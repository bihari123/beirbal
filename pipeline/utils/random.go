package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomEntry(input []ProdigyOutput, val string) (output ProdigyOutput) {
	rand.Seed(time.Now().UnixNano())
	output = input[rand.Intn(len(input))]
	output.Text = fmt.Sprintf(output.Text, val)
	output.Spans[0].End = GetWholeEntity(output.Text, output.Spans[0].Start)
	return
}
