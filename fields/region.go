package fields

import (
	"fmt"

	"example.com/utils"
)

var training_statements_region = []string{
	"He lives in %v region",
	"He belongs to %v region",
	"He went to %v for holidays",
}

func GetTextForFieldRegion(region string) string {
	return fmt.Sprintf(utils.GetRandom(training_statements_region), region)
}
