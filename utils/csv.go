package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func ReadCSV(filepath string) (data [][]string) {
	// open file
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// removing the whitespaces
	for indexRow, row := range data {
		for indexCol, entry := range row {
			data[indexRow][indexCol] = strings.TrimSpace(entry)
		}
	}

	return
}
