package utils

import "regexp"

// create a function to check if string is numeric
func IsNumeric(word byte) bool {
	return regexp.MustCompile(`\d`).MatchString(string(word))
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}
