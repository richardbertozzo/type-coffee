package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/richardbertozzo/type-coffee/coffee"
)

type handlerHTTP struct {
	service coffee.UseCase
}

// NewHandler creates the handlers HTTP giving the coffee use cases
func NewHandler(s coffee.UseCase) *handlerHTTP {
	return &handlerHTTP{
		service: s,
	}
}

func (h *handlerHTTP) ListCoffee(w http.ResponseWriter, r *http.Request) {
	c, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, c)
}

func (h *handlerHTTP) GetCoffeeByID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	if param == "" {
		http.Error(w, "id must be not null", http.StatusBadRequest)
		return
	}

	c, err := h.service.GetByID(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, c)
}

func (h *handlerHTTP) CreateCoffee(w http.ResponseWriter, r *http.Request) {

}
