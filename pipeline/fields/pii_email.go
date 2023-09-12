package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var EmailLabel = "PII_EMAIL"

type EmailEntry = utils.ProdigyOutput

var sampleTextEmail = []string{
	" email %v ",
	" contact %v ",
	" %v ",
	" email address %v ",
}

var training_statements_email = []EmailEntry{
	{
		Text: sampleTextEmail[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 7,
				End:   utils.GetWholeEntity(sampleTextEmail[0], 7),
				Label: EmailLabel,
			},
		},
	},
	{
		Text: sampleTextEmail[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextEmail[1], 9),
				Label: EmailLabel,
			},
		},
	},
	{
		Text: sampleTextEmail[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextEmail[2], 1),
				Label: EmailLabel,
			},
		},
	},
	{
		Text: sampleTextEmail[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 15,
				End:   utils.GetWholeEntity(sampleTextEmail[3], 15),
				Label: EmailLabel,
			},
		},
	},
}

func GetTextForFieldEmail(email string) EmailEntry {
	return utils.GetRandomEntry(training_statements_email, email)
}
