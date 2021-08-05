package db

import (
	"encoding/xml"
)

type DataXML struct {
	Data []byte
}

func (d DataXML) Read() (*Recipes, error) {
	recipesPtr := new(Recipes)
	err := xml.Unmarshal(d.Data, recipesPtr)
	return recipesPtr, err
}
