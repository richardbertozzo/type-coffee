package usecase

import (
	"errors"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

type useCase struct {
	db coffee.Repository
}

func NewService(db coffee.Repository) coffee.UseCase {
	return useCase{}
}

func (u useCase) GetByID(id string) (entity.Coffee, error) {
	if id == "" {
		return entity.Coffee{}, errors.New("id must be not blank")
	}

	return u.db.GetByID(id)
}

func (u useCase) Insert(c entity.Coffee) error {
	return u.db.Save(c)
}
