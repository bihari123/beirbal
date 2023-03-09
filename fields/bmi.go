package fields

import (
	"fmt"

	"example.com/utils"
)

var training_statements_bmi = []string{
	"her bmi is %v",
	"Her body mass index is %v",
	"His body mass index goes upto %v",
}

func GetTextForFieldBMI(bmi float64) string {
	return fmt.Sprintf(utils.GetRandom(training_statements_bmi), bmi)
}
