package fields

import (
	"fmt"

	"example.com/utils"
)

var training_statements_sex = []string{
	"She is a %v",
	"He is a %v",
	"His gender is %v",
	"her gender is %v",
	"he identifies as a %v",
	"she identifies as a %v",
}

func GetTextForFieldSex(sex string) string {
	return fmt.Sprintf(utils.GetRandom(training_statements_sex), sex)
}
