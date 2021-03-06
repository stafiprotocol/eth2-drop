// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package dao_user

import "drop/pkg/db"

type MetaData struct {
	db.BaseModel
	SyncStartDate      string `gorm:"type:varchar(30);not null;default:'0';column:sync_start_date"`       //sync dropflow start date
	DropFlowLatestDate string `gorm:"type:varchar(10);not null;default:'0';column:drop_flow_latest_date"` //latest date that has dropflow data
	LedgerLatestDate   string `gorm:"type:varchar(10);not null;default:'0';column:ledger_latest_date"`    //latest date apply to ledger, should <= DropFlowLatestDate
	LatestClaimRound   int64  `gorm:"type:bigint;unsigned;not null;column:latest_claim_round"`            //latest round
	DropIsOpen         int8   `gorm:"type:tinyint(1);not null;default:1;column:drop_is_open"`             //1 open 0 not
	RootHashSkipDate   string `gorm:"type:varchar(10);not null;default:'0';column:root_hash_skip_date"`   //root hash skip date
}

func UpOrInMetaData(db *db.WrapDb, c *MetaData) error {
	return db.Save(c).Error
}

func GetMetaData(db *db.WrapDb) (m *MetaData, err error) {
	m = &MetaData{}
	err = db.Take(m).Error
	return
}
