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
)

func SyncDropFlow(db *db.WrapDb, startDate, rethStatApi string) error {
	meta, err := dao_user.GetMetaData(db)
	if err != nil {
		return err
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

		dropRate, err := utils.GetDropRate(startDate, requestDay)
		if err != nil {
			return err
		}
		dropRateDecimal, err := decimal.NewFromString(dropRate)
		if err != nil {
			return fmt.Errorf("droprate:%s err :%s", dropRate, err)
		}

		//transaction start
		tx := db.NewTransaction()
		for _, l := range rspREth.Data.List {
			if !common.IsHexAddress(l.Address) {
				tx.RollbackTransaction()
				panic(fmt.Errorf("not common eth address: %s", l.Address))
			}
			rethAmountDecimal, err := decimal.NewFromString(l.Amount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("reth amount not right: %s", l.Amount))
			}
			dropAmountDecimal := rethAmountDecimal.Mul(dropRateDecimal).Div(decimal.New(1, 18))

			//get drop from db
			dropFlow, _ := dao_user.GetDropFlowByUserDate(tx, l.Address, requestDay)
			oldRethAmountDecimal, err := decimal.NewFromString(dropFlow.REthAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("old reth amount not right: %s", dropFlow.REthAmount))
			}
			oldDropAmountDecimal, err := decimal.NewFromString(dropFlow.DropAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("old drop amount not right: %s", dropFlow.DropAmount))
			}

			newRethAmountStr := rethAmountDecimal.Add(oldRethAmountDecimal).StringFixed(0)
			newDropAmountStr := dropAmountDecimal.Add(oldDropAmountDecimal).StringFixed(0)
			//update amount
			dropFlow.REthAmount = newRethAmountStr
			dropFlow.DropAmount = newDropAmountStr
			//update data
			dropFlow.UserAddress = l.Address
			dropFlow.DropRate = dropRate
			dropFlow.DepositDate = requestDay

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
			Address string `json:"address"`
			Amount  string `json:"amount"`
		} `json:"list"`
		Date string `json:"date"`
	} `json:"data"`
}
