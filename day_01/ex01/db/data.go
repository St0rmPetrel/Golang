package db

import (
	"encoding/xml"
	"sort"
	"fmt"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []Cake   `xml:"cake" json:"cake"`
}

type Cake struct {
	Name       string       `xml:"name" json:"name"`
	Time       string       `xml:"stovetime" json:"time"`
	Ingredient []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit"`
}

func (d *Recipes) Sort() {
	for _, cake := range d.Cake {
		sort.Sort(&cake)
	}
	sort.Sort(d)
}

func (d Recipes) Len() int {
	return len(d.Cake)
}

func (d Recipes) Less(i, j int) bool {
	return d.Cake[i].Name < d.Cake[j].Name
}

func (d *Recipes) Swap(i, j int) {
	tmp := d.Cake[i]
	d.Cake[i] = d.Cake[j]
	d.Cake[j] = tmp
}

func (d Cake) Len() int {
	return len(d.Ingredient)
}

func (d Cake) Less(i, j int) bool {
	return d.Ingredient[i].Name < d.Ingredient[j].Name
}

func (d *Cake) Swap(i, j int) {
	tmp := d.Ingredient[i]
	d.Ingredient[i] = d.Ingredient[j]
	d.Ingredient[j] = tmp
}

func (d Cake) Compare(src Cake) {
	if d.Time != src.Time {
		fmt.Print("CHANGED cooking time for cake ")
		fmt.Printf("\"%s\" - \"%s\" instead of \"%s\"\n", d.Name, d.Time, src.Time)
	}
//	for _,  := range src.Data.Cake {
//		if is_in_db, id := db.find(cake); is_in_db  {
//			db.Data.Cake[id].Compare(cake)
//		} else {
//			fmt.Printf("ADDED cake \"%s\"\n", cake.Name)
//		}
//	}
//	for _, cake := range db.Data.Cake {
//		if is_in_db, _ := src.find(cake); !is_in_db {
//			fmt.Printf("REMOVED cake \"%s\"\n", cake.Name)
//		}
//	}
}

//func (db ) find(key Cake) bool {
//	cake_id := sort.Search(len(db.Data.Cake), func(i int) bool {
//		return db.Data.Cake[i].Name >= key.Name
//	})
//	if cake_id < 0 || key.Name != db.Data.Cake[cake_id].Name {
//		return false, 0
//	}
//	return true, cake_id
//	return true
//}
