package main

import (
	_ "drop/cmd/ledgerd/docs"
	"drop/dao/migrate"
	"drop/pkg/config"
	"drop/pkg/db"
	"drop/pkg/log"
	"drop/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"runtime/debug"
)

func _main() error {
	cfg, err := config.Load("ledger_conf.toml")
	if err != nil {
		fmt.Printf("loadConfig err: %s", err)
		return err
	}
	log.InitLogFile(cfg.LogFilePath+"/ledger")
	logrus.Infof("config info:%+v ", cfg)

	//init db
	db, err := db.NewDB(&db.Config{
		Host:   cfg.Db.Host,
		Port:   cfg.Db.Port,
		User:   cfg.Db.User,
		Pass:   cfg.Db.Pwd,
		DBName: cfg.Db.Name,
		Mode:   cfg.Mode})
	if err != nil {
		logrus.Errorf("db err: %s", err)
		return err
	}
	logrus.Infof("db connect success")

	//interrupt signal
	ctx := server.ShutdownListener()

	defer func() {
		sqlDb, err := db.DB.DB()
		if err != nil {
			logrus.Errorf("db.DB() err: %s", err)
			return
		}
		logrus.Infof("shutting down the db ...")
		sqlDb.Close()
	}()

	err = migrate.AutoMigrate(db)
	if err != nil {
		logrus.Errorf("dao autoMigrate err: %s", err)
		return err
	}
	//server
	server, err := server.NewServer(cfg, db)
	if err != nil {
		logrus.Errorf("new server err: %s", err)
		return err
	}
	err = server.Start()
	if err != nil {
		logrus.Errorf("server start err: %s", err)
		return err
	}
	defer func() {
		logrus.Infof("shutting down server ...")
		server.Stop()
	}()

	<-ctx.Done()
	return nil
}

// @title drop API
// @version 1.0
// @description drop api document.

// @contact.name tk
// @contact.email tpkeeper.me@gmail.com

// @host xxxxx:8081
// @BasePath /api

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	debug.SetGCPercent(40)
	err := _main()
	if err != nil {
		os.Exit(1)
	}
}
