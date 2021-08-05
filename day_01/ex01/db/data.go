package db

import (
	"encoding/xml"
	"sort"
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
