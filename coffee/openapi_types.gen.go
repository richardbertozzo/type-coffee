// Package coffee provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package coffee

// Defines values for Characteristic.
const (
	Acidity    Characteristic = "acidity"
	Aftertaste Characteristic = "aftertaste"
	Aroma      Characteristic = "aroma"
	Body       Characteristic = "body"
	Flavor     Characteristic = "flavor"
	Sweetness  Characteristic = "sweetness"
)

// BestCoffees defines model for BestCoffees.
type BestCoffees struct {
	Characteristics []Characteristic `json:"characteristics"`
	ChatGpt         []Option         `json:"chat_gpt"`
	Database        *[]Option        `json:"database,omitempty"`
}

// Characteristic defines model for Characteristic.
type Characteristic string

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Option defines model for Option.
type Option struct {
	Details *map[string]interface{} `json:"details,omitempty"`
	Message string                  `json:"message"`
}

// GetBestTypeCoffeeParams defines parameters for GetBestTypeCoffee.
type GetBestTypeCoffeeParams struct {
	// Characteristics Characteristics to build your best type of coffee, up to 3 selected characteristics.
	Characteristics []Characteristic `form:"characteristics" json:"characteristics"`
}
