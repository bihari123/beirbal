package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var NameLabel = "PII_NAME"

type NameEntry = utils.ProdigyOutput

var sampleTextName = []string{
	" name %v ",
	" first name  %v ",
	" %v ",
	" last name %v ",
	" middle name %v ",
}

var training_statements_name = []NameEntry{
	{
		Text: sampleTextName[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 6,
				End:   utils.GetWholeEntity(sampleTextName[0], 6),
				Label: NameLabel,
			},
		},
	},
	{
		Text: sampleTextName[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 13,
				End:   utils.GetWholeEntity(sampleTextName[1], 13),
				Label: NameLabel,
			},
		},
	},
	{
		Text: sampleTextName[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextName[2], 1),
				Label: NameLabel,
			},
		},
	},
	{
		Text: sampleTextName[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 11,
				End:   utils.GetWholeEntity(sampleTextName[3], 11),
				Label: NameLabel,
			},
		},
	},
	{
		Text: sampleTextName[4],
		Spans: []prose.LabeledEntity{
			{
				Start: 13,
				End:   utils.GetWholeEntity(sampleTextName[3], 13),
				Label: NameLabel,
			},
		},
	},
}

func GetTextForFieldName(name string) NameEntry {
	return utils.GetRandomEntry(training_statements_name, name)
}
