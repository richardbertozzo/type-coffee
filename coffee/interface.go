package coffee

import (
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

type Coffee struct {
	UUID           string
	Name           string
	Image          string
	Description    string
	Caracteristics []string
}

type Repository interface {
	GetByID(id string) (entity.Coffee, error)
	ListByCaracteristic(entity.Caracteristic) ([]entity.Coffee, error)
	Save(entity.Coffee) error
}

type UseCase interface {
	GetByID(id string) (Coffee, error)
	Create(Coffee) error
	ListByCaracteristic(c string) ([]Coffee, error)
}
