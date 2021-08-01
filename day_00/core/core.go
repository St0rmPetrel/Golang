package core

import (
	"os"
)

type Result struct {
	Mean   float64
	Median float64
	Mode   int
	Sd     float64
}

func Calculate(ch chan int) Result {
	var ret Result
	var d data

	d.store_data(ch)
	if d.size == 0 {
		println("Empty Input")
		os.Exit(1)
	}
	ret.Mean = d.calculate_mean()
	ret.Median = d.calculate_median()
	ret.Mode = d.calculate_mode()
	ret.Sd = d.calculate_sd()
	return ret
}
