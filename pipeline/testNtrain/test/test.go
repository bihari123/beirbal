package test

import (
	"C"
	"fmt"
	"io/ioutil"

	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

import (
	"strings"
)

//export Test
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
		fmt.Println("the input text is ", entry.Text)
		ents := doc.Entities()
		tokens := doc.Tokens()

		for _, t := range tokens {

			token_label := strings.Split(t.Label, "-")
			if len(token_label) > 1 {

				label_already_found := false

				for _, ent := range ents {
					if token_label[1] == ent.Label {
						label_already_found = true
					}
				}

				if !label_already_found {
					ents = append(ents, prose.Entity{
						Text:  t.Text,
						Label: token_label[1],
					})
				}
			}
		}

		fmt.Println("the entities are ", ents)
		if len(ents) == 0 {
			if len(entry.Spans) == 0 {
				correct++
			} else {
				continue
			}
		} else if entry.Answer != "accept" {
			// If we rejected this entity during annotation, prose shouldn't
			// have labeled it.
			correct++
		} else {
			// Otherwise, we need to verify that we found the correct entities.
			expected := []string{}
			for _, span := range entry.Spans {
				expected = append(expected, entry.Text[span.Start:span.End])
			}
			// fmt.Printf("Expected: %v Ents: %v\n", expected, ents)
			/*
				if reflect.DeepEqual(expected, ents) {
					correct++
				}
			*/
			for _, expectedSpan := range entry.Spans {
				for _, foundSpan := range entry.Spans {
					if expectedSpan.Label == foundSpan.Label {
						correct++
					}
				}
			}
		}
	}
	fmt.Printf("Correct (%%): %f\n", correct/float64(len(test)))
}

// export GetLabels
func GetLabels(inputText string) (labels []string) {
	modelPath := "./output/model/ve_nlp_model"
	model := prose.ModelFromDisk(modelPath)

	doc, err := prose.NewDocument(
		inputText,
		prose.WithSegmentation(false),
		prose.UsingModel(model),
	)
	if err != nil {
		panic(err)
	}

	ents := doc.Entities()
	tokens := doc.Tokens()

	for _, t := range tokens {

		token_label := strings.Split(t.Label, "-")
		if len(token_label) > 1 {

			label_already_found := false

			for _, ent := range ents {
				if token_label[1] == ent.Label {
					label_already_found = true
				}
			}

			if !label_already_found {
				ents = append(ents, prose.Entity{
					Text:  t.Text,
					Label: token_label[1],
				})
			}
		}
	}

	for _, ent := range ents {
		labels = append(labels, ent.Label)
	}

	return
}

// Iterate over the doc's tokens:
/*	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}
*/
/*
	// Iterate over the doc's named-entities:
	for _, ent := range ents {
		fmt.Println(ent.Text, ent.Label)
		// 	// Go GPE
		// 	// Google GPE
	}
*/
