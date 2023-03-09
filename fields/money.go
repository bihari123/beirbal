package fields

import (
	"fmt"

	"example.com/utils"
)

var training_statements_money = []string{
	"He has $%v in his bank account",
	"He got a policy for $%v",
	"$ %v",
}

func GetTextForFieldMoney(money float64) string {
	return fmt.Sprintf(utils.GetRandom(training_statements_money), money)
}
