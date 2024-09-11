package config

import (
	"SebStudy/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	Env             string
	ImageUploadPath string
	NatsUrl         string
	ServerPort      string
	Logger          *logger.Config
	GRPC            *GRPC
}

type GRPC struct {
	Port string
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env variables: " + err.Error())
	}

	devMode, err := strconv.ParseBool(getEnv("DEV_MODE", "false"))
	if err != nil {
		devMode = true
	}

	loggerConfig := &logger.Config{
		LogLevel: getEnv("LOG_LEVEL", "debug"),
		DevMode:  devMode,
	}

	grpc := &GRPC{
		Port: getEnv("GRPC_PORT", ":50051"),
	}

	return &Config{
		ImageUploadPath: getEnv("IMAGE_UPLOAD_PATH", "./uploads"),
		NatsUrl:         getEnv("NATS_URL", "nats://127.0.0.1:4222"),
		ServerPort:      getEnv("SERVER_PORT", "50051"),
		Logger:          loggerConfig,
		GRPC:            grpc,
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}

	return val
}
