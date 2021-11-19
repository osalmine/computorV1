/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_plus_minus.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:06 by osalmine          #+#    #+#             */
/*   Updated: 2021/10/30 23:48:07 by osalmine         ###   ########.fr       */
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
			// fmt.Println("STRINGS:", input[:j], "SECOND:", input[j:])
			output = append(output, input[j:])
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
