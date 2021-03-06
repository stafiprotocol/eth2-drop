// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package task

import (
	dao_user "drop/dao/user"
	"drop/pkg/db"
	"drop/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func SyncDropFlow(db *db.WrapDb, startDate, rethStatApi string) error {
	meta, err := dao_user.GetMetaData(db)
	if err != nil {
		return err
	}
	if meta.DropIsOpen == 0 {
		logrus.Info("drop not open yet")
		return nil
	}
	yesterDay := utils.GetYesterdayUTC8Date()
	dayInMeta := meta.DropFlowLatestDate

	//only sync after 01:00
	newDaySeconds := utils.GetNewDayUtc8Seconds()
	if newDaySeconds < 60*60 {
		return nil
	}

	for dayInMeta < yesterDay {
		requestDay, err := utils.AddOneDay(dayInMeta)
		if err != nil {
			return err
		}
		dayInMeta = requestDay

		realUrl := fmt.Sprintf("%s?date=%s", rethStatApi, requestDay)
		rsp, err := http.Get(realUrl)
		if err != nil {
			return err
		}
		if rsp.StatusCode != http.StatusOK {
			return fmt.Errorf("status: %d", rsp.StatusCode)
		}
		rspBodyBts, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return err
		}
		if len(rspBodyBts) <= 0 {
			return fmt.Errorf("body err")
		}

		rspREth := RspREth{}
		err = json.Unmarshal(rspBodyBts, &rspREth)
		if err != nil {
			return err
		}

		if len(rspREth.Data.List) == 0 {
			//not store db now
			continue
		}
		if rspREth.Data.Date != requestDay {
			return fmt.Errorf("requestDay:%s != repDate:%s", requestDay, rspREth.Data.Date)
		}

		//transaction start
		tx := db.NewTransaction()
		for _, l := range rspREth.Data.List {
			//get droprate
			dropRate, err := utils.GetDropRateFromTimestamp(startDate, l.Timestamp)
			if err != nil {
				tx.RollbackTransaction()
				panic(err)
			}
			dropRateDecimal, err := decimal.NewFromString(dropRate)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("droprate:%s err :%s", dropRate, err))
			}

			//check address
			if !common.IsHexAddress(l.Address) {
				tx.RollbackTransaction()
				panic(fmt.Errorf("not common eth address: %s", l.Address))
			}
			//check tx hash
			if len(l.Hash) != 66 {
				tx.RollbackTransaction()
				panic(fmt.Errorf("tx hash len no right: address %s ,hash: %s", l.Address, l.Hash))
			}
			//checkout eth amount
			rethAmountDecimal, err := decimal.NewFromString(l.Amount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("reth amount not right: %s", l.Amount))
			}
			//call drop amount
			dropAmountDecimal := rethAmountDecimal.Mul(dropRateDecimal).Div(decimal.New(1, 18))

			//get drop from db
			dropFlow, _ := dao_user.GetDropFlowByUserTx(tx, l.Address, l.Hash)
			//update amount
			dropFlow.REthAmount = rethAmountDecimal.StringFixed(0)
			dropFlow.DropAmount = dropAmountDecimal.StringFixed(0)
			//update data
			dropFlow.UserAddress = l.Address
			dropFlow.DropRate = dropRate
			dropFlow.DepositDate = requestDay
			dropFlow.TxTimestamp = l.Timestamp
			dropFlow.Txhash = l.Hash

			err = dao_user.UpOrInDropFlow(tx, dropFlow)
			if err != nil {
				tx.RollbackTransaction()
				return fmt.Errorf("UpOrInDropFlow dropflow: %+v  err: %s", dropFlow, err)
			}
		}
		meta.DropFlowLatestDate = requestDay
		err = dao_user.UpOrInMetaData(tx, meta)
		if err != nil {
			tx.RollbackTransaction()
			return fmt.Errorf("UpOrInMetaData failed meta: %+v  err: %s", meta, err)
		}
		err = tx.CommitTransaction()
		if err != nil {
			panic(fmt.Errorf("tx.CommitTransaction err: %s", err))
		}

	}

	return nil
}

type RspREth struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Address   string `json:"address"`
			Amount    string `json:"amount"`
			Timestamp string `json:"operateTime"`
			Hash      string `json:"hash"`
		} `json:"list"`
		Date string `json:"date"`
	} `json:"data"`
}
