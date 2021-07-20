// Copyright 2020 tpkeeper
// SPDX-License-Identifier: LGPL-3.0-only

package server

import (
	"drop/task"
	"time"

	"github.com/sirupsen/logrus"
)

func (svr *Server) Task() {
	ticker := time.NewTicker(time.Duration(svr.taskTicker) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logrus.Info("SyncDropFlow start...")
			err := task.SyncDropFlow(svr.db, svr.syncStartDate, svr.rEthStatApi)
			if err != nil {
				logrus.Error("SyncDropFlow err: ", err)
			}
			logrus.Info("SyncDropFlow end")

			logrus.Info("ApllyToLedger start...")
			err = task.ApllyToLedger(svr.db)
			if err != nil {
				logrus.Error("ApllyToLedger err: ", err)
			}
			logrus.Info("ApllyToLedger end")

			logrus.Info("CheckAndDrop start...")
			err = task.CheckAndSnapshot(svr.db, svr.ethApi, svr.dropContract)
			if err != nil {
				logrus.Error("CheckAndSnapshot err: ", err)
			}
			logrus.Info("CheckAndDrop end")
		}

	}
}
