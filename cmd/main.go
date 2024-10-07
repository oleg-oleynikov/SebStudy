package main

import (
	"resume-server/config"
	"resume-server/logger"
	"resume-server/server"
)

func main() {
	cfg := config.InitConfig()

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.Fatal(server.NewServer(cfg, appLogger).Run())
}
