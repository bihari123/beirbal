package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV() (data [][]string) {
	// open file
	f, err := os.Open("./dataset/insurance.csv")
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

	return
}
