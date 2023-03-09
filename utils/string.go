package utils

// get the index where the current expression finishes
func GetWholeEntity(input string, startIndex int) int {
	for i := startIndex; i < len(input); i++ {
		if input[i] == ' ' {
			return i
		}
	}

	return len(input) - 1
}
