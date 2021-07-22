// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package task

import (
	"context"
	"drop/contract/fis_drop"
	dao_user "drop/dao/user"
	"drop/pkg/db"
	"drop/pkg/utils"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func CheckAndSnapshot(db *db.WrapDb, ethApi, fisDropContractAddress string) error {
	client, err := ethclient.Dial(ethApi)
	if err != nil {
		return err
	}
	fisDropContract, err := contract_fis_drop.NewFisDropREth(common.HexToAddress(fisDropContractAddress), client)
	if err != nil {
		return fmt.Errorf("NewFisDropREth err %s", err)
	}
	callOpts := bind.CallOpts{
		Pending:     false,
		From:        [20]byte{},
		BlockNumber: nil,
		Context:     context.Background(),
	}
	//snapshot when claim close && roundOnchain==round in snapshot
	isClaimOpen, err := fisDropContract.ClaimOpen(&callOpts)
	if err != nil {
		return err
	}
	//skip if claimopen
	if isClaimOpen {
		return fmt.Errorf("claim is open yet")
	}

	meta, err := dao_user.GetMetaData(db)
	if err != nil {
		return err
	}
	//skip if drop not open
	if meta.DropIsOpen == 0 {
		return fmt.Errorf("drop is not open yet")
	}
	roundOnchain, err := fisDropContract.ClaimRound(&callOpts)
	if err != nil {
		return err
	}
	//skip if roundonchain != round in db
	if roundOnchain.Int64() != meta.LatestClaimRound {
		return fmt.Errorf("round onchain != meta.latestClaimRound")
	}
	//sync claim data
	lastRound, err := dao_user.GetSnapshotLastRound(db)
	if err != nil {
		return err
	}
	list, err := dao_user.GetSnapshotListByRound(db, lastRound)
	if err != nil {
		return err
	}

	//transaction start
	tx := db.NewTransaction()
	for i, l := range list {
		time.Sleep(time.Millisecond * 200)
		//skip if has claimed
		if l.Claimed == 1 {
			continue
		}
		isClaimed, err := fisDropContract.IsClaimed(&callOpts, big.NewInt(lastRound), big.NewInt(int64(i)))
		if err != nil {
			tx.RollbackTransaction()
			return fmt.Errorf("fisDropContract.IsClaimed err:%s ,round:%d address:%s", err, lastRound, l.UserAddress)
		}
		if isClaimed {
			dropLedger, err := dao_user.GetDropLedgerByUser(tx, l.UserAddress)
			if err != nil {
				tx.RollbackTransaction()
				return fmt.Errorf("GetDropLedgerByUser err: %s", err.Error())
			}
			oldTotalClaimed, err := decimal.NewFromString(dropLedger.TotalClaimedDropAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(err)
			}
			newClaimed, err := decimal.NewFromString(l.DropAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(err)
			}
			//cal new total claimed
			newTotalClaimedStr := oldTotalClaimed.Add(newClaimed).StringFixed(0)

			//update db
			dropLedger.TotalClaimedDropAmount = newTotalClaimedStr
			l.Claimed = 1
			err = dao_user.UpdateSnapshot(tx, l)
			if err != nil {
				tx.RollbackTransaction()
				return err
			}
			err = dao_user.UpOrInDropLedger(tx, dropLedger)
			if err != nil {
				tx.RollbackTransaction()
				return err
			}
		}
	}

	//gen new snapshot
	dropLedgerList, err := dao_user.GetDropLedgerList(tx)
	if err != nil {
		tx.RollbackTransaction()
		return err
	}

	hashSnapshotThisRound := false
	for _, l := range dropLedgerList {
		totalClaimedDropAmountDeci, err := decimal.NewFromString(l.TotalClaimedDropAmount)
		if err != nil {
			tx.RollbackTransaction()
			panic(err)
		}
		totalDropAmountDeci, err := decimal.NewFromString(l.TotalDropAmount)
		if err != nil {
			tx.RollbackTransaction()
			panic(err)
		}

		newDropAmount := decimal.NewFromInt(0)
		if totalDropAmountDeci.GreaterThanOrEqual(totalClaimedDropAmountDeci) {
			newDropAmount = totalDropAmountDeci.Sub(totalClaimedDropAmountDeci)
		}

		//skip if drop amount is zero
		if newDropAmount.Equal(decimal.NewFromInt(0)) {
			continue
		}

		snapShot := dao_user.Snapshot{
			UserAddress: l.UserAddress,
			Round:       lastRound + 1,
			DropAmount:  newDropAmount.StringFixed(0),
			Claimed:     0,
		}
		hashSnapshotThisRound = true
		err = dao_user.UpOrInSnapshot(tx, &snapShot)
		if err != nil {
			tx.RollbackTransaction()
			return err
		}
	}

	if !hashSnapshotThisRound {
		//if has no snapshot, skip root hash today
		meta.RootHashSkipDate = utils.GetNowUTC8Date()
	} else {
		//new round flag
		meta.LatestClaimRound = lastRound + 1
	}

	//update meta data
	err = dao_user.UpOrInMetaData(tx, meta)
	if err != nil {
		tx.RollbackTransaction()
		return err
	}

	err = tx.CommitTransaction()
	if err != nil {
		panic(fmt.Errorf("tx.CommitTransaction err: %s", err))
	}
	logrus.Info("tx commitTransaction ok")

	return nil
}
