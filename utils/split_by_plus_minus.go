package utils

import (
	"strings"
)

func splitBy(r rune) bool {
	return r == '+' || r == '-'
}

func SplitByPlusMinus(input string) []string {
	var output []string
	for {
		if j := strings.LastIndexFunc(input, splitBy); j >= 0 {
			// fmt.Println("STRINGS:", input[:j], input[j:])
			output = append(output, input[j:])
			input = input[:j]
			// fmt.Println("output", output)
		} else {
			// fmt.Println("Nothing left, input:", input)
			output = append(output, input)
			break
		}
	}
	// fmt.Println("Final output", output)

	// splittedArray := strings.FieldsFunc(input, splitBy)

	return output
}
