package utils

import "strings"

func SplitByEqual(stringToSplit string) []string {
	splittedString := strings.Split(stringToSplit, "=")
	return splittedString
}
