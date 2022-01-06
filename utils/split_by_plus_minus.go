/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_plus_minus.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:06 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/06 22:42:11 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"strings"
)

func splitBy(r rune) bool {
	return r == '+' || r == '-'
}

func indexIsBetweenBrackets(index int, bracketPositions []int) bool {
	for i, bracketPosition := range bracketPositions {
		if i%2 == 0 {
			if index >= bracketPosition && index <= bracketPositions[i+1] {
				return true
			}
		}
	}
	return false
}

func SplitByPlusMinus(input string, bracketPositions []int) []string {
	var output []string
	var tmp string
	for {
		if j := strings.LastIndexFunc(input, splitBy); j >= 0 {
			if !indexIsBetweenBrackets(j, bracketPositions) {
				output = append(output, input[j:]+tmp)
				input = input[:j]
				tmp = ""
			} else {
				tmp = input[j:] + tmp
				input = input[:j]
			}
		} else {
			if len(input) > 0 {
				output = append(output, input+tmp)
			}
			break
		}
	}
	return output
}
