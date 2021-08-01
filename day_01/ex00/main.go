package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Cakes struct {
	Cake []ACake
}

type ACake struct {
	Name        string
	Time        string
	Ingredients []Ingredient
}

type Ingredient struct {
	Ingredient_name  string
	Ingredient_count string
	Ingredient_unit  string
}

func main() {
	jsonFile, err := os.Open("cakes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cakes Cakes
	json.Unmarshal(byteValue, &cakes)
	fmt.Println(cakes)
}
