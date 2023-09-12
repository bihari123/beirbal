package train

import (
	"fmt"
	"os"

	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

func Train(filePath, modelName string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("splitting the data into test and train data\n\n")
	train, _ := utils.Split(utils.ReadProdigy(data))
	fmt.Printf("data splitting completed\n\n")
	fmt.Printf("training started")
	i := 0
	for _, t := range train {
		if t.Accept {
			i++
		}
	}
	fmt.Printf("total number of test data: %v \n\n\n", len(train))
	model := prose.ModelFromData(modelName, prose.UsingEntities(train))
	os.Chdir("./output/model/")
	fmt.Printf("Model complete.vWriting it on disk\n\n")
	if utils.Exists(fmt.Sprintf("./%v", modelName)) {
		os.RemoveAll(fmt.Sprintf("./%v", modelName))
	}
	model.Write(modelName) // Save the model to disk.
	fmt.Printf("Model saved to disk\n\n\n")
	// model = onnx.Model
	// fmt.Print(model)
	os.Chdir("../../")
	fmt.Printf("training complete\n\n\n")

}
