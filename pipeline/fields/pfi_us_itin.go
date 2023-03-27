package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var UsItinLabel = "PFI_US_ITIN"

type UsItinEntry = utils.ProdigyOutput

var sampleTextUsItin = []string{
	"ssn %v ",
	"taxpayer id %v ",
	"tax identification %v ",
	"itin %v ",
	"i.t.i.n. %v ",
	"ssns %v ",
	"tin %v ",
	"itins %v ",
	"taxid %v ",
}

var training_statements_us_itin = utils.GenTrainingStatement(sampleTextUsItin, UsItinLabel)

func GetTextForFieldUSItin(itin string) UsItinEntry {
	return utils.GetRandomEntry(training_statements_us_itin, itin)
}
