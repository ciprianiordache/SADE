package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"sade-backend/api/middleware"
	"sade-backend/api/server"
	"sade-backend/config"
	"sade-backend/pkg/logger"
)

var srv *server.Server
var router *gin.Engine
var log *logger.Logger

func init() {
	log = logger.New("../log/logfile.log")

	configData, err := os.ReadFile("../config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	conf, err := config.New(configData, "../.env")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Static("/media", filepath.Join("..", "media"))
	router.Use(middleware.Logger())
	router.Use(middleware.Session(conf.Session.Name, conf.Session.Key))
	router.Use(middleware.CORSMiddleware(conf.OriginURL))
	router.Use(middleware.RateLimit())

	srv = server.New(conf.HTTPPort, router)
}
