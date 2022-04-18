package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/richardbertozzo/type-coffee/coffee/handler"
	"github.com/richardbertozzo/type-coffee/coffee/repository"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	db := repository.NewMemoryDB()
	uc := usecase.NewService(db)
	h := handler.NewHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/coffee", h.ListCoffee)
	r.Get("/coffee/{id}", h.GetCoffeeByID)
	r.Post("/coffee", h.CreateCoffee)

	log.Printf("http server runs on %s \n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("error on start http server %v", err)
	}
}
