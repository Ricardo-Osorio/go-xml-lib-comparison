package data

import (
	_ "embed"
	"fmt"
	"os"
)

var (
	// alternative to loading the file on demand
	//go:embed input.xml
	RawXML []byte
)

func LoadEntireFile(path string) ([]byte, error) {
	if path == "" {
		// default to location from project root dir
		path = "pkg/data/input.xml"
	}

	// Read a local file into memory
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read local file (source of xml): %s", err.Error())
	}
	return contents, nil
}
