package main

import (
	"bufio"
	"io"
	"strings"
)

func line_counter(r *bufio.Reader) (uint, error) {
	var count uint = 0

	for {
		_, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err == nil {
			count++
		} else {
			return 0, err
		}
	}
	return count, nil
}

func char_counter(r *bufio.Reader) (uint, error) {
	var count uint = 0

	for {
		_, _, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if err == nil {
			count++
		} else {
			return 0, err
		}
	}
	return count, nil
}

func word_counter(r *bufio.Reader) (uint, error) {
	var count uint = 0
	var f bool

	word := strings.Builder{}
	for {
		r, _, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if err == nil {
			if r == '\n' {
				count = newword(count, &word, &f)
			} else if r == ' ' {
				count = newword(count, &word, &f)
			} else {
				f = true
				word.WriteRune(r)
			}
		} else {
			return 0, err
		}
	}
	return count, nil
}

func newword(count uint, w *strings.Builder, f *bool) uint {
	if *f {
		w.Reset()
		return count + 1
	}
	*f = false
	return count
}
