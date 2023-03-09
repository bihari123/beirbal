package fields

import (
	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

var BmiLabel = "pii-bmi"

type BmiEntry = utils.ProdigyOutput

var training_statements_bmi = []BmiEntry{
	{
		Text:  "her bmi is %v",
		Spans: []prose.LabeledEntity{{Start: 11, End: 15, Label: BmiLabel}},
	},
	{
		Text:  "Her body mass index is %v",
		Spans: []prose.LabeledEntity{{Start: 23, End: 27, Label: BmiLabel}},
	},

	{
		Text: "His body mass index goes upto %v",

		Spans: []prose.LabeledEntity{{Start: 30, End: 34, Label: BmiLabel}},
	},
}

func GetTextForFieldBMI(bmi string) string {
	return utils.GetRandomEntry(training_statements_bmi, bmi)
}
