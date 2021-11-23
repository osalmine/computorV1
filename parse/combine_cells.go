/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   combine_cells.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:42 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:44:22 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"fmt"

	"github.com/osalmine/computorV1/utils"
)

func isExponentCombined(combinedExponents []int, exponent int) bool {
	for _, combinedExponent := range combinedExponents {
		if combinedExponent == exponent {
			return true
		}
	}
	return false
}

func combineTwoCells(cell1 Cell, cell2 Cell) Cell {
	var combinedCell Cell
	combinedCell.Exponent = cell1.Exponent
	combinedCell.Coefficient = cell1.Coefficient + cell2.Coefficient
	combinedCell.Variable = cell1.Variable || cell2.Variable
	combinedCell.Visual.CapitalX = cell1.Visual.CapitalX || cell2.Visual.CapitalX
	combinedCell.Visual.DisplayExponent = cell1.Visual.DisplayExponent || cell2.Visual.DisplayExponent
	return combinedCell
}

func CombineCells(cells []Cell, options Options) []Cell {
	var combinedCells []Cell
	var combinedExponents []int
	for currentCellIndex, cell := range cells {
		combinationCell := cell
		currentCellExponent := cell.Exponent
		for _, nextCell := range cells[currentCellIndex+1:] {
			if currentCellExponent == nextCell.Exponent &&
				!isExponentCombined(combinedExponents, currentCellExponent) {
				combinationCell = combineTwoCells(combinationCell, nextCell)
				currentCellIndex++
			}
		}
		if combinationCell.Coefficient != 0 &&
			!isExponentCombined(combinedExponents, currentCellExponent) {
			combinedCells = append(combinedCells, combinationCell)
		}
		combinedExponents = append(combinedExponents, currentCellExponent)
	}
	if options.ShowCells {
		fmt.Println("Combined cells:")
		utils.PrettyPrintCells(combinedCells)
		fmt.Print("\n")
	}
	return combinedCells
}
