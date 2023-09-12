package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var PhoneNumberLabel = "PII_PHONE_NUMBER"

type PhoneNumberEntry = utils.ProdigyOutput

var sampleTextPhoneNumber = []string{
	" contact %v ",
	" phone:  %v ",
	" %v ",
	" phone number: %v ",
	" mobile number: %v ",
	" contact number: %v ",
}

var training_statements_phone_number = []CountryEntry{
	{
		Text: sampleTextPhoneNumber[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextPhoneNumber[0], 9),
				Label: PhoneNumberLabel,
			},
		},
	},
	{
		Text: sampleTextPhoneNumber[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextPhoneNumber[1], 9),
				Label: PhoneNumberLabel,
			},
		},
	},
	{
		Text: sampleTextPhoneNumber[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextPhoneNumber[2], 1),
				Label: PhoneNumberLabel,
			},
		},
	},
	{
		Text: sampleTextPhoneNumber[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextPhoneNumber[3], 15),
				Label: PhoneNumberLabel,
			},
		},
	},
	{
		Text: sampleTextPhoneNumber[4],
		Spans: []prose.LabeledEntity{
			{
				Start: 16,
				End:   utils.GetWholeEntity(sampleTextPhoneNumber[3], 16),
				Label: PhoneNumberLabel,
			},
		},
	},
}

func GetTextForFieldPhoneNumber(phone_number string) PhoneNumberEntry {
	return utils.GetRandomEntry(training_statements_phone_number, phone_number)
}
