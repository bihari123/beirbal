package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var AddressLabel = "PII_ADDRESS"

type AddressEntry = utils.ProdigyOutput

var sampleTextAddress = []string{
	" address %v ",
	" street %v ",
	" %v ",
	" residence %v ",
}

var training_statements_address = []AddressEntry{
	{
		Text: sampleTextAddress[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextAddress[0], 9),
				Label: AddressLabel,
			},
		},
	},
	{
		Text: sampleTextAddress[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 8,
				End:   utils.GetWholeEntity(sampleTextAddress[1], 8),
				Label: AddressLabel,
			},
		},
	},
	{
		Text: sampleTextAddress[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextAddress[2], 1),
				Label: AddressLabel,
			},
		},
	},
	{
		Text: sampleTextAddress[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 11,
				End:   utils.GetWholeEntity(sampleTextAddress[3], 11),
				Label: AddressLabel,
			},
		},
	},
}

func GetTextForFieldAddress(address string) AddressEntry {
	return utils.GetRandomEntry(training_statements_address, address)
}
