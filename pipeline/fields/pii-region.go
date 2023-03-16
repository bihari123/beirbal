package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var RegionLabel = "geo-location"

type RegionEntry = utils.ProdigyOutput

var sampleTextRegion = []string{
	"He lives in %v region",
	"He belongs to %v region",
	"He went to %v for holidays",
}

var training_statements_region = []utils.ProdigyOutput{
	RegionEntry{
		Text: sampleTextRegion[0],
		Spans: []prose.LabeledEntity{
			{Start: 12, End: utils.GetWholeEntity(sampleTextRegion[0], 12), Label: RegionLabel},
		},
	},

	RegionEntry{
		Text: sampleTextRegion[1],
		Spans: []prose.LabeledEntity{
			{Start: 14, End: utils.GetWholeEntity(sampleTextRegion[1], 14), Label: RegionLabel},
		},
	},

	RegionEntry{
		Text: sampleTextRegion[2],
		Spans: []prose.LabeledEntity{
			{Start: 11, End: utils.GetWholeEntity(sampleTextRegion[2], 11), Label: RegionLabel},
		},
	},
}

func GetTextForFieldRegion(region string) RegionEntry {
	return utils.GetRandomEntry(training_statements_region, region)
}
