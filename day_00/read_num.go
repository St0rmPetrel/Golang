package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func read_num(ch chan int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if str != "" {
			str = str[:(len(str) - 1)]
			i, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			ch <- i
		}
		if err != nil {
			break
		}
	}
	close(ch)
}
