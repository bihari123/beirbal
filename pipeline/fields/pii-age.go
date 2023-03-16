package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var AgeLabel = "pii-age"

type AgeEntry = utils.ProdigyOutput

var training_statements_age = []utils.ProdigyOutput{
	AgeEntry{
		Text:  "His age is %v years ",
		Spans: []prose.LabeledEntity{{Start: 11, End: 19, Label: AgeLabel}},
	},
	AgeEntry{
		Text:  "His age is %v yrs ",
		Spans: []prose.LabeledEntity{{Start: 11, End: 17, Label: AgeLabel}},
	},
	AgeEntry{
		Text:  "Her age is %v yrs ",
		Spans: []prose.LabeledEntity{{Start: 11, End: 17, Label: AgeLabel}},
	},
	AgeEntry{
		Text:  "Her age is %v years",
		Spans: []prose.LabeledEntity{{Start: 11, End: 19, Label: AgeLabel}},
	},

	AgeEntry{
		Text:  "He is %v yrs old ",
		Spans: []prose.LabeledEntity{{Start: 6, End: 16, Label: AgeLabel}},
	},

	AgeEntry{
		Text:  "She is %v yrs old ",
		Spans: []prose.LabeledEntity{{Start: 7, End: 16, Label: AgeLabel}},
	},
	AgeEntry{
		Text:  "She is %v years old ",
		Spans: []prose.LabeledEntity{{Start: 7, End: 18, Label: AgeLabel}},
	},
	AgeEntry{
		Text:  "he is %v yrs old ",
		Spans: []prose.LabeledEntity{{Start: 6, End: 16, Label: AgeLabel}},
	},
}

func GetTextForFieldAge(age string) AgeEntry {
	return utils.GetRandomEntry(training_statements_age, age)
}
