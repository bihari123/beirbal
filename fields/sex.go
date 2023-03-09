package fields

import (
	"example.com/utils"
	"github.com/jdkato/prose/v2"
)

var SexLabel = "pii-sex"

type SexEntry = utils.ProdigyOutput

var sampleTextSex = []string{
	"She is a %v",
	"He is a %v",
	"His gender is %v",
	"her gender is %v",
	"he identifies as a %v",
	"she identifies as a %v",
}

var training_statements_sex = []SexEntry{
	{
		Text: sampleTextSex[0],
		Spans: []prose.LabeledEntity{
			{Start: 9, End: utils.GetWholeEntity(sampleTextSex[0], 9), Label: SexLabel},
		},
	},

	{
		Text: sampleTextSex[1],
		Spans: []prose.LabeledEntity{
			{Start: 8, End: utils.GetWholeEntity(sampleTextSex[1], 8), Label: SexLabel},
		},
	},

	{
		Text: sampleTextSex[2],
		Spans: []prose.LabeledEntity{
			{Start: 14, End: utils.GetWholeEntity(sampleTextSex[2], 14), Label: SexLabel},
		},
	},
	{
		Text: sampleTextSex[3],
		Spans: []prose.LabeledEntity{
			{Start: 14, End: utils.GetWholeEntity(sampleTextSex[3], 14), Label: SexLabel},
		},
	},
	{
		Text: sampleTextSex[4],
		Spans: []prose.LabeledEntity{
			{Start: 19, End: utils.GetWholeEntity(sampleTextSex[4], 19), Label: SexLabel},
		},
	},
	{
		Text: sampleTextSex[5],
		Spans: []prose.LabeledEntity{
			{Start: 20, End: utils.GetWholeEntity(sampleTextSex[5], 20), Label: SexLabel},
		},
	},
}

func GetTextForFieldSex(sex string) string {
	return utils.GetRandomEntry(training_statements_sex, sex)
}
