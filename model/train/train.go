package model

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

func train(filePath, modelName string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	train, test := utils.Split(utils.ReadProdigy(data))

	// Here, we're training a new model named PRODUCT with the training portion
	// of our annotated data.
	//
	// Depending on your hardware, this should take around 1 - 3 minutes.
	model := prose.ModelFromData(modelName, prose.UsingEntities(train))
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
	model.Write(modelName) // Save the model to disk.
	// model = onnx.Model
	// fmt.Print(model)
}
