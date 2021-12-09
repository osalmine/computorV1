/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_coefficients.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:12:52 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/07 16:04:38 by osalmine         ###   ########.fr       */
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

func captureCoefficientWithoutPower(baseNumber string, coefficients *[]float64) {
	baseNumber = removeBrackets(baseNumber)
	if len(baseNumber) > 3 && (!unicode.IsDigit(rune(baseNumber[2])) && baseNumber[2] != '.') {
		panic("Invalid syntax")
	}
	if len(baseNumber) >= 3 {
		if (baseNumber[0] == '-' && baseNumber[1] == '+') || (baseNumber[0] == '+' && baseNumber[1] == '-') {
			baseNumber = "-" + string(baseNumber[2:])
		} else if (baseNumber[0] == '-' && baseNumber[1] == '-') || (baseNumber[0] == '+' && baseNumber[1] == '+') {
			baseNumber = string(baseNumber[2:])
		} else if !unicode.IsDigit(rune(baseNumber[2])) && baseNumber[2] != '.' {
			panic("Invalid syntax")
		}
	}
	value, err := strconv.ParseFloat(baseNumber, 64)
	if err != nil {
		panic(err)
	}
	*coefficients = append(*coefficients, value)
}

func captureCoefficientWithPower(originalString string, splittedByPower []string, coefficients *[]float64) {
	var baseValue string
	valueBetweenBrackets, found := getStringInBetween(originalString, "(", ")")
	if found {
		baseValue = valueBetweenBrackets
	} else {
		baseValue = splittedByPower[0]
	}
	value, err := strconv.ParseFloat(baseValue, 64)
	if err != nil {
		panic(err)
	}
	exponent, _ := strconv.Atoi(splittedByPower[1])
	value = utils.Pow(value, exponent)
	if originalString[0] == '-' {
		value = -value
	}
	*coefficients = append(*coefficients, value)
}

func captureCoefficient(str string, coefficients *[]float64) {
	splittedByPower := utils.SplitByPower(str)
	if len(splittedByPower) == 1 {
		captureCoefficientWithoutPower(splittedByPower[0], coefficients)
	} else {
		captureCoefficientWithPower(str, splittedByPower, coefficients)
	}
}

func addNegativeCoefficient(coefficients *[]float64, manualAddedCoefficient *bool) {
	*coefficients = append(*coefficients, -1)
	*manualAddedCoefficient = true
}

func getCoefficients(input []string) ([]float64, bool) {
	var coefficients []float64
	var manualAddedCoefficient bool
	for _, str := range input {
		if !strings.Contains(str, "x") && !strings.Contains(str, "X") {
			captureCoefficient(str, &coefficients)
		} else if str[0] == '-' {
			addNegativeCoefficient(&coefficients, &manualAddedCoefficient)
		}
	}
	return coefficients, manualAddedCoefficient
}
