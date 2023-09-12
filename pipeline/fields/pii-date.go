package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var DateLabel = "PII_DATE"

type DateEntry = utils.ProdigyOutput

var sampleTextDate = []string{
	"date %v ",
	"%v ",
}

var training_statements_date = utils.GenTrainingStatement(sampleTextDate, DateLabel)

func GetTextForFieldDate(date string) DateEntry {
	return utils.GetRandomEntry(training_statements_date, date)
}
