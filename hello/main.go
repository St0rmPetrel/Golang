package main

import (
	"fmt"
	"log"
	"github.com/St0rmPetrel/Golang/hello/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Telman")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
