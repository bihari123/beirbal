package test

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

func Test(filepath, modelPath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	_, test := utils.Split(utils.ReadProdigy(data))
	model := prose.ModelFromDisk(modelPath)
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
