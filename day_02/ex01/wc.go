package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func wc(name string, fl Flags, wg *sync.WaitGroup, mprint *sync.Mutex) {
	var count uint

	defer wg.Done()
	file, err := os.Open(name)
	if err != nil {
		save_printf(mprint, "Error: can't open file \"%s\"\n", name)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	switch {
	case fl.l:
		count, err = line_counter(rd)
	case fl.m:
		count, err = char_counter(rd)
	case fl.w:
		count, err = word_counter(rd)
	}
	if err != nil {
		save_printf(mprint, "Error: can't read file \"%s\"\n", name)
		return
	}
	save_printf(mprint, "%v\t%s\n", count, name)
}

func save_printf(m *sync.Mutex, format string, a ...interface{}) {
	m.Lock()
	fmt.Printf(format, a...)
	m.Unlock()
}
