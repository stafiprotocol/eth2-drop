// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package server

import (
	"drop/api"
	dao_user "drop/dao/user"
	"drop/pkg/config"
	"drop/pkg/db"
	"drop/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	listenAddr string
	httpServer *http.Server
	taskTicker int64

	ethApi        string
	rEthStatApi   string
	dropContract  string
	dropTime      int64
	chainId       int64
	syncStartDate string
	dropIsOpen    bool
	db            *db.WrapDb
}

func NewServer(cfg *config.Config, dao *db.WrapDb) (*Server, error) {
	s := &Server{
		listenAddr:    cfg.ListenAddr,
		taskTicker:    cfg.TaskTicker,
		ethApi:        cfg.EthApi,
		rEthStatApi:   cfg.REthStatApi,
		dropContract:  cfg.DropContract,
		dropTime:      cfg.DropTime,
		syncStartDate: cfg.SyncStartDate,
		dropIsOpen:    cfg.DropIsOpen,

		chainId: cfg.ChainId,
		db:      dao,
	}

	handler := s.InitHandler()

	s.httpServer = &http.Server{
		Addr:         s.listenAddr,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return s, nil
}

func (svr *Server) InitHandler() http.Handler {
	return api.InitRouters(svr.db)
}

func (svr *Server) ApiServer() {
	logrus.Infof("Gin server start on %s", svr.listenAddr)
	err := svr.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.Errorf("Gin server start err: %s", err.Error())
		shutdownRequestChannel <- struct{}{} //shutdown server
		return
	}
	logrus.Infof("Gin server done on %s", svr.listenAddr)
}

//check and init dropFlowLatestDate LedgerLatestDate
func (svr *Server) InitOrUpdateMetaData() error {
	if len(svr.syncStartDate) != 8 {
		return fmt.Errorf("syncStartDate format err")
	}
	meta, _ := dao_user.GetMetaData(svr.db)
	if svr.syncStartDate > meta.DropFlowLatestDate {
		newDay, err := utils.SubOneDay(svr.syncStartDate)
		if err != nil {
			return err
		}
		meta.DropFlowLatestDate = newDay
	}

	if svr.syncStartDate > meta.LedgerLatestDate {
		newDay, err := utils.SubOneDay(svr.syncStartDate)
		if err != nil {
			return err
		}
		meta.LedgerLatestDate = newDay
	}
	meta.SyncStartDate = svr.syncStartDate
	if svr.dropIsOpen {
		meta.DropIsOpen = 1
	} else {
		meta.DropIsOpen = 0
	}
	return dao_user.UpOrInMetaData(svr.db, meta)
}

func (svr *Server) Start() error {
	err := svr.InitOrUpdateMetaData()
	if err != nil {
		return err
	}
	utils.SafeGoWithRestart(svr.ApiServer)
	utils.SafeGoWithRestart(svr.Task)
	return nil
}

func (svr *Server) Stop() {
	if svr.httpServer != nil {
		err := svr.httpServer.Close()
		if err != nil {
			logrus.Errorf("Problem shutdown Gin server :%s", err.Error())
		}
	}
}
