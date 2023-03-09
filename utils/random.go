package utils

import (
	"math/rand"
	"time"

	"github.com/jdkato/prose/v2"
)

func GetRandomEntry(input [] ProdigyOutput, val string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(input[rand.Intn(len(input))]
}
