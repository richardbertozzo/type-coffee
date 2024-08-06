package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const defaultEnvFilePath = ".env"

type Config struct {
	Port       string
	ChatGPTKey string
	DBCfg      DBConfig
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
	chatGPTKey := "CHAT_GPT_KEY"
	dbUrlKey := "DB_URL"
	dbDatabaseKey := "DB_DATABASE"
	dbUsernameKey := "DB_USERNAME"
	dbPwdKey := "DB_PASSWORD"

	err := godotenv.Load(getEnvPath(pathEnv))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv(portKey)
	if port == "" {
		port = ":3000"
	}

	chatGptValue := os.Getenv(chatGPTKey)
	if chatGptValue == "" {
		log.Fatal("CHAT_GPT_KEY ENV is required")
	}

	return Config{
		Port:       port,
		ChatGPTKey: chatGptValue,

		DBCfg: DBConfig{
			DBUrl:      os.Getenv(dbUrlKey),
			DBDatabase: os.Getenv(dbDatabaseKey),
			DBUser:     os.Getenv(dbUsernameKey),
			DBPassword: os.Getenv(dbPwdKey),
		},
	}
}
