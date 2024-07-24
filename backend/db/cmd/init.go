package cmd

import (
	"os"
	"sade-backend/config"
	"sade-backend/db"
	"sade-backend/pkg/logger"
)

var dbConn *db.Connection
var log *logger.Logger

func init() {
	log = logger.New("../log/logfile.log")
	configData, err := os.ReadFile("../config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	conf, err := config.New(configData, "../.env")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	dbConn = db.New(conf.DbConnection.Server, conf.DbConnection.User, conf.DbConnection.Pass, conf.DbConnection.DataBase, conf.DbConnection.Driver)
	dbConn.CreateTable()
}
