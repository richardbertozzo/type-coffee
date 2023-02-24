package coffee

// Characteristic represent the entity of a coffee characteristic, example: strong
type Characteristic string

func (c Characteristic) String() string {
	return string(c)
}

const (
	Aroma      = "aroma"
	Flavor     = "flavor"
	Aftertaste = "aftertaste"
	Acidity    = "acidity"
	Body       = "body"
)

func (c Characteristic) IsValid() bool {
	switch c {
	case Aroma, Flavor, Aftertaste, Acidity, Body:
		return true
	}

	return false
}
