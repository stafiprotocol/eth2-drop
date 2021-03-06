// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package user_handlers

import (
	dao_user "drop/dao/user"
	"drop/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RspDropRate struct {
	DropRate string `json:"drop_rate"`
}

// @Summary get drop rate
// @Description get drop rate
// @Tags v1
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropRate}
// @Router /v1/drop_rate [get]
func (h *Handler) HandleGetDropRate(c *gin.Context) {
	timeStamp := c.Query("timestamp")
	nowDate := utils.GetNowUTC8Date()
	meta, err := dao_user.GetMetaData(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	//return 0 if not open yet
	if meta.DropIsOpen == 0 {
		utils.Ok(c, "success", RspDropRate{
			DropRate: "0",
		})
		return
	}

	if len(timeStamp) > 0 {
		rate, err := utils.GetDropRateFromTimestamp(meta.SyncStartDate, timeStamp)
		if err != nil {
			utils.Err(c, err.Error())
			return
		}
		utils.Ok(c, "success", RspDropRate{
			DropRate: rate,
		})
		return
	}

	dropRate, err := utils.GetDropRate(meta.SyncStartDate, nowDate)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	utils.Ok(c, "success", RspDropRate{
		DropRate: dropRate,
	})
}
