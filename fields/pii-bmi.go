package fields

import (
	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

var BmiLabel = "pii-bmi"

type BmiEntry = utils.ProdigyOutput

var training_statements_bmi = []BmiEntry{
	{
		Text: "her bmi is %v ",
		Spans: []prose.LabeledEntity{
			{Start: 11, End: utils.GetWholeEntity("her bmi is %v ", 11), Label: BmiLabel},
		},
	},
	{
		Text: "Her body mass index is %v ",
		Spans: []prose.LabeledEntity{
			{
				Start: 23,
				End:   utils.GetWholeEntity("Her body mass index is %v ", 23),
				Label: BmiLabel,
			},
		},
	},

	{
		Text: "His body mass index goes upto %v ",

		Spans: []prose.LabeledEntity{
			{
				Start: 30,
				End:   utils.GetWholeEntity("His body mass index goes upto %v ", 30),
				Label: BmiLabel,
			},
		},
	},
}

func GetTextForFieldBMI(bmi string) BmiEntry {
	return utils.GetRandomEntry(training_statements_bmi, bmi)
}
