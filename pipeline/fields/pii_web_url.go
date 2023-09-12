package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var WebUrlLabel = "PII_WEB_URL"

type WebUrlEntry = utils.ProdigyOutput

var sampleTextWebUrl = []string{
	" link %v ",
	" url %v ",
	" %v ",
}

var training_statements_web_url = []WebUrlEntry{
	{
		Text: sampleTextWebUrl[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 6,
				End:   utils.GetWholeEntity(sampleTextWebUrl[0], 6),
				Label: WebUrlLabel,
			},
		},
	},
	{
		Text: sampleTextWebUrl[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 5,
				End:   utils.GetWholeEntity(sampleTextWebUrl[1], 5),
				Label: WebUrlLabel,
			},
		},
	},
	{
		Text: sampleTextWebUrl[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextWebUrl[2], 1),
				Label: WebUrlLabel,
			},
		},
	},
}

func GetTextForFieldWebUrl(web_url string) WebUrlEntry {
	return utils.GetRandomEntry(training_statements_web_url, web_url)
}
