package main

import (
	"log"

	"github.com/St0rmPetrel/Golang/jun_test_1/api"
	"github.com/St0rmPetrel/Golang/jun_test_1/db"
)

func main() {
	rdb, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	api.Up(rdb)
}
