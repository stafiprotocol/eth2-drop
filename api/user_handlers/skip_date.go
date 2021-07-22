// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package user_handlers

import (
	dao_user "drop/dao/user"
	"drop/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RspSkipDate struct {
	SkipDate string `json:"skip_date"`
}

// @Summary get skip date
// @Description get skip date, will not set root hash and just close claim if now day skip
// @Tags v1
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropInfo}
// @Router /v1/skip_date [get]
func (h *Handler) HandleGetSkipDate(c *gin.Context) {
	meta, err := dao_user.GetMetaData(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	rsp := RspSkipDate{
		SkipDate: meta.RootHashSkipDate,
	}
	utils.Ok(c, "success", rsp)
}
