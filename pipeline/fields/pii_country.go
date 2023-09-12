package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var CountryLabel = "PII_COUNTRY"

type CountryEntry = utils.ProdigyOutput

var sampleTextCountry = []string{
	" country %v ",
	" country:  %v ",
	" %v ",
	" nationality: %v ",
}

var training_statements_country = []CountryEntry{
	{
		Text: sampleTextCountry[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextCountry[0], 9),
				Label: CountryLabel,
			},
		},
	},
	{
		Text: sampleTextCountry[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 11,
				End:   utils.GetWholeEntity(sampleTextCountry[1], 11),
				Label: CountryLabel,
			},
		},
	},
	{
		Text: sampleTextCountry[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextCountry[2], 1),
				Label: CountryLabel,
			},
		},
	},
	{
		Text: sampleTextCountry[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 14,
				End:   utils.GetWholeEntity(sampleTextCountry[3], 14),
				Label: CountryLabel,
			},
		},
	},
}

func GetTextForFieldCountry(country string) CountryEntry {
	return utils.GetRandomEntry(training_statements_country, country)
}
