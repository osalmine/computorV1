/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   combine_cells.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:42 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:11:18 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

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

func CombineCells(cells []Cell) []Cell {
	// fmt.Printf("\n\nCOMBINING\n\n")
	var combinedCells []Cell
	var combinedExponents []int
	for currentCellIndex, cell := range cells {
		// fmt.Printf("\nCELL: %+v\n", cell)
		combinationCell := cell
		currentCellExponent := cell.Exponent
		// fmt.Printf("combinationCell:\t%+v\n", combinationCell)
		for _, nextCell := range cells[currentCellIndex+1:] {
			// fmt.Printf("NEXT CELL:\t\t%+v\n", nextCell)
			if currentCellExponent == nextCell.Exponent &&
				!isExponentCombined(combinedExponents, currentCellExponent) {
				// fmt.Println("COMBINATION FOUND")
				combinationCell = combineTwoCells(combinationCell, nextCell)
				currentCellIndex++
				// fmt.Printf("done combinationCell:\t%+v\n", combinationCell)
			}
		}
		if combinationCell.Coefficient != 0 &&
			!isExponentCombined(combinedExponents, currentCellExponent) {
			combinedCells = append(combinedCells, combinationCell)
		}
		combinedExponents = append(combinedExponents, currentCellExponent)
		// fmt.Println("combinedExponents", combinedExponents)
	}
	// fmt.Printf("combinedCells: %+v\n", combinedCells)
	return combinedCells
}
