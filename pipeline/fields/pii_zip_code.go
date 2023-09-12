package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var ZipCodeLabel = "PII_ZIP_CODE"

type ZipCodeEntry = utils.ProdigyOutput

var sampleTextZipCode = []string{
	" zip code %v ",
	" zip:  %v ",
	" %v ",
}

var training_statements_zip_code = []ZipCodeEntry{
	{
		Text: sampleTextZipCode[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 10,
				End:   utils.GetWholeEntity(sampleTextZipCode[0], 10),
				Label: ZipCodeLabel,
			},
		},
	},
	{
		Text: sampleTextZipCode[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 7,
				End:   utils.GetWholeEntity(sampleTextZipCode[1], 7),
				Label: ZipCodeLabel,
			},
		},
	},
	{
		Text: sampleTextZipCode[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextZipCode[2], 1),
				Label: ZipCodeLabel,
			},
		},
	},
}

func GetTextForFieldZipCode(zipcode string) ZipCodeEntry {
	return utils.GetRandomEntry(training_statements_zip_code, zipcode)
}
