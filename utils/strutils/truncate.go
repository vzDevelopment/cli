package strutils

// Truncates the string so that it is equal or less to the provided limit.
func TruncateString(input string, limit int) string {
	if len(input) <= limit {
		return input
	}

	runes := []rune(input)
	return string(runes[0:limit])
}
