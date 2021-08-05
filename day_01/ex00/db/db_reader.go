package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DBReader interface {
	Read() (*Recipes, error)
}

func GetDBReader(name string) (DBReader, error) {
	var f *os.File
	var err error
	var b []byte

	f, err = os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if is_json, _ := filepath.Match("*.json", filepath.Base(name)); is_json {
		return DataJSON{b}, nil
	} else if is_xml, _ := filepath.Match("*.xml", filepath.Base(name)); is_xml {
		return DataXML{b}, nil
	}
	return nil, &BadExtensionError{name}
}

type BadExtensionError struct {
	name string
}

func (e *BadExtensionError) Error() string {
	return fmt.Sprintf("bad filename extension in \"%s\"", e.name)
}
