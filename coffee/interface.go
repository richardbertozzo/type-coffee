package coffee

import (
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

// Coffee represents the Coffee entity with all characteristics
type Coffee struct {
	UUID           string
	Name           string
	Image          string
	Description    string
	Caracteristics []string
}

// Repository the methods to write and read of Coffee (database abstraction)
type Repository interface {
	Reader
	Writer
}

// Reader the reader methods to get coffee data
type Reader interface {
	GetByID(id string) (entity.Coffee, error)
	ListByCaracteristic(entity.Caracteristic) ([]entity.Coffee, error)
}

// Writer the writer methods of coffee database
type Writer interface {
	Save(entity.Coffee) error
}

// UseCase the usecases/business rules layer
type UseCase interface {
	GetByID(id string) (Coffee, error)
	Create(Coffee) error
	ListByCaracteristic(c string) ([]Coffee, error)
}
