package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var UsHealthBenificiaryNumbeLabel = "PHI_HEALTH_BENIFICIARY_NUMBER"

type UsHealthBenificiaryNumberEntry = utils.ProdigyOutput

var sampleTextUsHealthBenificiaryNumber = []string{
	"HBN %v ",
	"health benificiary number %v ",
	" %v ",
}

var training_statements_us_health_benificiary = utils.GenTrainingStatement(
	sampleTextUsHealthBenificiaryNumber,
	UsHealthBenificiaryNumbeLabel,
)

func GetTextForFieldUSHealthBenificiaryNumber(
	benificiary_number string,
) UsHealthBenificiaryNumberEntry {
	return utils.GetRandomEntry(training_statements_us_health_benificiary, benificiary_number)
}
