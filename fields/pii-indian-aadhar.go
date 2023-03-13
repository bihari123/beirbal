package fields

import (
	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

var AadharLabel = "pii-indian-aadhar"

type AadharEntry = utils.ProdigyOutput

var sampleTextAadhar = []string{
	"The indian unique identification for this account is %v ",
	"His indian unique identification for is %v ",
	"Aadhar Card number %v ",
	"Aadhar ID %v ",
	"UDIAI ID %v ",
	"UDAIA number %v ",
	"aadhar# %v ",
	"uid %v ",
	"aadhaar %v ",
	"आधार" + " %v",
}

// training the money data for 3-7 figures bank amount
var training_statements_aadhar = []utils.ProdigyOutput{
	AadharEntry{
		Text: sampleTextAadhar[0],
		Spans: []prose.LabeledEntity{
			{Start: 53, End: utils.GetWholeEntity(sampleTextAadhar[0], 53), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[1],
		Spans: []prose.LabeledEntity{
			{Start: 40, End: utils.GetWholeEntity(sampleTextAadhar[1], 40), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[2],
		Spans: []prose.LabeledEntity{
			{Start: 19, End: utils.GetWholeEntity(sampleTextAadhar[2], 19), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[3],
		Spans: []prose.LabeledEntity{
			{Start: 10, End: utils.GetWholeEntity(sampleTextAadhar[3], 10), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[4],
		Spans: []prose.LabeledEntity{
			{Start: 13, End: utils.GetWholeEntity(sampleTextAadhar[4], 13), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[5],
		Spans: []prose.LabeledEntity{
			{Start: 8, End: utils.GetWholeEntity(sampleTextAadhar[5], 8), Label: AadharLabel},
		},
	},
	AadharEntry{
		Text: sampleTextAadhar[6],
		Spans: []prose.LabeledEntity{
			{Start: 4, End: utils.GetWholeEntity(sampleTextAadhar[6], 4), Label: AadharLabel},
		},
	},

	AadharEntry{
		Text: sampleTextAadhar[7],
		Spans: []prose.LabeledEntity{
			{Start: 8, End: utils.GetWholeEntity(sampleTextAadhar[7], 8), Label: AadharLabel},
		},
	},

	AadharEntry{
		Text: sampleTextAadhar[8],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextAadhar[8], 5), Label: AadharLabel},
		},
	},
}

func GetTextForFieldAadhar(aadharId string) AadharEntry {
	return utils.GetRandomEntry(training_statements_aadhar, aadharId)
}
