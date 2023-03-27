package main

import (
	"strings"

	"github.com/jdkato/prose/v2"
)
import "C"

// compile the code as:
// go build -ldflags="-s -w" -buildmode=c-shared -o libgo.dll main.go

// export main
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
}

//export GetLabels
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
