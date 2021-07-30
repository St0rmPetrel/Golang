package main

import (
	"fmt"
	"github.com/St0rmPetrel/Golang/tree/master/greetings"
)

func main() {
	message := greetings.Hello("Telman")
	fmt.Println(message)
}
