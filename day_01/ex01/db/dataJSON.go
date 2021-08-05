package db

import (
	"encoding/json"
)

type DataJSON struct {
	Data []byte
}

func (d DataJSON) Read() (*Recipes, error) {
	recipesPtr := new(Recipes)
	err := json.Unmarshal(d.Data, recipesPtr)
	return recipesPtr, err
}
