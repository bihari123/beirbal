package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var MoneyLabel = "PFI_FINANCE"

type MoneyEntry = utils.ProdigyOutput

var sampleTextMoney = []string{
	"He has $%v in his bank account",
	"He got a policy for $%v ",
	"$%v ",
	"He has $%v in his bank account",
	"He got a policy for $%v ",
}

// training the money data for 3-7 figures bank amount
var training_statements_money = []utils.ProdigyOutput{
	MoneyEntry{
		Text: sampleTextMoney[0],
		Spans: []prose.LabeledEntity{
			{Start: 7, End: utils.GetWholeEntity(sampleTextMoney[0], 7), Label: MoneyLabel},
		},
	},
	MoneyEntry{
		Text: sampleTextMoney[1],
		Spans: []prose.LabeledEntity{
			{Start: 20, End: utils.GetWholeEntity(sampleTextMoney[1], 20), Label: MoneyLabel},
		},
	},

	MoneyEntry{
		Text: sampleTextMoney[2],
		Spans: []prose.LabeledEntity{
			{Start: 0, End: utils.GetWholeEntity(sampleTextMoney[2], 0), Label: MoneyLabel},
		},
	},
	MoneyEntry{
		Text: sampleTextMoney[3],
		Spans: []prose.LabeledEntity{
			{Start: 7, End: utils.GetWholeEntity(sampleTextMoney[3], 7), Label: MoneyLabel},
		},
	},
	MoneyEntry{
		Text: sampleTextMoney[4],
		Spans: []prose.LabeledEntity{
			{Start: 20, End: utils.GetWholeEntity(sampleTextMoney[4], 20), Label: MoneyLabel},
		},
	},
}

func GetTextForFieldMoney(money string) MoneyEntry {
	return utils.GetRandomEntry(training_statements_money, money)
}
