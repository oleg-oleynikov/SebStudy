package main

import (
	"SebStudy/config"
	"SebStudy/server"

	"SebStudy/logger"
)

func main() {
	cfg := config.InitConfig()

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.Fatal(server.NewServer(cfg, appLogger).Run())
}
