package config

import (
	"os"
	"resume-server/internal/infrastructure/mongodb"
	"resume-server/logger"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env                        string
	Logger                     *logger.Config
	GRPC                       *GRPC
	Mongo                      *mongodb.Config
	MongoCollections           *MongoCollections
	AuthServerURL              string
	EventStoreConnectionString string
	EventStorePrefix           string
	EventStoreGroupName        string
	PublicKeyPath              string
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
		Logger:                     loggerConfig,
		GRPC:                       grpc,
		Mongo:                      mongo,
		MongoCollections:           mongoCollections,
		AuthServerURL:              getEnv("AUTH_SERVER_URL", "sso-service:50052"),
		EventStoreConnectionString: getEnv("EVENT_STORE_CONN", "esdb://eventstore.db:2113?tls=false"),
		EventStorePrefix:           getEnv("EVENT_STORE_PREFIX", "heychar-resume-"),
		EventStoreGroupName:        getEnv("EVENT_STORE_GROUP_NAME", "resume"),
		PublicKeyPath:              getEnv("PUBLIC_KEY_PATH", "certs/jwtRSA256-public.pem"),
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}

	return val
}
