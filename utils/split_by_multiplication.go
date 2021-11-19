/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_multiplication.go                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:04 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/18 22:00:42 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"regexp"
	"strings"
)

// func SplitByMultiplication(stringToSplit string) []string {
// 	splittedString := strings.Split(stringToSplit, "*")
// 	return splittedString
// }

func splitByMult(r rune) bool {
	return r == '*' || r == 'x' || r == 'X'
}

func SplitByMultiplication(input string) []string {
	var output []string
	// fmt.Println("!regexp.MustCompile(`[0-9]+`).MatchString(input) || regexp.MustCompile(`[0-9]+$`).MatchString(input)", !regexp.MustCompile(`[0-9]+`).MatchString(input) || regexp.MustCompile(`[0-9]+$`).MatchString(input))
	if !regexp.MustCompile(`[0-9]+`).MatchString(input) || regexp.MustCompile(`[0-9]+$`).MatchString(input) {
		return []string{input}
	}
	for {
		if j := strings.LastIndexFunc(input, splitByMult); j >= 0 {
			// fmt.Println("STRINGS:", input[:j], "SECOND:", input[j:])
			if input[j:] != "*" {
				output = append(output, input[j:])
			}
			input = input[:j]
			// fmt.Println("output", output)
		} else {
			// fmt.Printf("Nothing left, input: %s, type: %T\n", input, input)
			if len(input) > 0 {
				output = append(output, input)
			}
			break
		}
	}
	// fmt.Println("Final output", output)

	// splittedArray := strings.FieldsFunc(input, splitBy)

	return output
}
