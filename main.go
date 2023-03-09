package main

import (
	"fmt"

	"example.com/utils"
)

func main() {
	// test()
	// utils.ConvertGonToJson()
	data := utils.ReadCSV()

	fmt.Print(data)
}
