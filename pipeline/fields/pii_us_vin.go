package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var UsVinLabel = "PII_VIN"

type UsVinEntry = utils.ProdigyOutput

var sampleTextUsVin = []string{
	"vin %v ",
	"us vehicle identification number %v ",
	"us vin %v ",
}

var training_statements_us_vin = utils.GenTrainingStatement(sampleTextUsVin, UsVinLabel)

func GetTextForFieldUSVin(vin string) UsVinEntry {
	return utils.GetRandomEntry(training_statements_us_vin, vin)
}
