package main

import (
	"bufio"
	"os"
	"log"
	"syscall"
	"io"
)

func main() {
	args := make([]string, 1)
	args = append(os.Args[1:])
	r := bufio.NewReader(os.Stdin)
	for {
		str, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			args = append(args, str[:len(str)-1])
		}
	}
	if len(args) < 1 {
		println("Error: no arguments")
		return
	}
	cmd := search_in_path(args[0])
	if err := syscall.Exec(cmd, args, os.Environ()); err != nil {
		log.Fatal(err)
	}
}

func search_in_path(cmd string) string {
	return cmd
}
