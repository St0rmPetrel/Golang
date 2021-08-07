package main

import (
	"fmt"
	"os"
)

func Find(d string, f Flags) error {
	entry, _ := os.ReadDir(d)
	for _, file := range entry {
		name := addPrefix(file.Name(), d)
		if file.IsDir() {
			if f.d {
				fmt.Printf("%s\n", name)
			}
			Find(name, f)
		} else {
			if f.f {
				fmt.Printf("%s\n", name)
			}
		}
	}
	return nil
}

func addPrefix(name, prefix string) string {
	ret := prefix
	if ret != "" && ret[len(ret)-1] != '/' {
		ret += "/"
	}
	ret += name
	return ret
}
