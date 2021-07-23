// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"context"
	contract_fis_drop "drop/contract/fis_drop"
	"drop/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func getRootHash(ledgerApi string, round int64) (string, error) {
	realUrl := fmt.Sprintf("%s/api/v1/root_hash?round=%d", ledgerApi, round)
	rsp, err := http.Get(realUrl)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get root hash status: %d", rsp.StatusCode)
	}
	rspBodyBts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if len(rspBodyBts) <= 0 {
		return "", fmt.Errorf("get root hash err %s", fmt.Errorf("body zero err"))
	}
	rspRootHash := RspRootHash{}
	err = json.Unmarshal(rspBodyBts, &rspRootHash)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if rspRootHash.Status != "80000" {
		return "", fmt.Errorf("get root hash err %s", fmt.Errorf("status err:%s", rspRootHash.Status))
	}
	return rspRootHash.Data.RootHash, nil
}

func getSkipDate(ledgerApi string) (string, error) {
	realUrl := fmt.Sprintf("%s/api/v1/skip_date", ledgerApi)
	rsp, err := http.Get(realUrl)
	if err != nil {
		return "", fmt.Errorf("get skip date err %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get skip date status: %d", rsp.StatusCode)
	}
	rspBodyBts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("get skip date err %s", err)
	}
	if len(rspBodyBts) <= 0 {
		return "", fmt.Errorf("get skip date err %s", fmt.Errorf("body zero err"))
	}
	rspSkipDate := RspSkipDate{}
	err = json.Unmarshal(rspBodyBts, &rspSkipDate)
	if err != nil {
		return "", fmt.Errorf("get skip date err %s", err)
	}
	if rspSkipDate.Status != "80000" {
		return "", fmt.Errorf("get skip date err %s", fmt.Errorf("status err:%s", rspSkipDate.Status))
	}
	return rspSkipDate.Data.SkipDate, nil
}

func getDropFLowLatest(ledgerApi string) (string, error) {
	realUrl := fmt.Sprintf("%s/api/v1/drop_flow_latest", ledgerApi)
	rsp, err := http.Get(realUrl)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get getDropFLow status: %d", rsp.StatusCode)
	}
	rspBodyBts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if len(rspBodyBts) <= 0 {
		return "", fmt.Errorf("get getDropFLow err %s", fmt.Errorf("body zero err"))
	}
	rspDropFlow := RspDropFlow{}
	err = json.Unmarshal(rspBodyBts, &rspDropFlow)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if rspDropFlow.Status != "80000" {
		return "", fmt.Errorf("get getDropFLow status err %s", rspDropFlow.Status)
	}
	return rspDropFlow.Data.DropFlowLatestDate, nil
}

type RspRootHash struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		RootHash string `json:"root_hash"`
	} `json:"data"`
}

type RspSkipDate struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		SkipDate string `json:"skip_date"`
	} `json:"data"`
}

type RspDropFlow struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		DropFlowLatestDate string `json:"drop_flow_latest_date"`
	} `json:"data"`
}

func waitClient(ethApi string) *ethclient.Client {
	retry := 0
	var client *ethclient.Client
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("dial client reach retry"))
		}
		client, err = ethclient.Dial(ethApi)
		if err != nil {
			logrus.Warn("dail client failed ,watting...", " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
		break
	}
	return client
}

func waitFisDropContract(dropContract string, client *ethclient.Client) *contract_fis_drop.FisDropREth {
	retry := 0
	var fisDropContract *contract_fis_drop.FisDropREth
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("NewFisDropREth reach retry"))
		}
		fisDropContract, err = contract_fis_drop.NewFisDropREth(common.HexToAddress(dropContract), client)
		if err != nil {
			logrus.Warn("NewFisDropREth failed ,watting...", " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
		break
	}
	return fisDropContract
}

func waitDateHashDrop(fisDropContract *contract_fis_drop.FisDropREth, dateHash [32]byte) bool {
	retry := 0
	var dateHasDrop bool
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("get DateDrop reach retry"))
		}
		dateHasDrop, err = fisDropContract.DateDrop(&callOpts, dateHash)
		if err != nil {
			logrus.Warn("get DateDrop failed ,watting...", " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
		break
	}
	return dateHasDrop
}

func waitDropFlowLatestDate(ledgerApi string) string {
	retry := 0
	dropFlowLatestDate := ""
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("getDropFLowLatest  reach retry"))
		}
		dropFlowLatestDate, err = getDropFLowLatest(ledgerApi)
		if err != nil {
			logrus.Warnf("getDropFLowLatest failed: %s", err)
			time.Sleep(waitTime)
			continue
		}
		break
	}
	return dropFlowLatestDate
}

func waitToGetSkipDate(ledgerApi string) string {
	retry := 0
	skipDate := ""
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("getSkipDate  reach retry"))
		}
		skipDate, err = getSkipDate(ledgerApi)
		if err != nil {
			logrus.Warnf("getSkipDate failed: %s", err)
			time.Sleep(waitTime)
			continue
		}
		break
	}
	return skipDate
}

func waitToGetClaimOpen(fisDropContract *contract_fis_drop.FisDropREth) bool {
	retry := 0
	var claimOpen bool
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("get ClaimOpen  reach retry"))
		}
		claimOpen, err = fisDropContract.ClaimOpen(&callOpts)
		if err != nil {
			logrus.Warnf("get ClaimOpen failed: %s", err)
			time.Sleep(waitTime)
			continue
		}
		break
	}
	return claimOpen
}

