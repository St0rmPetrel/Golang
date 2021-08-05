package main

import (
	"fmt"
	"log"

	"github.com/St0rmPetrel/Golang/day_01/ex00/db"
)

func test_print(name string) {
	dbr, err := db.GetDBReader(name)
	if err != nil {
		log.Fatal(err)
	}
	cakes, _ := dbr.Read()
	fmt.Println(cakes)
}

func main() {
	test_print("data/cakes.json")
	test_print("data/cakes.xml")
}
