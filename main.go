package main

import "example.com/model/test"

func main() {
	// train.LoadRawData()
	// train.Train("./dataset/testFile.jsonl", "demo_model")
	test.Test("./dataset/testFile.jsonl", "demo_model")
}
