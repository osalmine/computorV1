/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_coefficients.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:12:52 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/30 12:12:28 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/osalmine/computorV1/utils"
)

func getStringInBetween(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

func removeBrackets(str string) string {
	var output string
	output = strings.Replace(str, "(", "", -1)
	output = strings.Replace(output, ")", "", -1)
	return output
}

func getCoefficients(input []string) []float64 {
	var coefficients []float64
	for _, str := range input {
		if !strings.Contains(str, "x") && !strings.Contains(str, "X") {
			// fmt.Println("input before power split:", str)
			splittedByPower := utils.SplitByPower(str)
			// fmt.Println("splittedByPower", splittedByPower)
			baseNumber := splittedByPower[0]
			// fmt.Println("baseNumber", baseNumber)
			if len(splittedByPower) == 1 {
				baseNumber = removeBrackets(baseNumber)
				// fmt.Println("baseNumber after bracket remove:", baseNumber)
				if len(baseNumber) > 3 && (!unicode.IsDigit(rune(baseNumber[2])) && baseNumber[2] != '.') {
					panic("Invalid syntax")
				}
				if len(baseNumber) == 3 {
					if (baseNumber[0] == '-' && baseNumber[1] == '+') || (baseNumber[0] == '+' && baseNumber[1] == '-') {
						baseNumber = "-" + string(baseNumber[2])
					} else if (baseNumber[0] == '-' && baseNumber[1] == '-') || (baseNumber[0] == '+' && baseNumber[1] == '+') {
						baseNumber = string(baseNumber[2])
					} else {
						panic("Invalid syntax")
					}
				}
				// fmt.Println("Base number after combine:", baseNumber)
				value, err := strconv.ParseFloat(baseNumber, 64)
				if err != nil {
					panic(err)
				}
				if len(splittedByPower) > 1 {
					exponent, _ := strconv.Atoi(splittedByPower[1])
					value = utils.Pow(value, exponent)
				}
				// fmt.Println("Value added to coefficients:", value)
				coefficients = append(coefficients, value)
			} else {
				valueBetweenBrackets, found := getStringInBetween(str, "(", ")")
				// valueBetweenBrackets := strings.TrimLeft(strings.TrimRight(baseNumber, ")"), "(")
				// fmt.Println("valueBetweenBrackets", valueBetweenBrackets, "found:", found)
				if found {
					value, err := strconv.ParseFloat(valueBetweenBrackets, 64)
					if err != nil {
						panic(err)
					}
					// if len(splittedByPower) > 1 {
					exponent, _ := strconv.Atoi(splittedByPower[1])
					value = utils.Pow(value, exponent)
					// }
					if str[0] == '-' {
						value = -value
					}
					// fmt.Println("Value added to coefficients:", value)
					coefficients = append(coefficients, value)
				} else {
					value, err := strconv.ParseFloat(baseNumber, 64)
					if err != nil {
						panic(err)
					}
					// if len(splittedByPower) > 1 {
					exponent, _ := strconv.Atoi(splittedByPower[1])
					value = utils.Pow(value, exponent)
					// }
					if str[0] == '-' {
						value = -value
					}
					// fmt.Println("Value added to coefficients:", value)
					coefficients = append(coefficients, value)
				}
			}
		} else if str[0] == '-' {
			// fmt.Println("Value added to coefficients:", -1)
			coefficients = append(coefficients, -1)
		}
	}
	return coefficients
}
