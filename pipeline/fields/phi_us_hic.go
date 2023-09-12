package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var UsHicLabel = "PHI_US_HIC"

type UsHicEntry = utils.ProdigyOutput

var sampleTextUsHci = []string{
	" health insurance claim %v ",
	" hic %v ",
	" %v ",
	" us hic number %v ",
	" health insurance claim number %v ",
	" hic number %v ",
}

var training_statements_us_hic = []UsHicEntry{
	{
		Text: sampleTextUsHci[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 24,
				End:   utils.GetWholeEntity(sampleTextUsHci[0], 24),
				Label: UsHicLabel,
			},
		},
	},
	{
		Text: sampleTextUsHci[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 5,
				End:   utils.GetWholeEntity(sampleTextUsHci[1], 5),
				Label: UsHicLabel,
			},
		},
	},
	{
		Text: sampleTextUsHci[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextUsHci[2], 1),
				Label: UsHicLabel,
			},
		},
	},
	{
		Text: sampleTextUsHci[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextUsHci[3], 15),
				Label: UsHicLabel,
			},
		},
	},
	{
		Text: sampleTextUsHci[4],
		Spans: []prose.LabeledEntity{
			{
				Start: 31,
				End:   utils.GetWholeEntity(sampleTextUsHci[4], 31),
				Label: UsHicLabel,
			},
		},
	},
	{
		Text: sampleTextUsHci[5],
		Spans: []prose.LabeledEntity{
			{
				Start: 12,
				End:   utils.GetWholeEntity(sampleTextUsHci[5], 12),
				Label: UsHicLabel,
			},
		},
	},
}

func GetTextForFieldHCI(hci string) UsHicEntry {
	return utils.GetRandomEntry(training_statements_us_hic, hci)
}
