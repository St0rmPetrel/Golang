package main

import (
	"flag"
	"os"
)

type flags struct {
	mean	bool
	median	bool
	mode	bool
	sd		bool
}

func init_flags() flags {
	var f flags

	flag.BoolVar(&(f.mean), "mean", false, "display mean (average)")
	flag.BoolVar(&(f.median), "median", false, "display median")
	flag.BoolVar(&(f.mode), "mode", false, "display mode")
	flag.BoolVar(&(f.sd), "sd", false, "display standard deviation")
	flag.Parse()
	tail := flag.Args()
	if (len(tail) != 0) {
		println("Wrong number of arguments")
		os.Exit(1)
	}
	if !(f.mean || f.median || f.mode || f.sd) {
		f.mean, f.median, f.mode, f.sd = true, true, true, true
	}
	return f
}
