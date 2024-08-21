package initializers

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger(env string) {
	if env == "debug" {
		// log.SetReportCaller(true) // If you want see where from log
		log.SetLevel(log.DebugLevel)
	} else if env == "prod" {
		log.SetLevel(log.InfoLevel)
	}
}
