package db

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
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
	db.Data.Sort()
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

func (db DataBase) Compare(src DataBase) {
	for _, cake := range src.Data.Cake {
		if is_in_db, id := db.find(cake); is_in_db  {
			db.Data.Cake[id].Compare(cake)
		} else {
			fmt.Printf("ADDED cake \"%s\"\n", cake.Name)
		}
	}
	for _, cake := range db.Data.Cake {
		if is_in_db, _ := src.find(cake); !is_in_db {
			fmt.Printf("REMOVED cake \"%s\"\n", cake.Name)
		}
	}
}

func (db DataBase) find(key Cake) (bool, int) {
	cake_id := sort.Search(len(db.Data.Cake), func(i int) bool {
		return db.Data.Cake[i].Name >= key.Name
	})
	if cake_id < 0 || key.Name != db.Data.Cake[cake_id].Name {
		return false, 0
	}
	return true, cake_id
}
