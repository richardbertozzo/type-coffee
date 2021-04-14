package entity

import (
	"errors"
	"strings"
)

// Caracteristic represent the entity of a coffee caracteristic, example: strong
type Caracteristic struct {
	name string
}

func (c Caracteristic) String() string {
	return c.name
}

var (
	strong = Caracteristic{
		name: "forte",
	}
	weak = Caracteristic{
		name: "fraco",
	}
	suable = Caracteristic{
		name: "suave",
	}
	fullBodied = Caracteristic{
		name: "encorpado",
	}
	sweat = Caracteristic{
		name: "doce",
	}
	creamy = Caracteristic{
		name: "cremoso",
	}
	unknown = Caracteristic{
		name: "desconhecido",
	}
)

// NewCaracteristic creates a Caracteristic
func NewCaracteristic(s string) (Caracteristic, error) {
	switch s {
	case strings.ToLower("forte"):
		return strong, nil
	case strings.ToLower("fraco"):
		return weak, nil
	case strings.ToLower("suave"):
		return suable, nil
	case strings.ToLower("encorpado"):
		return fullBodied, nil
	case strings.ToLower("doce"):
		return sweat, nil
	case strings.ToLower("cremoso"):
		return creamy, nil
	default:
		return unknown, errors.New("Caracteristic is unknown")
	}
}

// NewCaracteristics creates a slice of Caracteristics
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
