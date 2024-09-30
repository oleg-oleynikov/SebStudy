package config

import (
	"SebStudy/internal/infrastructure/mongodb"
	"SebStudy/logger"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	ImageUploadPath  string
	NatsUrl          string
	Logger           *logger.Config
	GRPC             *GRPC
	Mongo            *mongodb.Config
	MongoCollections *MongoCollections
	AuthServerURL    string
}

type GRPC struct {
	Port string
}

type MongoCollections struct {
	Resumes string
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
		Port: getEnv("GRPC_PORT", ":50055"),
	}

	mongoConnTimeout, err := time.ParseDuration(getEnv("MONGO_CONNECT_TIMEOUT", "5s"))
	if err != nil {
		mongoConnTimeout = time.Second * 5
	}

	mongo := &mongodb.Config{
		URI:            getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Username:       getEnv("MONGO_USER", "admin"),
		Password:       getEnv("MONGO_PASSWORD", "admin"),
		Db:             getEnv("MONGO_DB", "resume"),
		ConnectTimeout: mongoConnTimeout,
	}

	mongoCollections := &MongoCollections{
		Resumes: getEnv("MONGO", "resumes"),
	}

	return &Config{
		ImageUploadPath:  getEnv("IMAGE_UPLOAD_PATH", "./uploads"),
		NatsUrl:          getEnv("NATS_URL", "nats://resume-nats:4225"),
		Logger:           loggerConfig,
		GRPC:             grpc,
		Mongo:            mongo,
		MongoCollections: mongoCollections,
		AuthServerURL:    getEnv("AUTH_SERVER_URL", "sso-service:50052"),
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}

	return val
}
