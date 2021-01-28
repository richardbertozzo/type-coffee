package entity

import (
	"errors"
	"fmt"
)

// Coffee represent the entity of a Coffee
type Coffee struct {
	UUID           string
	Name           string
	Image          Link
	Description    string
	caracteristics []Caracteristic
}

func (c Coffee) String() string {
	return fmt.Sprintf("ID: %s - Name: %s", c.UUID, c.Name)
}

func (c Coffee) IsZero() bool {
	return c.UUID == ""
}

func (c Coffee) Caracteristics() (caracs []string) {
	for _, c := range c.caracteristics {
		caracs = append(caracs, c.String())
	}
	return
}

func New(uuid, name, d string, l Link, c []Caracteristic) (Coffee, error) {
	if name == "" {
		return Coffee{}, errors.New("Name is blank")
	} else if uuid == "" {
		return Coffee{}, errors.New("UUID is null")
	} else if len(c) < 1 || len(c) > 4 {
		return Coffee{}, errors.New("Coffee must has more than 0 and less than 4 caracteristic")
	}

	return Coffee{
		UUID:           uuid,
		Name:           name,
		Image:          l,
		Description:    d,
		caracteristics: c,
	}, nil
}
