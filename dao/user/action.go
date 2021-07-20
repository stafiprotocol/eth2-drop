// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package dao_user

import "drop/pkg/db"

type DropAction struct {
	db.BaseModel
	DropDate         string `gorm:"type:varchar(10);not null;default:'0';column:drop_date;uniqueIndex"`
	CloseClaimTxHash string `gorm:"type:varchar(66);not null;default:'0x';column:close_claim_tx_hash"`
	SetRootTxhash    string `gorm:"type:varchar(66);not null;default:'0x';column:set_root_tx_hash"`
	Round            int64  `gorm:"type:bigint;unsigned;not null;column:round;uniqueIndex"`
	RootHash         string `gorm:"type:varchar(66);not null;default:'0x';column:root_hash"`
}

func UpOrInDropAction(db *db.WrapDb, c *DropAction) error {
	return db.Save(c).Error
}

func GetDropActionByDate(db *db.WrapDb, date string) (c *DropAction, err error) {
	c = &DropAction{}
	err = db.Take(c, "drop_date = ?", date).Error
	return
}
