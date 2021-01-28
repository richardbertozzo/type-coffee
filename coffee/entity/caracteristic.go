package entity

import (
	"errors"
	"strings"
)

type Caracteristic struct {
	name string
}

func (c Caracteristic) String() string {
	return c.name
}

var (
	Strong = Caracteristic{
		name: "forte",
	}
	Weak = Caracteristic{
		name: "fraco",
	}
	Suable = Caracteristic{
		name: "suave",
	}
	FullBodied = Caracteristic{
		name: "encorpado",
	}
	Sweat = Caracteristic{
		name: "doce",
	}
	Creamy = Caracteristic{
		name: "cremoso",
	}
	Unknown = Caracteristic{
		name: "desconhecido",
	}
)

func NewCaracteristic(s string) (Caracteristic, error) {
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
		carac, err := NewCaracteristic(c)
		caracs = append(caracs, carac)
		errs = append(errs, err)
	}

	return caracs, errs[0]
}
