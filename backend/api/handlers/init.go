package handlers

import (
	"os"
	"sade-backend/config"
	"sade-backend/db/cmd"
	"sade-backend/pkg/auth"
	"sade-backend/pkg/ffmpeg"
	"sade-backend/pkg/logger"
	"sade-backend/pkg/notifier"
	"time"
)

var Log *logger.Logger
var MediaTable *cmd.DataTable
var TransactionTable *cmd.DataTable
var A *auth.Auth
var Fw *ffmpeg.FFMpeg
var N *notifier.Notifier
var StripeKey string
var OriginPort string

func init() {
	Log = logger.New("../log/handler_log.log")
	configData, err := os.ReadFile("../config.yaml")
	if err != nil {
		Log.Fatalf("Error reading config file: %s", err)
	}
	conf, err := config.New(configData, "../.env")
	if err != nil {
		Log.Fatalf("Error loading config: %s", err)
	}
	StripeKey = conf.Gateway.ApiKey
	OriginPort = conf.OriginURL

	//DataTable initialization
	userTable := cmd.New("users")
	linkTable := cmd.New("links")
	TransactionTable = cmd.New("transaction")
	MediaTable = cmd.New("media")

	//Auth,  Notifier and FFMPeg initialization
	A = auth.New(userTable, linkTable, time.Hour)
	N = notifier.New(conf.Notifier.Host, conf.Notifier.APIKey)
	Fw = ffmpeg.New(conf.Ffmpeg.AudioWatermark, conf.Ffmpeg.VideoWatermark, conf.Ffmpeg.ImgWatermark)
}
