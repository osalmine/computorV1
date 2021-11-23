/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_plus_minus.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:06 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:48:48 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
			output = append(output, input[j:])
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
