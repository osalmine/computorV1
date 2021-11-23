/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_multiplication.go                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:04 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:48:42 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"regexp"
	"strings"
)

func splitByMult(r rune) bool {
	return r == '*' || r == 'x' || r == 'X'
}

func SplitByMultiplication(input string) []string {
	var output []string
	if !regexp.MustCompile(`[0-9]+`).MatchString(input) || regexp.MustCompile(`[0-9]+$`).MatchString(input) {
		return []string{input}
	}
	for {
		if j := strings.LastIndexFunc(input, splitByMult); j >= 0 {
			if input[j:] != "*" {
				output = append(output, input[j:])
			}
			input = input[:j]
		} else {
			if len(input) > 0 {
				output = append(output, input)
			}
			break
		}
	}
	return output
}
