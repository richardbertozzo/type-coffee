package usecase

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type useCase struct {
	chatGPTProvider  coffee.Service
	dbService        coffee.Service
	dbServiceEnabled bool
}

// New returns a new coffee use case
func New(chatGPTProvider, dbService coffee.Service) *useCase {
	return &useCase{
		chatGPTProvider:  chatGPTProvider,
		dbService:        dbService,
		dbServiceEnabled: dbService != nil,
	}
}

func (u useCase) GetBestCoffees(ctx context.Context, filter coffee.Filter) (*coffee.BestCoffees, error) {
	lenChar := len(filter.Characteristics)
	if lenChar == 0 {
		return nil, errors.New("must have minimum 1 input characteristic")
	} else if lenChar > 3 {
		return nil, errors.New("must have up to 3 input characteristics")
	}

	filter.Limit = setDefaultLimit(filter.Limit)

	g, ctx := errgroup.WithContext(ctx)

	var chatGptOpts []coffee.Option
	g.Go(func() error {
		opts, err := u.chatGPTProvider.GetCoffeeOptionsByCharacteristics(ctx, filter)
		if err != nil {
			return err
		}

		for _, opt := range opts {
			chatGptOpts = append(chatGptOpts, coffee.Option{
				Message: opt.Message,
			})
		}

		return nil
	})

	var dbOpts []coffee.Option
	if u.dbServiceEnabled {
		g.Go(func() error {
			opts, err := u.dbService.GetCoffeeOptionsByCharacteristics(ctx, filter)
			if err != nil {
				return err
			}

			for _, opt := range opts {
				dbOpts = append(dbOpts, coffee.Option{
					Message: opt.Message,
				})
			}

			return nil
		})

	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &coffee.BestCoffees{
		Characteristics: filter.Characteristics,
		ChatGpt:         chatGptOpts,
		Database:        &dbOpts,
	}, nil
}

func setDefaultLimit(limit int) int {
	defaultLimit := 5
	if limit == 0 || limit > 5 {
		return defaultLimit
	}
	return limit
}
