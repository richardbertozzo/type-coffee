package entity

import (
	"errors"
	"fmt"
)

var (
	errNameBlank = errors.New("Name is blank")
)

// Coffee represent the entity of a Coffee
type Coffee struct {
	UUID           string
	Name           string
	Description    string
	caracteristics []Caracteristic
}

func (c Coffee) String() string {
	return fmt.Sprintf("ID: %s - Name: %s", c.UUID, c.Name)
}

func (c Coffee) IsZero() bool {
	return c.UUID == ""
}

func New(uuid, name string, c []Caracteristic) (Coffee, error) {
	if name == "" {
		return Coffee{}, errNameBlank
	} else if uuid == "" {
		return Coffee{}, errors.New("UUID is null")
	}

	return Coffee{
		UUID:           uuid,
		Name:           name,
		caracteristics: c,
	}, nil
}
