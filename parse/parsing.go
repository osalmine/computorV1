/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parsing.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:47 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/24 20:48:45 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/utils"
)

type RawCell = computorV1.RawCell
type VisualRepresentation = computorV1.VisualRepresentation
type Cell = computorV1.Cell

func combineRawCellSlices(leftSideRawCells, rightSideRawCells []RawCell) []RawCell {
	return append(leftSideRawCells, rightSideRawCells...)
}

func Parse(input string, options Options) []Cell {
	noWhitespacesInput := utils.RemoveWhitespace(input)
	utils.PrintOnOption(options.Verbose, "Whitespaces removed:", noWhitespacesInput)
	inputSplitByEquation := utils.SplitByEqual(noWhitespacesInput)
	utils.PrintOnOption(options.Verbose, "Input split by equation:", inputSplitByEquation)
	splitLeft := utils.SplitByPlusMinus(inputSplitByEquation[0])
	splitRight := utils.SplitByPlusMinus(inputSplitByEquation[1])
	utils.PrintOnOption(options.Verbose, "Both sides split by + and -:", splitLeft, splitRight)
	leftSideRawCells := parseToRawCells(splitLeft)
	rightSideRawCells := parseToRawCells(splitRight)
	rightSideRawCells = negateRawCells(rightSideRawCells)
	combinedRawCells := combineRawCellSlices(leftSideRawCells, rightSideRawCells)
	if options.ShowCells {
		utils.PrettyPrintRawCells(combinedRawCells)
	}
	cells := transformRawCellsToCells(combinedRawCells)
	if options.ShowCells {
		utils.PrettyPrintCells(cells)
	}
	utils.PrintOnOption(options.Verbose, "Parsing done\n")
	return cells
}
