package utils

import "strings"

func RemoveWhitespace(inputString string) string {
	output := strings.Join(strings.Fields(inputString), "")
	return output
}
