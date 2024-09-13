package config

import (
	"SebStudy/internal/infrastructure/mongodb"
	"SebStudy/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	Env              string
	ImageUploadPath  string
	NatsUrl          string
	ServerPort       string
	Logger           *logger.Config
	GRPC             *GRPC
	Mongo            *mongodb.Config
	MongoCollections MongoCollections
}

type GRPC struct {
	Port string
}

type MongoCollections struct {
	Orders string
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

	mongo := &mongodb.Config{
		URI:      getEnv("URI", "mongodb://localhost:27017"),
		Username: getEnv("Username", "admin"),
		Password: getEnv("Password", "admin"),
		Db:       "resume",
	}

	return &Config{
		ImageUploadPath: getEnv("IMAGE_UPLOAD_PATH", "./uploads"),
		NatsUrl:         getEnv("NATS_URL", "nats://127.0.0.1:4222"),
		ServerPort:      getEnv("SERVER_PORT", "50051"),
		Logger:          loggerConfig,
		GRPC:            grpc,
		Mongo:           mongo,
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}

	return val
}
