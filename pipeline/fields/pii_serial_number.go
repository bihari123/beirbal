package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var SerialNumberLabel = "PII_SERIAL_NUMBER"

type SerialNumberEntry = utils.ProdigyOutput

var sampleTextSerialNumber = []string{
	"serial number %v ",
	"device id %v ",
	" %v ",
}

var training_statements_serial_number = utils.GenTrainingStatement(
	sampleTextSerialNumber,
	SerialNumberLabel,
)

func GetTextForFieldSerialNumber(serial_number string) SerialNumberEntry {
	return utils.GetRandomEntry(training_statements_serial_number, serial_number)
}
