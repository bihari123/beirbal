package main

import (
	"fmt"
	"os"

	"github.com/bihari123/beirbal/pipeline/testNtrain/test"
	"github.com/bihari123/beirbal/pipeline/testNtrain/train"
)

func main() {
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
