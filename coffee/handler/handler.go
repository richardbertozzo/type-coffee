package handler

import (
	"net/http"

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
	//TODO implement me
	panic("implement me")
}
