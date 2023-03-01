package main

import (
	"context"
	"log"
	"net/http"
	"os"

	openapiMiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/handler"
	"github.com/richardbertozzo/type-coffee/coffee/service"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	chatGptKey := os.Getenv("CHAT_GPT_KEY")
	if chatGptKey == "" {
		log.Fatal("CHAT_GPT_KEY ENV is required")
	}

	var dbService coffee.Service
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("database mode service enabled")
		dbPool, err := database.NewConnection(context.Background(), dbURL)
		if err != nil {
			log.Fatal(err)
		}
		dbService = service.NewDatabaseService(dbPool)
	}

	swagger, err := coffee.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	swagger.Servers = nil

	chatGPTService, err := service.NewChatGPTProvider(chatGptKey)
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(chatGPTService, dbService)
	h := handler.NewHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(openapiMiddleware.OapiRequestValidator(swagger))

	coffee.HandlerFromMux(h, r)

	log.Printf("http server runs on %s \n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("error on start http server %v", err)
	}
}
