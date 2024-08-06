package handler

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type handlerHttp struct {
	service coffee.UseCase
}

func NewHandler(s coffee.UseCase) *handlerHttp {
	return &handlerHttp{
		service: s,
	}
}

func (h *handlerHttp) GetBestTypeCoffee(w http.ResponseWriter, r *http.Request, params coffee.GetBestTypeCoffeeParams) {
	bestCoffees, err := h.service.GetBestCoffees(r.Context(), coffee.Filter{
		Characteristics: params.Characteristics,
	})

	if err != nil {
		render.JSON(w, r, coffee.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	render.JSON(w, r, coffee.BestCoffees{
		Characteristics: bestCoffees.Characteristics,
		Database:        bestCoffees.Database,
		Gemini:          bestCoffees.Gemini,
		Disclaimer:      bestCoffees.Disclaimer,
	})
}
