/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_coefficients.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:12:52 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:12:59 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"strconv"
	"strings"

	"github.com/osalmine/computorV1/utils"
)

func getCoefficients(input []string) []float64 {
	var coefficients []float64
	// fmt.Println("input", input)
	for _, str := range input {
		// fmt.Println("str", str)
		if !strings.Contains(str, "x") && !strings.Contains(str, "X") {
			splittedByPower := utils.SplitByPower(str)
			// fmt.Println("splittedByPower", splittedByPower)
			baseNumber := splittedByPower[0]
			value, _ := strconv.ParseFloat(baseNumber, 64)
			if len(splittedByPower) > 1 {
				exponent, _ := strconv.Atoi(splittedByPower[1])
				value = utils.Pow(value, exponent)
			}
			coefficients = append(coefficients, value)
			// fmt.Printf("COEFFICIENTS: %+v\n", coefficients)
		} else if str[0] == '-' {
			coefficients = append(coefficients, -1)
		}
	}
	return coefficients
}
