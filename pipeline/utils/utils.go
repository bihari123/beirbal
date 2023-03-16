package utils

import (
	"encoding/gob"
	"fmt"
	"os"
)

// checkError panics if `err` is not `nil`.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConvertGonToJson() {
	var data []int

	// open data file
	dataFile, err := os.Open("./PRODUCT/Maxent/mapping.gob")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile.Close()

	fmt.Println(data)
}
