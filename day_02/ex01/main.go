package main

import (
	"flag"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mprint sync.Mutex

	fl, args, err := flag_init()
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range args {
		wg.Add(1)
		go wc(arg, fl, &wg, &mprint)
	}
	wg.Wait()
}

func flag_init() (Flags, []string, error) {
	var l, m, w bool

	flag.BoolVar(&l, "l", false, "counting lines")
	flag.BoolVar(&m, "m", false, "counting characters")
	flag.BoolVar(&w, "w", false, "counting words")
	flag.Parse()
	if flag.NArg() < 1 || flag.NFlag() > 1 {
		return Flags{}, nil, &ArgError{}
	}
	if !l && !m && !w {
		w = true
	}
	return Flags{l, m, w}, flag.Args(), nil
}

type Flags struct {
	l, m, w bool
}

type ArgError struct {
}

func (err *ArgError) Error() string {
	return "Wrong count of arguments"
}
