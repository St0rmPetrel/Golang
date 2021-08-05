package main

import (
	"flag"
	"log"

	"github.com/St0rmPetrel/Golang/day_01/ex01/db"
)

func main() {
	name, err_flag := flag_init()
	log.SetFlags(0)
	if err_flag != nil {
		log.Fatal(err_flag)
	}
	db := db.NewDB()
	err_load := db.LoadData(name)
	if err_load != nil {
		log.Fatal(err_load)
	}
	db.PrintData()
}

func flag_init() (string, error) {
	var name string

	flag.StringVar(&name, "f", "", "file for load data")
	flag.Parse()
	if (name == "") || (flag.NArg() != 0) {
		return name, &BadFlagError{}
	}
	return name, nil
}

type BadFlagError struct {
}

func (err *BadFlagError) Error() string {
	return "Flag initialization error"
}
