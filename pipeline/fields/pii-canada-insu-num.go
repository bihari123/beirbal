package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

// CSIN = canada social insurance number
var CSINLabel = "pii-indian-aadhar"

type CSINEntry = utils.ProdigyOutput

var sampleTextCSIN = []string{
	"sin %v ",
	"social insurance %v ",
	"numero d'assurance sociale %v ",
	"sins %v ",
	"ssn %v ",
	"ssns %v ",
	"social security %v ",
	"numero d'assurance social %v ",
	"national identification number %v ",
	"national id %v ",
	"sin# %v ",
	"soc ins %v ",
	"social ins %v ",
	"NAS %v ",
	"numéro d’assurance social %v ",
	"numéro d’assurance sociale %v ",
	"assurance social %v ",
	"carte d’assurance sociale %v ",
}

// training the money data for 3-7 figures bank amount
var training_statements_CSIN = []utils.ProdigyOutput{
	CSINEntry{
		Text: sampleTextCSIN[0],
		Spans: []prose.LabeledEntity{
			{Start: 4, End: utils.GetWholeEntity(sampleTextCSIN[0], 4), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[1],
		Spans: []prose.LabeledEntity{
			{Start: 17, End: utils.GetWholeEntity(sampleTextCSIN[1], 17), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[2],
		Spans: []prose.LabeledEntity{
			{Start: 27, End: utils.GetWholeEntity(sampleTextCSIN[2], 27), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[3],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextCSIN[3], 5), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[4],
		Spans: []prose.LabeledEntity{
			{Start: 4, End: utils.GetWholeEntity(sampleTextCSIN[4], 4), Label: CSINLabel},
		},
	},

	CSINEntry{
		Text: sampleTextCSIN[5],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextCSIN[5], 5), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[6],
		Spans: []prose.LabeledEntity{
			{Start: 16, End: utils.GetWholeEntity(sampleTextCSIN[6], 16), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[7],
		Spans: []prose.LabeledEntity{
			{Start: 26, End: utils.GetWholeEntity(sampleTextCSIN[7], 26), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[8],
		Spans: []prose.LabeledEntity{
			{Start: 31, End: utils.GetWholeEntity(sampleTextCSIN[8], 31), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[9],
		Spans: []prose.LabeledEntity{
			{Start: 12, End: utils.GetWholeEntity(sampleTextCSIN[9], 12), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[10],
		Spans: []prose.LabeledEntity{
			{Start: 5, End: utils.GetWholeEntity(sampleTextCSIN[10], 5), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[11],
		Spans: []prose.LabeledEntity{
			{Start: 8, End: utils.GetWholeEntity(sampleTextCSIN[11], 8), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[12],
		Spans: []prose.LabeledEntity{
			{Start: 11, End: utils.GetWholeEntity(sampleTextCSIN[12], 11), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[13],
		Spans: []prose.LabeledEntity{
			{Start: 4, End: utils.GetWholeEntity(sampleTextCSIN[13], 4), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[14],
		Spans: []prose.LabeledEntity{
			{Start: 26, End: utils.GetWholeEntity(sampleTextCSIN[14], 26), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[15],
		Spans: []prose.LabeledEntity{
			{Start: 27, End: utils.GetWholeEntity(sampleTextCSIN[15], 27), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[16],
		Spans: []prose.LabeledEntity{
			{Start: 17, End: utils.GetWholeEntity(sampleTextCSIN[16], 17), Label: CSINLabel},
		},
	},
	CSINEntry{
		Text: sampleTextCSIN[17],
		Spans: []prose.LabeledEntity{
			{Start: 26, End: utils.GetWholeEntity(sampleTextCSIN[17], 26), Label: CSINLabel},
		},
	},
}

func GetTextForFieldCSIN(csin string) CSINEntry {
	return utils.GetRandomEntry(training_statements_CSIN, csin)
}
