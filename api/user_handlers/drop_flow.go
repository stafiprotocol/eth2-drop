// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package user_handlers

import (
	dao_user "drop/dao/user"
	"drop/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RspDropFlow struct {
	DropFlowLatestDate string `json:"drop_flow_latest_date"`
}

// @Summary get drop flow
// @Description get drop flow latest date, will not close claim to set roothash when no dropflow yesterday
// @Tags v1
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropFlow}
// @Router /v1/drop_flow_latest [get]
func (h *Handler) HandleGetDropFlow(c *gin.Context) {
	meta, err := dao_user.GetMetaData(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}

	rsp := RspDropFlow{
		DropFlowLatestDate: meta.DropFlowLatestDate,
	}
	utils.Ok(c, "success", rsp)
}
