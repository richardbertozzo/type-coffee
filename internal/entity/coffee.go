package entity

// Coffee represent the entity of a Coffee
type Coffee struct {
	UUID          string   `json:"uuid"`
	Name          string   `json:"name"`
	Caracteristic []string `json:"caracteristics"`
}
