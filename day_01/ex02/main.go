package main

import (
	"flag"
	"log"
	"os"
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
	file_old, err_open_old := os.Open(name_old)
	if err_open_old != nil {
		log.Fatal(err_open_old)
	}
	defer file_old.Close()
	file_new, err_open_new := os.Open(name_new)
	if err_open_new != nil {
		log.Fatal(err_open_new)
	}
	defer file_new.Close()
	err := diff(file_old, file_new)
	if err != nil {
		log.Fatal(err)
	}
}

func flag_init() (string, string, error) {
	var name_old, name_new string

	flag.StringVar(&name_old, "old", "", "file for compare")
	flag.StringVar(&name_new, "new", "", "file for compare")
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
