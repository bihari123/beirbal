package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
	"github.com/spf13/cast"
)
import "C"

// compile the code as:
// go build -ldflags="-s -w" -
//
//export Sum
func Sum(a, b int) int {
	return a + b
}

// stringtest :
//
//export stringtest
func stringtest(name *C.char) {
	s := strings.Fields(C.GoString(name))
	fmt.Println(s)
}

//export GetLabels
func GetLabels(inputText string) *C.char {
	// inputText := strings.Fields(C.GoString(inputBytes))
	var labels string
	fmt.Println(cast.ToString(inputText))

	modelPath := "../../../pipeline/output/model/ve_nlp_model"
	// modelPath := "./pipeline//output/model/ve_nlp_model"

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
		labels += "" + ent.Label
	}

	fmt.Println("labels", labels)

	return C.CString(labels)
}

//export TokeniseThis
func TokeniseThis() {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument("Go is an open-source programming language created at Google.")
	if err != nil {
		/* log.Fatal(err) */
		fmt.Println("error: ", err.Error())
		return
	}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	// Iterate over the doc's named-entities:
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Go GPE
		// Google GPE
	}

	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {
		fmt.Println(sent.Text)
		// Go is an open-source programming language created at Google.
	}
}

func main() {
	/*
		os.Chdir("./pipeline")
		fmt.Println("loading raw data into prodigy json format\n ")
		train.LoadRawData()
		os.Chdir("../")

		os.Chdir("./pipeline")
		fmt.Println("\n\ntraining golang prose lib")
		train.Train("./output/json/testFile.jsonl", "ve_nlp_model")
		os.Chdir("../")
	*/
	/*
		os.Chdir("./pipeline")
		fmt.Println("\n\ntesting using in golang prose lib")
		test.Test("./output/json/testFile.jsonl", "./output/model/ve_nlp_model")
		os.Chdir("../")
	*/

	/*
		fmt.Println("training using spaCy")

		cmd := exec.Command("python3", "./spaCy-integration/ve_nlp_pipeline.py")

		out, err := cmd.Output()
		if err != nil {
			println(err.Error())
			return
		}
		fmt.Println(string(out[:]))
	*/
	os.Chdir("./pipeline/dataset/pdf")
	utils.GeneratePdf()
	os.Chdir("../../../")
}
