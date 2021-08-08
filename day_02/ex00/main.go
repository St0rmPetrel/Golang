package main

import (
	"flag"
	"log"
)

func main() {
	dir, f, err := init_flags()
	if err != nil {
		log.Fatal(err)
	}
	Find(dir, f)
}

type Flags struct {
	f   bool
	d   bool
	sl  bool
	ext string
}

func init_flags() (string, Flags, error) {
	var name, ext string
	var f, d, sl bool

	flag.StringVar(&ext, "ext", "", "find files only with extention ext")
	flag.BoolVar(&f, "f", false, "find files")
	flag.BoolVar(&d, "d", false, "find direcoris")
	flag.BoolVar(&sl, "sl", false, "find symlink")
	flag.Parse()
	if !f && !d && !sl {
		f, d, sl = true, true, true
	}
	if !f && ext != "" || flag.NArg() != 1 {
		return "", Flags{}, &BadFlagError{}
	}
	name = flag.Arg(0)
	return name, Flags{f, d, sl, ext}, nil
}

type BadFlagError struct {
}

func (err *BadFlagError) Error() string {
	return "Bad args or flags"
}
