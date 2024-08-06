package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	openapiMiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/handler"
	"github.com/richardbertozzo/type-coffee/coffee/provider"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

const defaultEnvFilePath = ".env"

type config struct {
	Port       string
	ChatGPTKey string
	DBCfg      dbConfig
}

type dbConfig struct {
	DBUrl      string
	DBDatabase string
	DBUser     string
	DBPassword string
}

func loadConfig() config {
	isDocker := "IS_DOCKER"
	portKey := "PORT"
	chatGPTKey := "CHAT_GPT_KEY"
	dbUrlKey := "DB_URL"
	dbDatabaseKey := "DB_DATABASE"
	dbUsernameKey := "DB_USERNAME"
	dbPwdKey := "DB_PASSWORD"

	isDockerEnv := os.Getenv(isDocker)
	err := godotenv.Load(defaultEnvFilePath)
	if err != nil && isDockerEnv == "" {
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

	return config{
		Port:       port,
		ChatGPTKey: chatGptValue,

		DBCfg: dbConfig{
			DBUrl:      os.Getenv(dbUrlKey),
			DBDatabase: os.Getenv(dbDatabaseKey),
			DBUser:     os.Getenv(dbUsernameKey),
			DBPassword: os.Getenv(dbPwdKey),
		},
	}
}

func main() {
	cfg := loadConfig()

	var db coffee.Service
	if cfg.DBCfg.DBUrl != "" {
		log.Println("database mode service enabled")
		dbURL := database.BuildURL(cfg.DBCfg.DBUrl, cfg.DBCfg.DBDatabase, cfg.DBCfg.DBDatabase, cfg.DBCfg.DBPassword)

		dbPool, err := database.NewConnection(context.Background(), dbURL)
		if err != nil {
			log.Fatal(err)
		}
		db = provider.NewDatabase(dbPool)
	}

	swagger, err := coffee.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	swagger.Servers = nil

	chatGPT, err := provider.NewChatGPTProvider(cfg.ChatGPTKey)
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(chatGPT, db)
	h := handler.NewHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(openapiMiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&openapiMiddleware.Options{
			Options: openapi3filter.Options{
				ExcludeRequestBody:    true,
				ExcludeResponseBody:   true,
				IncludeResponseStatus: true,
				MultiError:            false,
			},
		},
	))

	coffee.HandlerFromMux(h, r)

	log.Printf("http server runs on %s \n", cfg.Port)
	err = http.ListenAndServe(cfg.Port, r)
	if err != nil {
		log.Fatalf("error on start http server %v", err)
	}
}
