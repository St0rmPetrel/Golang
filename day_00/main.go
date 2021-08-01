package main

import (
	"fmt"

	"github.com/St0rmPetrel/Golang/day_00/core"
)

func main() {
	ch := make(chan int, 5)
	flags := init_flags()
	go read_num(ch)
	res := core.Calculate(ch)
	print_res(&res, &flags)
}

func print_res(res *core.Result, f *flags) {
	if f.mean {
		fmt.Printf("Mean:\t%.2f\n", res.Mean)
	}
	if f.median {
		fmt.Printf("Median:\t%.2f\n", res.Median)
	}
	if f.mode {
		fmt.Printf("Mode:\t%d\n", res.Mode)
	}
	if f.sd {
		fmt.Printf("SDec:\t%.2f\n", res.Sd)
	}
}
