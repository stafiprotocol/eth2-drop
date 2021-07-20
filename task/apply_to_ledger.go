package task

import (
	dao_user "drop/dao/user"
	"drop/pkg/db"
	"drop/pkg/utils"
	"fmt"

	"github.com/shopspring/decimal"
)

func ApllyToLedger(db *db.WrapDb) error {
	meta, err := dao_user.GetMetaData(db)
	if err != nil {
		return err
	}
	dropFlowLatestDate := meta.DropFlowLatestDate
	ledgerLatesDate := meta.LedgerLatestDate

	for ledgerLatesDate < dropFlowLatestDate {
		willApplyDate, err := utils.AddOneDay(ledgerLatesDate)
		if err != nil {
			return err
		}
		ledgerLatesDate = willApplyDate
		dropList, err := dao_user.GetDropFlowListByDate(db, willApplyDate)
		if err != nil {
			return err
		}
		if len(dropList) == 0 {
			continue
		}
		tx := db.NewTransaction()
		for _, drop := range dropList {
			ledger, _ := dao_user.GetDropLedgerByUser(tx, drop.UserAddress)
			ledger.UserAddress = drop.UserAddress
			oldTotalREthAmountDecimal, err := decimal.NewFromString(ledger.TotalREthAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("ledger total reth amount failed, amount:%s err:%s", ledger.TotalREthAmount, err))
			}
			dropREthAmountDecimal, err := decimal.NewFromString(drop.REthAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("drop flow reth amount failed, amount:%s err:%s", drop.REthAmount, err))
			}
			newTotalREthAmountStr := oldTotalREthAmountDecimal.Add(dropREthAmountDecimal).StringFixed(0)

			oldTotalDropAmountDecimal, err := decimal.NewFromString(ledger.TotalDropAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("ledger total drop amount failed, amount:%s err:%s", ledger.TotalDropAmount, err))
			}
			dropDropAmountDecimal, err := decimal.NewFromString(drop.DropAmount)
			if err != nil {
				tx.RollbackTransaction()
				panic(fmt.Errorf("drop flow drop amount failed, amount:%s err:%s", drop.DropAmount, err))
			}
			newTotalDropAmountStr := oldTotalDropAmountDecimal.Add(dropDropAmountDecimal).StringFixed(0)

			ledger.TotalREthAmount = newTotalREthAmountStr
			ledger.TotalDropAmount = newTotalDropAmountStr
			ledger.LatestDate = willApplyDate
			err = dao_user.UpOrInDropLedger(tx, ledger)
			if err != nil {
				tx.RollbackTransaction()
				return fmt.Errorf("UpOrInDropLedger failed, ledger: %+v  err: %s", ledger, err)
			}
		}
		meta.LedgerLatestDate = willApplyDate

		err = dao_user.UpOrInMetaData(tx, meta)
		if err != nil {
			tx.RollbackTransaction()
			return fmt.Errorf("UpOrInMetaData failed， meta: %+v  err: %s", meta, err)
		}
		err = tx.CommitTransaction()
		if err != nil {
			panic(fmt.Errorf("tx.CommitTransaction err: %s",err))
		}

	}
	return nil
}
