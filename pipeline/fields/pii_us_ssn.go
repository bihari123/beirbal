package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var UsSSNLabel = "PII_US_SSN"

type UsSsnEntry = utils.ProdigyOutput

var sampleTextUsSsn = []string{
	"ssn %v ",
	" %v ",
	" social security %v ",
	"ssns %v ",
	" social security number %v ",
}

var training_statements_us_ssn = utils.GenTrainingStatement(sampleTextUsSsn, UsSSNLabel)

func GetTextForFieldUSSsn(ssn string) UsSsnEntry {
	return utils.GetRandomEntry(training_statements_us_ssn, ssn)
}
