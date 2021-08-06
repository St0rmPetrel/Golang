package main

import (
	"flag"
	"log"

	"github.com/St0rmPetrel/Golang/day_01/ex01/db"
)

func main() {
	name_old, name_new, err_flag := flag_init()
	log.SetFlags(0)
	if err_flag != nil {
		log.Fatal(err_flag)
	}
	if name_old == name_new {
		return
	}
	db_new := db.NewDB()
	db_old := db.NewDB()
	err_load_old := db_old.LoadData(name_old)
	if err_load_old != nil {
		log.Fatal(err_load_old)
	}
	err_load_new := db_new.LoadData(name_new)
	if err_load_new != nil {
		log.Fatal(err_load_new)
	}
	db_old.Compare(*db_new)
}

func flag_init() (string, string, error) {
	var name_old, name_new string

	flag.StringVar(&name_old, "old", "", "file for load data")
	flag.StringVar(&name_new, "new", "", "file for load data")
	flag.Parse()
	if name_old == "" || name_new == "" || flag.NArg() != 0 {
		return name_old, name_new, &BadFlagError{}
	}
	return name_old, name_new, nil
}

type BadFlagError struct {
}

func (err *BadFlagError) Error() string {
	return "Flag initialization error"
}
