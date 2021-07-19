package user_handlers

import (
	dao_user "drop/dao/user"
	"drop/pkg/utils"
	"drop/pkg/utils/distributor"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RspRootHash struct {
	RootHash string `json:"root_hash"`
}

// @Summary get root hash
// @Description get root hash
// @Tags v1
// @Param round query string true "round"
// @Produce json
// @Success 200 {object} utils.Rsp{data=RspDropInfo}
// @Router /v1/root_hash [get]
func (h *Handler) HandleGetRootHash(c *gin.Context) {
	roundStr := c.Query("round")
	round, err := strconv.Atoi(roundStr)
	if err != nil {
		utils.Err(c, "param err")
		return
	}
	meta, err := dao_user.GetMetaData(h.db)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	if meta.LatestClaimRound != int64(round) {
		utils.Err(c, "new round not ready yet")
		return
	}
	list, err := dao_user.GetSnapshotListByRound(h.db, meta.LatestClaimRound)
	if err != nil {
		utils.Err(c, err.Error())
		return
	}
	if len(list) == 0 {
		utils.Err(c, "drop list len is zero")
		return
	}
	hash := distributor.GetRootHash(list)
	rsp := RspRootHash{
		RootHash: hash,
	}
	utils.Ok(c, "success", rsp)
}
