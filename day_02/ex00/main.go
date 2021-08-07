package main

import "log"

func main() {
	dir, f, err := init_flags()
	if err != nil {
		log.Fatal(err)
	}
	if err := dir.Find(f); err != nil {
		log.Fatal(err)
	}
}

type Flags struct {
	f   bool
	d   bool
	sl  bool
	ext string
}

func init_flags() (Dir, Flags, error) {
	return Dir{}, Flags{}, nil
}
