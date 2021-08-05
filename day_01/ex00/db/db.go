package db

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type DataBase struct {
	Data *Recipes
}

func NewDB() *DataBase {
	return new(DataBase)
}

func (db *DataBase) LoadData(name string) error {
	reader, err := GetDBReader(name)
	if err != nil {
		return err
	}
	db.Data, err = reader.Read()
	return err
}

func (db DataBase) PrintData() error {
	b, err := db.marshal()
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func (db DataBase) marshal() ([]byte, error) {
	if db.Data.XMLName.Local != "" {
		return json.MarshalIndent(db.Data, "", "    ")
	}
	return xml.MarshalIndent(db.Data, "", "    ")
}
