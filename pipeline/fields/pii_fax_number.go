package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var FaxLabel = "PII_FAX_NUMBER"

type FaxEntry = utils.ProdigyOutput

var sampleTextFax = []string{
	" fax %v ",
	" fax number:  %v ",
	" %v ",
}

var training_statements_fax = []FaxEntry{
	{
		Text: sampleTextFax[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 5,
				End:   utils.GetWholeEntity(sampleTextFax[0], 5),
				Label: FaxLabel,
			},
		},
	},
	{
		Text: sampleTextFax[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 14,
				End:   utils.GetWholeEntity(sampleTextFax[1], 14),
				Label: FaxLabel,
			},
		},
	},
	{
		Text: sampleTextFax[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextFax[2], 1),
				Label: FaxLabel,
			},
		},
	},
}

func GetTextForFieldFax(fax string) FaxEntry {
	return utils.GetRandomEntry(training_statements_fax, fax)
}
