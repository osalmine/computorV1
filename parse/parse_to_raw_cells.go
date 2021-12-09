/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse_to_raw_cells.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:14:36 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/07 16:05:01 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"strings"

	"github.com/osalmine/computorV1/utils"
)

func parseToRawCells(input []string) []RawCell {
	var RawCells []RawCell
	for _, str := range input {
		splitByMultiplication := utils.SplitByMultiplication(str)
		rawCellCoefficients, manuallyAddedCoefficient := getCoefficients(splitByMultiplication)
		RawCells = append(RawCells, RawCell{
			VariableCount: countVariables(str),
			Coefficients:  rawCellCoefficients,
			Exponents:     getExponents(splitByMultiplication),
			Visual: VisualRepresentation{
				CapitalX:           getVariableCapitalization(str),
				DisplayExponent:    strings.Contains(str, "x^") || strings.Contains(str, "X^"),
				DisplayCoefficient: rawCellCoefficients != nil && !manuallyAddedCoefficient,
			},
		})
	}
	return RawCells
}
