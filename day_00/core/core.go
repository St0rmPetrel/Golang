package core

import (
	"os"
)

type Result struct {
	Mean   float32
	Median float32
	Mode   float32
	Sd     float32
}

func Calculate(ch chan int) Result {
	var ret Result
	var d data

	d.store_data(ch)
	if d.size == 0 {
		println("Empty Input")
		os.Exit(1)
	}
	ret.Mean = float32(d.sum / float64(d.size))
	ret.Median = d.calculate_median()
	return ret
}
