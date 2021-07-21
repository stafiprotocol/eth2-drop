// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package user_handlers

import (
	dao_user "drop/dao/user"
	"drop/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DropInfo struct {
	UserAddress       string `json:"user_address"`
	TotalDropAmount   string `json:"total_drop_amount"`
	TotalREthAmount   string `json:"total_eth_amount"`
	ClaimedDropAmount string `json:"claimed_drop_amount"`
}

type RspDropInfo struct {
	DropIsOpen bool        `json:"drop_is_open"`
	DropInfo   interface{} `json:"drop_info"`
	DropList   []Drop      `json:"drop_list"`
	TxList     []string    `json:"tx_list"`
}

// @Summary get user drop info
// @Description get user drop info
// @Tags v1
// @Param user_address query string true "user address"
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropInfo}
// @Router /v1/drop_info [get]
func (h *Handler) HandleGetDropInfo(c *gin.Context) {
	userAddress := c.Query("user_address")
	if !common.IsHexAddress(userAddress) {
		utils.Err(c, "param err")
		return
	}
	//get drop info
	userInfo, err := dao_user.GetDropLedgerByUser(h.db, userAddress)
	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Err(c, err.Error())
		return
	}
	var dropInfo interface{} = struct{}{}
	if err != nil && err == gorm.ErrRecordNotFound {
		dropInfo = struct{}{}
	}
	if err == nil {
		dropInfo = DropInfo{
			UserAddress:       userAddress,
			TotalDropAmount:   userInfo.TotalDropAmount,
			TotalREthAmount:   userInfo.TotalREthAmount,
			ClaimedDropAmount: userInfo.TotalClaimedDropAmount,
		}
	}
	rsp := RspDropInfo{
		DropInfo: dropInfo,
	}

	//get drop list
	dropList := make([]Drop, 0)
	lastRound, err := dao_user.GetSnapshotLastRound(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	list, err := dao_user.GetSnapshotListByRound(h.db, lastRound)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	for _, l := range list {
		dropList = append(dropList, Drop{
			UserAddress: l.UserAddress,
			DropAmount:  l.DropAmount,
		})
	}
	rsp.DropList = dropList

	//get txlist
	rsp.TxList = make([]string, 0)
	txList, err := dao_user.GetDropFlowListByUser(h.db, userAddress)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	for _, tx := range txList {
		rsp.TxList = append(rsp.TxList, tx.Txhash)
	}

	meta, err := dao_user.GetMetaData(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	rsp.DropIsOpen = meta.DropIsOpen == 1
	utils.Ok(c, "success", rsp)
}
