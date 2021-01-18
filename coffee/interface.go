package coffee

import "github.com/richardbertozzo/type-coffee/coffee/entity"

type Repository interface {
	GetByID(id string) (entity.Coffee, error)
	Save(entity.Coffee) error
}

type UseCase interface {
	GetByID(id string) (entity.Coffee, error)
	Insert(entity.Coffee) error
}
