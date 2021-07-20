// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package user_handlers

import (
	"drop/dao/user"
	"drop/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Drop struct {
	UserAddress string `json:"account"`
	DropAmount  string `json:"amount"`
}

type RspDropList struct {
	DropList []Drop `json:"drop_list"`
}

// @Summary get drop list
// @Description get drop list
// @Tags v1
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropList}
// @Router /v1/drop_list [get]
func (h *Handler) HandleGetDropList(c *gin.Context) {
	rsp := RspDropList{}
	rsp.DropList = make([]Drop, 0)

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
		rsp.DropList = append(rsp.DropList, Drop{
			UserAddress: l.UserAddress,
			DropAmount:  l.DropAmount,
		})
	}

	utils.Ok(c, "success", rsp)
}
