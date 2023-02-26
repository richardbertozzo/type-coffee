package usecase

import (
	"context"
	"errors"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type useCase struct {
	chatGPTProvider coffee.Provider
}

// NewService returns a new coffee use case
func NewService(chatGPTProvider coffee.Provider) coffee.UseCase {
	return useCase{
		chatGPTProvider: chatGPTProvider,
	}
}

func (u useCase) GetBestCoffees(ctx context.Context, filter coffee.Filter) (*coffee.BestCoffees, error) {
	lenChar := len(filter.Characteristics)
	if lenChar == 0 {
		return nil, errors.New("must have minimum 1 characteristic input")
	} else if lenChar > 3 {
		return nil, errors.New("must have up to 3 characteristic input")
	}

	opts, err := u.chatGPTProvider.GetCoffeeOptionsByCharacteristics(ctx, filter)
	if err != nil {
		return nil, err
	}

	options := make([]coffee.Option, len(opts))
	for i, opt := range opts {
		options[i] = coffee.Option{
			Message: opt.Message,
		}
	}

	return &coffee.BestCoffees{
		Characteristics: filter.Characteristics,
		ChatGpt:         options,
		Database:        nil,
	}, nil
}
