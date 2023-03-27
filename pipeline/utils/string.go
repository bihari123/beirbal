package utils

import (
	"github.com/jdkato/prose/v2"
)

// get the index where the current expression finishes
func GetWholeEntity(input string, startIndex int) int {
	for i := startIndex; i < len(input)-1; i++ {
		if (input[i] == ' ') &&
			!(IsNumeric(input[i-1]) && IsNumeric(input[i+1])) {
			return i
		}
	}

	return len(input) - 1
}

func GenTrainingStatement(
	inputArr []string,
	label string,
) (trainingStatements []ProdigyOutput) {
	for _, elem := range inputArr {
		trainingStatements = append(trainingStatements, ProdigyOutput{
			Text: elem,
			Spans: []prose.LabeledEntity{
				{
					Start: GetStart(elem),
					Label: label,
				},
			},
			Answer: GetAnswer(),
		},
		)
	}

	return
}

func GetStart(input string) int {
	for index := range input {
		if index < len(input)-1 {
			if input[index] == '%' && input[index+1] == 'v' {
				return index
			}
		}
	}
	return 0
}
