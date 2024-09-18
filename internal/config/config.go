package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const defaultEnvFilePath = ".env"

type Config struct {
	Port         string
	GRPCPort     string
	GRPCEnabled  bool
	GeminiAPIKey string
	DBCfg        DBConfig
}

type DBConfig struct {
	DBUrl      string
	DBDatabase string
	DBUser     string
	DBPassword string
}

func getEnvPath(path string) string {
	if path != "" {
		return path
	}

	return defaultEnvFilePath
}

func LoadConfig() Config {
	pathEnv := os.Getenv("CONFIG_PATH")
	portKey := "PORT"
	gRPCPortKey := "GRPC_PORT"
	grpcEnabledKey := "GRPC_ENABLED"
	geminiAPIKey := "GEMINI_API_KEY"
	dbUrlKey := "DB_URL"
	dbDatabaseKey := "DB_DATABASE"
	dbUsernameKey := "DB_USERNAME"
	dbPwdKey := "DB_PASSWORD"

	err := godotenv.Load(getEnvPath(pathEnv))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := getEnvWithDefault(portKey, ":3000")
	grpcPort := getEnvWithDefault(gRPCPortKey, ":9000")

	grpcEnabledStr := os.Getenv(grpcEnabledKey)
	grpcEnabled, err := strconv.ParseBool(grpcEnabledStr)
	if err != nil {
		log.Fatal("GRPC_ENABLED ENV is not boolean type")
	}

	geminiAPIValue := os.Getenv(geminiAPIKey)
	if geminiAPIValue == "" {
		log.Fatal("GEMINI_API_KEY ENV is required")
	}

	return Config{
		Port:         port,
		GRPCPort:     grpcPort,
		GRPCEnabled:  grpcEnabled,
		GeminiAPIKey: geminiAPIValue,

		DBCfg: DBConfig{
			DBUrl:      os.Getenv(dbUrlKey),
			DBDatabase: os.Getenv(dbDatabaseKey),
			DBUser:     os.Getenv(dbUsernameKey),
			DBPassword: os.Getenv(dbPwdKey),
		},
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
