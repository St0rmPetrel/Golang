package main

import (
	"fmt"
	"github.com/St0rmPetrel/Golang/day_00/core"
)

//type Result struct {
//	mean	float32
//	median	float32
//	mode	float32
//	sd		float32
//}

func main() {
	ch := make(chan int, 5)
	flags := init_flags()
	go read_num(ch)
	res := core.Calculate(ch)
	res.print(flags)
}

func (res *core.Result) print(f flags) {
	if (f.mean) {
		fmt.Printf("Mean:\t%.2f\n", res.mean)
	}
	if (f.median) {
		fmt.Printf("Median:\t%.2f\n", res.median)
	}
	if (f.mode) {
		fmt.Printf("Mode:\t%.2f\n", res.mode)
	}
	if (f.sd) {
		fmt.Printf("Std dec:\t%.2f\n", res.sd)
	}
}

//func calculate(ch chan int) result {
//	var ret result
//
//	for {
//		res, ok := <-ch
//		if ok == false {
//			println("Channel Close ")
//			break
//		}
//		println("Channel Open :", res)
//	}
//	return ret
//}
