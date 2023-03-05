package coffee

func (c Characteristic) String() string {
	return string(c)
}

func (c Characteristic) IsValid() bool {
	switch c {
	case Acidity, Aftertaste, Aroma, Body, Flavor, Sweetness:
		return true
	}
	return false
}

func ListAllCharacteristic() []string {
	return []string{
		Acidity.String(),
		Aftertaste.String(),
		Aroma.String(),
		Body.String(),
		Flavor.String(),
		Sweetness.String(),
	}
}

func ConvertToCharacteristic(strs []string) []Characteristic {
	c := make([]Characteristic, len(strs))
	for i, str := range strs {
		c[i] = Characteristic(str)
	}

	return c
}
