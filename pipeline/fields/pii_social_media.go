package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var SocialMediaLabel = "PII_SOCIAL_MEDIA"

type SocialMediaEntry = utils.ProdigyOutput

var sampleTextSocialMedia = []string{
	" facebook %v ",
	" linkedin %v ",
	" %v ",
	" instagram %v",
	" twitter %v ",
	" youtube %v ",
	" pininterest %v ",
	" fb %v ",
	" insta %v ",
}

var training_statements_social_media = []SocialMediaEntry{
	{
		Text: sampleTextSocialMedia[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 10,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[0], 10),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 10,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[1], 10),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[2], 1),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[3],
		Spans: []prose.LabeledEntity{
			{
				Start: 11,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[3], 11),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[4],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[4], 9),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[5],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[5], 9),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[6],
		Spans: []prose.LabeledEntity{
			{
				Start: 13,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[6], 13),
				Label: SocialMediaLabel,
			},
		},
	},
	{
		Text: sampleTextSocialMedia[7],
		Spans: []prose.LabeledEntity{
			{
				Start: 4,
				End:   utils.GetWholeEntity(sampleTextSocialMedia[7], 4),
				Label: SocialMediaLabel,
			},
		},
	},
}

func GetTextForFieldSocialMedia(social string) SocialMediaEntry {
	return utils.GetRandomEntry(training_statements_social_media, social)
}
