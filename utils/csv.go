package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCsv(path string) *csv.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	// remove header row
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return reader
}
