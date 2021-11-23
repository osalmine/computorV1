/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   transform_raw_cells.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:15:26 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:47:42 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

func checkForVariable(rawCell RawCell) bool {
	return rawCell.VariableCount != 0
}

func transformRawCellCoefficient(coefficients []float64) float64 {
	var result float64 = 1
	for _, coefficient := range coefficients {
		result *= coefficient
	}
	return result
}

func transformRawCellExponents(rawCell RawCell) int {
	var result int
	for _, exponent := range rawCell.Exponents {
		result += exponent
	}
	return result
}

func transformRawCellsToCells(rawCells []RawCell) []Cell {
	var Cells []Cell
	for _, rawCell := range rawCells {
		Cells = append(Cells, Cell{
			Coefficient: transformRawCellCoefficient(rawCell.Coefficients),
			Exponent:    transformRawCellExponents(rawCell),
			Variable:    checkForVariable(rawCell),
			Visual:      rawCell.Visual,
		})
	}
	return Cells
}