func waitToGetClaimRound(fisDropContract *contract_fis_drop.FisDropREth) *big.Int {
	retry := 0
	var claimRoundOnchain *big.Int
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("get ClaimRound  reach retry"))
		}
		claimRoundOnchain, err = fisDropContract.ClaimRound(&callOpts)
		if err != nil {
			logrus.Warnf("get ClaimRound failed: %s", err)
			time.Sleep(waitTime)
			continue
		}
		break
	}
	return claimRoundOnchain
}

func sendCloseClaimTxAndWait(client *ethclient.Client, fisDropContract *contract_fis_drop.FisDropREth, txOpts *bind.TransactOpts) {
	retry := 0
	var tx *types.Transaction
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("send CloseClaim tx reach retry"))
		}
		tx, err = fisDropContract.CloseClaim(txOpts)
		if err != nil {
			logrus.Warnf("send CloseClaim tx failed: %s", err)
			//return if executed or voted
			if strings.Contains(err.Error(), "proposal already executed/cancelled") ||
				strings.Contains(err.Error(), "relayer already voted") {
				return
			}
			time.Sleep(waitTime)
			continue
		}
		break
	}

	//wait close claim tx onchain
	retry = 0
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("check ClaimClose tx onchain reach retry"))
		}
		_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err == nil && !isPending {
			break
		} else {
			logrus.Warn("check CloseClaim tx onchain failed ,watting...", " isPending ", isPending, " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
	}
}

func waitUntilClaimClose(fisDropContract *contract_fis_drop.FisDropREth) {
	retry := 0
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("check ClaimOpen   reach retry"))
		}
		open, err := fisDropContract.ClaimOpen(&callOpts)
		if err == nil && !open {
			break
		} else {
			logrus.Warn("check ClaimOpen failed ,watting...", " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
	}
}

func waitUntilClaimOpen(fisDropContract *contract_fis_drop.FisDropREth) {
	retry := 0
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("check ClaimOpen   reach retry"))
		}
		open, err := fisDropContract.ClaimOpen(&callOpts)
		if err == nil && open {
			break
		} else {
			logrus.Warn("check ClaimOpen failed ,watting...", " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
	}
}

func waitToGetRootHash(ledgerAPi string, round int64) (rootHash string, skip bool) {
	retry := 0
	willUseRootHash := ""
	skipDate := ""
	skipToday := false
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("waitToGetRootHashOrSkipDate  reach retry"))
		}
		willUseRootHash, err = getRootHash(ledgerAPi, round)
		if err != nil {
			logrus.Warnf("getRootHash failed: %s", err)
			skipDate, err = getSkipDate(ledgerAPi)
			if err != nil {
				logrus.Warnf("getSkipDate failed: %s", err)
				time.Sleep(waitTime * 2)
				continue
			}
			if utils.GetNowUTC8Date() != skipDate {
				time.Sleep(waitTime * 2)
				continue
			} else {
				skipToday = true
				break
			}
		}
		break
	}
	return willUseRootHash, skipToday
}

func sendSetRootHashTxAndWait(client *ethclient.Client, fisDropContract *contract_fis_drop.FisDropREth,
	txOpts *bind.TransactOpts, dateHash [32]byte, _merkleRoot [32]byte) {
	retry := 0
	var err error
	var tx *types.Transaction
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("send SetMerkleRoot tx reach retry"))
		}
		tx, err = fisDropContract.SetMerkleRoot(txOpts, dateHash, _merkleRoot)
		if err != nil {
			logrus.Warnf("send SetMerkleRoot tx failed: %s", err)
			//return if executed or voted
			if strings.Contains(err.Error(), "proposal already executed/cancelled") ||
				strings.Contains(err.Error(), "relayer already voted") ||
				strings.Contains(err.Error(), "this date has drop") {
				return
			}
			time.Sleep(waitTime)
			continue
		}
		break
	}

	//wait root hash set tx onchain
	retry = 0
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("check SetMerkleRoot tx  reach retry"))
		}
		_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err == nil && !isPending {
			break
		} else {
			logrus.Warn("check SetMerkleRoot tx failed ,watting...", " isPending ", isPending, " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
	}

}

func sendOpenClaimTxAndWait(client *ethclient.Client, fisDropContract *contract_fis_drop.FisDropREth, txOpts *bind.TransactOpts) {
	retry := 0
	var tx *types.Transaction
	var err error
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("send OpenClaim tx reach retry"))
		}
		tx, err = fisDropContract.OpenClaim(txOpts)
		if err != nil {
			logrus.Warnf("send OpenClaim tx failed: %s", err)
			//return if executed or voted
			if strings.Contains(err.Error(), "proposal already executed/cancelled") ||
				strings.Contains(err.Error(), "relayer already voted") {
				return
			}
			time.Sleep(waitTime)
			continue
		}
		break
	}

	//wait open claim tx onchain
	retry = 0
	for {
		if retry > reTryLimit {
			panic(fmt.Errorf("check OpenClaim tx onchain reach retry"))
		}
		_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err == nil && !isPending {
			break
		} else {
			logrus.Warn("check OpenClaim tx onchain failed ,watting...", " isPending ", isPending, " err ", err)
			time.Sleep(waitTime)
			retry++
			continue
		}
	}
}
