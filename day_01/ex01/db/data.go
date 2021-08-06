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
	for _, ingrd := range src.Ingredient {
		if is_in_cake, id := d.find(ingrd); is_in_cake  {
			d.Ingredient[id].Compare(ingrd, src)
		} else {
			fmt.Printf("ADDED ingredient \"%s\"\n", ingrd.Name)
		}
	}
	for _, ingrd := range d.Ingredient {
		if is_in_cake, _ := src.find(ingrd); !is_in_cake {
			fmt.Printf("REMOVED ingredient \"%s\"\n", ingrd.Name)
		}
	}
}

func (d Cake) find(key Ingredient) (bool, int) {
	id := sort.Search(len(d.Ingredient), func(i int) bool {
		return d.Ingredient[i].Name >= key.Name
	})
	if id < 0 || key.Name != d.Ingredient[id].Name {
		return false, 0
	}
	return true, id
}

func (d Ingredient) Compare(src Ingredient, c Cake) {
	if d.Unit == src.Unit {
		if d.Count != src.Count {
			fmt.Print("CHANGED unit count for ingredient ")
			fmt.Printf("\"%s\" for cake \"%s\" - ", d.Name, c.Name)
			fmt.Printf("\"%s\" instead of \"%s\"\n", d.Count, src.Count)
		}
	} else {
		fmt.Print("CHANGED unit for ingredient ")
		fmt.Printf("\"%s\" for cake \"%s\" - ", d.Name, c.Name)
		fmt.Printf("\"%s\" instead of \"%s\"\n", d.Unit, src.Unit)
	}
}
