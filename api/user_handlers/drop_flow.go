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
// @Description get drop flow
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