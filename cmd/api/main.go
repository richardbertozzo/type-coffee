package main

import (
	"context"
	"log"
	"net/http"

	openapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/handler"
	"github.com/richardbertozzo/type-coffee/coffee/provider"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

type config struct {
	Port       string
	DBUrl      string
	ChatGPTKey string
}

func loadConfig(filePath string) config {
	portKey := "PORT"
	dbKey := "DATABASE_URL"
	chatGPTKey := "CHAT_GPT_KEY"

	viper.SetConfigFile(filePath)
	viper.SetDefault(portKey, ":3000")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalf("config not found: %s", err)
		} else {
			log.Fatalf("error in read config: %s", err)
		}
	}

	chatGptKey := viper.GetString(chatGPTKey)
	if chatGptKey == "" {
		log.Fatal("CHAT_GPT_KEY ENV is required")
	}

	return config{
		Port:       viper.GetString(portKey),
		DBUrl:      viper.GetString(dbKey),
		ChatGPTKey: viper.GetString(chatGPTKey),
	}
}

func main() {
	cfg := loadConfig(".env")

	var db coffee.Service
	if cfg.DBUrl != "" {
		log.Println("database mode service enabled")
		dbPool, err := database.NewConnection(context.Background(), cfg.DBUrl)
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
