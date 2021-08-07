package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Data struct {
	path []string
}

func (d *Data) LoadNew(f *os.File) error {
	d.path = nil
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d.path = append(d.path, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	sort.Strings(d.path)
	return nil
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) Compare(f *os.File, prefix string) error {
	scanner := bufio.NewScanner(f)
	end := len(d.path)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		if id := sort.SearchStrings(d.path, scanner.Text()); id == end ||
			d.path[id] != scanner.Text() {

			fmt.Printf("%s %s\n", prefix, scanner.Text())
		}
	}
	return nil
}

func diff(f_o, f_n *os.File) error {
	d := NewData()
	if err := d.LoadNew(f_o); err != nil {
		return err
	}
	if err := d.Compare(f_n, "ADDED"); err != nil {
		return err
	}

	if _, err := f_n.Seek(0, 0); err != nil {
		return err
	}
	if _, err := f_o.Seek(0, 0); err != nil {
		return err
	}
	if err := d.LoadNew(f_n); err != nil {
		return err
	}
	if err := d.Compare(f_o, "REMOVED"); err != nil {
		return err
	}
	return nil
}
