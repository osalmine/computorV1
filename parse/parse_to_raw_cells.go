/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse_to_raw_cells.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:14:36 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:14:42 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"strings"

	"github.com/osalmine/computorV1/utils"
)

func parseToRawCells(input []string) []RawCell {
	// fmt.Println("parseToRawCells INPUT", input)
	var RawCells []RawCell
	for _, str := range input {
		// fmt.Println("parseToRawCells STR", str)
		splitByMultiplication := utils.SplitByMultiplication(str)
		// fmt.Println("SPLIT BY MULTIPLICATION", splitByMultiplication)
		RawCells = append(RawCells, RawCell{
			VariableCount: countVariables(str),
			Coefficients:  getCoefficients(splitByMultiplication),
			Exponents:     getExponents(splitByMultiplication),
			Visual: VisualRepresentation{
				CapitalX:        getVariableCapitalization(str),
				DisplayExponent: strings.Contains(str, "x^") || strings.Contains(str, "X^"),
			},
		})
	}
	// fmt.Printf("RawCells %+v\n", RawCells)
	return RawCells
}