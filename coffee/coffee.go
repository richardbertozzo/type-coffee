package coffee

import (
	"context"
)

// UseCase the use-cases/business rules layer
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

type Service interface {
	GetCoffeeOptionsByCharacteristics(context.Context, Filter) ([]OptionProvider, error)
}
