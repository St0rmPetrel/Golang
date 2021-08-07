package main

type Dir struct {
	Name string
}

func (d Dir) Find(f Flags) error {
	// recurcive search os.ReadDir
	return nil
}
