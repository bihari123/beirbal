package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomEntry(input []ProdigyOutput, val string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(input[rand.Intn(len(input))])
}
