/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_multiplication.go                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:04 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 17:11:17 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"regexp"
	"strings"
)

func splitByMult(r rune) bool {
	return r == '*' || r == 'x' || r == 'X'
	// return r == '*'
}

func SplitByMultiplication(input string) []string {
	var output []string
	// fmt.Println("input", input)
	// fmt.Println("!regexp.MustCompile(`[0-9]+`).MatchString(input)", !regexp.MustCompile(`[0-9]+`).MatchString(input))
	// fmt.Println("regexp.MustCompile(`[0-9]+$`).MatchString(input)", regexp.MustCompile(`[0-9]+$`).MatchString(input))
	// if !regexp.MustCompile(`[0-9]+`).MatchString(input) || regexp.MustCompile(`[0-9]+$`).MatchString(input) {
	// 	return []string{input}
	// }
	if !regexp.MustCompile(`[0-9]+`).MatchString(input) {
		return strings.Split(input, "*")
	}
	for {
		if j := strings.LastIndexFunc(input, splitByMult); j >= 0 {
			// fmt.Println("input[j:]", input[j:])
			// fmt.Println("input[j]", input[j])
			if j != 0 && (input[j-1] == '+' || input[j-1] == '-') {
				output = append(output, input[j-1:])
				j--
			} else if input[j:] != "*" {
				if input[j] == '*' {
					j++
				}
				output = append(output, input[j:])
			}
			// fmt.Println("INPUT BEF", input)
			input = input[:j]
			// fmt.Println("INPUT AFT", input)
		} else {
			if len(input) > 0 {
				output = append(output, input)
			}
			break
		}
	}
	// output := strings.Split(input, "*")
	return output
}
