package dao_user

import "drop/pkg/db"

type MetaData struct {
	db.BaseModel
	DropRate           string `gorm:"type:varchar(30);not null;default:'0';column:drop_rate"`
	DropFlowLatestDate string `gorm:"type:varchar(10);not null;default:'0';column:drop_flow_latest_date"`
	LedgerLatestDate   string `gorm:"type:varchar(10);not null;default:'0';column:ledger_latest_date"`
	LatestClaimRound   int64  `gorm:"type:bigint;unsigned;not null;column:latest_claim_round"`
}

func UpOrInMetaData(db *db.WrapDb, c *MetaData) error {
	return db.Save(c).Error
}

func GetMetaData(db *db.WrapDb) (m *MetaData, err error) {
	m = &MetaData{}
	err = db.Take(m).Error
	return
}
