package main

import "example.com/testNtrain/train"

func main() {
	train.LoadRawData()
	train.Train("./output/json/testFile.jsonl", "demo_model")
	// test.Test("./dataset/testFile.jsonl", "demo_model")
}
