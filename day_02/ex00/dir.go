package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Find(d string, f Flags) {
	entry, _ := os.ReadDir(d)
	for _, file := range entry {
		name := addPrefix(file.Name(), d)
		if file.IsDir() {
			if f.d {
				fmt.Printf("%s\n", name)
			}
			Find(name, f)
		} else if fi, _ := file.Info(); fi.Mode()&os.ModeSymlink != 0 {
			if f.sl {
				printLink(name)
			}
		} else {
			if f.f {
				printFile(name, f.ext)
			}
		}
	}
}

func addPrefix(name, prefix string) string {
	ret := prefix
	if ret != "" && ret[len(ret)-1] != '/' {
		ret += "/"
	}
	ret += name
	return ret
}

func printFile(name, ext string) {
	if ext == "" {
		fmt.Printf("%s\n", name)
		return
	}
	if strings.HasSuffix(name, "."+ext) {
		fmt.Printf("%s\n", name)
	}
}

func printLink(name string) {
	link, _ := os.Readlink(name)
	prefix := filepath.Dir(name)
	link = addPrefix(link, prefix)
	if _, err := os.Stat(link); os.IsNotExist(err) {
		fmt.Printf("%s -> [broken]\n", name)
		return
	}
	fmt.Printf("%s -> %s\n", name, link)
}
