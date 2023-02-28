package main

import (
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

	swagger, err := coffee.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	swagger.Servers = nil

	provider, err := service.NewChatGPTProvider(chatGptKey)
	if err != nil {
		log.Fatal(err)
	}
	uc := usecase.NewUseCase(provider)
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
