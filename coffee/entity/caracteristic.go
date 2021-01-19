package entity

import (
	"errors"
	"strings"
)

type Caracteristic struct {
	Name string
}

var (
	Strong = Caracteristic{
		Name: "forte",
	}
	Weak = Caracteristic{
		Name: "fraco",
	}
	Suable = Caracteristic{
		Name: "suave",
	}
	FullBodied = Caracteristic{
		Name: "encorpado",
	}
	Sweat = Caracteristic{
		Name: "doce",
	}
	Creamy = Caracteristic{
		Name: "cremoso",
	}
	Unknown = Caracteristic{
		Name: "desconhecido",
	}
)

func newCaracteristic(s string) (Caracteristic, error) {
	switch s {
	case strings.ToLower("forte"):
		return Strong, nil
	case strings.ToLower("fraco"):
		return Weak, nil
	case strings.ToLower("suave"):
		return Suable, nil
	case strings.ToLower("encorpado"):
		return FullBodied, nil
	case strings.ToLower("doce"):
		return Sweat, nil
	case strings.ToLower("cremoso"):
		return Creamy, nil
	default:
		return Unknown, errors.New("Caracteristic is unknown")
	}
}

func NewCaracteristics(caracStr []string) ([]Caracteristic, error) {
	if len(caracStr) <= 0 {
		return []Caracteristic{}, errors.New("Caracterics must be greater than zero")
	}

	caracs := []Caracteristic{}
	errs := []error{}
	for _, c := range caracStr {
		carac, err := newCaracteristic(c)
		caracs = append(caracs, carac)
		errs = append(errs, err)
	}

	return caracs, errs[0]
}
