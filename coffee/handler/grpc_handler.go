package handler

import (
	"context"

	"github.com/richardbertozzo/type-coffee/coffee"
	pb "github.com/richardbertozzo/type-coffee/coffee/api"
)

type handlerGrpc struct {
	service coffee.UseCase
	pb.UnimplementedCoffeeServiceServer
}

func NewGrpcHandler(service coffee.UseCase) pb.CoffeeServiceServer {
	return &handlerGrpc{service: service}
}

func (s *handlerGrpc) GetBestTypeCoffee(ctx context.Context, req *pb.GetBestTypeCoffeeRequest) (*pb.GetBestTypeCoffeeResponse, error) {
	bestCoffees, err := s.service.GetBestCoffees(ctx, coffee.Filter{
		Characteristics: convertCharacteristics(req.Characteristics),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetBestTypeCoffeeResponse{
		BestCoffee: &pb.BestCoffee{
			Gemini:     convertOptions(*bestCoffees.Gemini),
			Database:   convertOptions(*bestCoffees.Database),
			Disclaimer: bestCoffees.Disclaimer,
		},
	}, nil
}

func convertCharacteristics(pbCharacs []pb.Characteristic) []coffee.Characteristic {
	cs := make([]coffee.Characteristic, len(pbCharacs))
	for i, c := range pbCharacs {
		cs[i] = mapCharacteristic(c)
	}
	return cs
}

func convertOptions(opts []coffee.Option) []*pb.Option {
	if len(opts) == 0 {
		return nil
	}

	os := make([]*pb.Option, len(opts))
	for i, opt := range opts {
		var d map[string]string
		if opt.Details != nil && len(*opt.Details) > 0 {
			d = make(map[string]string, len(*opt.Details))
			for key, val := range *opt.Details {
				d[key] = val.(string)
			}
		}

		os[i] = &pb.Option{
			Message: opt.Message,
			Details: d,
		}
	}

	return os
}

func mapCharacteristic(c pb.Characteristic) coffee.Characteristic {
	switch c {
	case pb.Characteristic_ACIDITY:
		return coffee.Acidity
	case pb.Characteristic_AFTERTASTE:
		return coffee.Aftertaste
	case pb.Characteristic_AROMA:
		return coffee.Aroma
	case pb.Characteristic_BODY:
		return coffee.Body
	case pb.Characteristic_FLAVOR:
		return coffee.Flavor
	case pb.Characteristic_SWEETNESS:
		return coffee.Sweetness
	default:
		return coffee.Flavor
	}
}
