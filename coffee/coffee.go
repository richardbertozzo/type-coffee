package coffee

import (
	"context"

	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

// Repository the methods to write and read of Coffee (database abstraction)
type Repository interface {
	Reader
	Writer
}

// Reader the reader methods to get coffee data
type Reader interface {
	GetByID(id string) (entity.Coffee, error)
	List() ([]entity.Coffee, error)
	ListByCaracteristic(Characteristic) ([]entity.Coffee, error)
}

// Writer the writer methods of coffee database
type Writer interface {
	Save(entity.Coffee) error
}

// UseCase the usecases/business rules layer
type UseCase interface {
	GetBestCoffees(context.Context, Filter) (*BestCoffees, error)
}

// OptionProvider mean a selected option coffee
type OptionProvider struct {
	Message string
	Details interface{}
}

type Filter struct {
	Characteristics []Characteristic
	Limit           int
	Sort            bool // 0 = desc - 1 = asc
}

type Provider interface {
	GetCoffeeOptionsByCharacteristics(context.Context, Filter) ([]OptionProvider, error)
}
