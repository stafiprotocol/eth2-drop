package dao_user

import "drop/pkg/db"

// flow drop info
type DropFlow struct {
	db.BaseModel
	UserAddress string `gorm:"type:varchar(42);not null;default:'';column:user_address;uniqueIndex:uni_idx_user_date"`
	REthAmount  string `gorm:"type:varchar(30);not null;default:'0';column:reth_amount"`
	DropRate    string `gorm:"type:varchar(30);not null;default:'0';column:drop_rate"`
	DropAmount  string `gorm:"type:varchar(30);not null;default:'0';column:drop_amount"`
	DepositDate string `gorm:"type:varchar(10);not null;default:'0';column:deposit_date;uniqueIndex:uni_idx_user_date"`
}

func UpOrInDropFlow(db *db.WrapDb, c *DropFlow) error {
	return db.Save(c).Error
}

func GetDropFlowListByDate(db *db.WrapDb, date string) (cmp []*DropFlow, err error) {
	err = db.Find(&cmp, "deposit_date = ?", date).Error
	return
}
