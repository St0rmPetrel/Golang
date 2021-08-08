package main

import (
	"flag"
	"log"
)

func main() {
	f, err := flag_init()
	if err != nil {
		log.Fatal(err)
	}
	println(f.l, f.m, f.w)
}

func flag_init() (Flags, error) {
	var l, m, w bool

	flag.BoolVar(&l, "l", false, "counting lines")
	flag.BoolVar(&m, "m", false, "counting characters")
	flag.BoolVar(&w, "w", false, "counting words")
	flag.Parse()
	if flag.NArg() < 1 || flag.NFlag() > 1 {
		return Flags{}, &ArgError{}
	}
	if !l && !m && !w {
		w = true
	}
	return Flags{l, m, w}, nil
}

type Flags struct {
	l, m, w bool
}

type ArgError struct {
}

func (err *ArgError) Error() string {
	return "Wrong count of arguments"
}
