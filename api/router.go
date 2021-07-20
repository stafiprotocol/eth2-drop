// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"drop/api/user_handlers"
	"drop/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouters(db *db.WrapDb) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/static", "./static")
	router.Use(Cors())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandler := user_handlers.NewHandler(db)
	router.GET("api/v1/drop_rate", userHandler.HandleGetDropRate)
	router.GET("api/v1/drop_info", userHandler.HandleGetDropInfo)
	router.GET("api/v1/root_hash", userHandler.HandleGetRootHash)
	router.GET("api/v1/drop_flow_latest", userHandler.HandleGetDropFlow)

	return router
}
