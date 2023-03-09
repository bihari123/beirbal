package utils

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/jdkato/prose/v2"
)

func test() {
	data, err := ioutil.ReadFile("./annoted-reddit-product-ner.jsonl")
	if err != nil {
		panic(err)
	}
	_, test := Split(ReadProdigy(data))
	model := prose.ModelFromDisk("./PRODUCT")
	correct := 0.0
	for _, entry := range test {
		// Create a document without segmentation, which isn't required for NER.
		doc, err := prose.NewDocument(
			entry.Text,
			prose.WithSegmentation(false),
			prose.UsingModel(model))
		if err != nil {
			panic(err)
		}
		ents := doc.Entities()

		if entry.Answer != "accept" && len(ents) == 0 {
			// If we rejected this entity during annotation, prose shouldn't
			// have labeled it.
			correct++
		} else {
			// Otherwise, we need to verify that we found the correct entities.
			expected := []string{}
			for _, span := range entry.Spans {
				expected = append(expected, entry.Text[span.Start:span.End])
			}
			fmt.Printf("Expected: %v Ents: %v\n", expected, ents)
			if reflect.DeepEqual(expected, ents) {
				correct++
			}
		}
	}
	fmt.Printf("Correct (%%): %f\n", correct/float64(len(test)))
}

func train() {
	data, err := ioutil.ReadFile("./annoted-reddit-product-ner.jsonl")
	if err != nil {
		panic(err)
	}
	train, test := Split(ReadProdigy(data))

	// Here, we're training a new model named PRODUCT with the training portion
	// of our annotated data.
	//
	// Depending on your hardware, this should take around 1 - 3 minutes.
	model := prose.ModelFromData("PRODUCT", prose.UsingEntities(train))
	// Now, let's test our model:
	correct := 0.0
	for _, entry := range test {
		// Create a document without segmentation, which isn't required for NER.
		doc, err := prose.NewDocument(
			entry.Text,
			prose.WithSegmentation(false),
			prose.UsingModel(model))
		if err != nil {
			panic(err)
		}
		ents := doc.Entities()

		if entry.Answer != "accept" && len(ents) == 0 {
			// If we rejected this entity during annotation, prose shouldn't
			// have labeled it.
			correct++
		} else {
			// Otherwise, we need to verify that we found the correct entities.
			expected := []string{}
			for _, span := range entry.Spans {
				expected = append(expected, entry.Text[span.Start:span.End])
			}
			fmt.Printf("Expected: %v Ents: %v\n", expected, ents)
			if reflect.DeepEqual(expected, ents) {
				correct++
			}
		}
	}
	fmt.Printf("Correct (%%): %f\n", correct/float64(len(test)))
	model.Write("PRODUCT") // Save the model to disk.
	// model = onnx.Model
	// fmt.Print(model)
}

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
