package dao_user

import (
	"drop/pkg/db"

	"gorm.io/gorm"
)

// aggregate drop info
type DropLedger struct {
	db.BaseModel
	UserAddress            string `gorm:"type:varchar(42);not null;default:'';column:user_address;uniqueIndex"`
	TotalDropAmount        string `gorm:"type:varchar(30);not null;default:'0';column:total_drop_amount"`
	TotalClaimedDropAmount string `gorm:"type:varchar(30);not null;default:'0';column:total_claimed_drop_amount"`
	TotalREthAmount        string `gorm:"type:varchar(30);not null;default:'0';column:total_reth_amount"`
	LatestDate             string `gorm:"type:varchar(10);not null;default:'0';column:latest_date;"` //last dropflow's latestdate
}

func UpOrInDropLedger(db *db.WrapDb, c *DropLedger) error {
	return db.Save(c).Error
}

func GetDropLedgerList(db *db.WrapDb) (cmp []*DropLedger, err error) {
	err = db.Find(&cmp).Error
	return
}

func GetDropLedgerByUser(db *db.WrapDb, user string) (banker *DropLedger, err error) {
	banker = &DropLedger{}
	err = db.Take(banker, "user_address = ?", user).Error
	if err == gorm.ErrRecordNotFound {
		banker.TotalClaimedDropAmount = "0"
		banker.TotalDropAmount = "0"
		banker.TotalREthAmount = "0"
	}
	return
}
