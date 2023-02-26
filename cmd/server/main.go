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
	"github.com/richardbertozzo/type-coffee/coffee/repository"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	swagger, err := coffee.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	db := repository.NewMemoryDB()
	uc := usecase.NewService(db)
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
