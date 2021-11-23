/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parsing.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:47 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/22 17:28:35 by osalmine         ###   ########.fr       */
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
	// fmt.Println("leftRight:", leftRight)
	utils.PrintOnOption(options.Verbose, "Input split by equation:", inputSplitByEquation)
	splitLeft := utils.SplitByPlusMinus(inputSplitByEquation[0])
	// fmt.Println("left side:", splitLeft)
	splitRight := utils.SplitByPlusMinus(inputSplitByEquation[1])
	// fmt.Println("right side:", splitRight)
	utils.PrintOnOption(options.Verbose, "Both sides split by + and -:", splitLeft, splitRight)
	leftSideRawCells := parseToRawCells(splitLeft)
	rightSideRawCells := parseToRawCells(splitRight)
	rightSideRawCells = negateRawCells(rightSideRawCells)
	// fmt.Printf("\nleftSideRawCells: %+v\n\n", leftSideRawCells)
	// fmt.Printf("rightSideRawCells: %+v\n", rightSideRawCells)
	combinedRawCells := combineRawCellSlices(leftSideRawCells, rightSideRawCells)
	// fmt.Printf("\ncombinedRawCells: %+v\n", combinedRawCells)
	cells := transformRawCellsToCells(combinedRawCells)
	utils.PrintOnOption(options.Verbose, "Parsing done\n")
	return cells
}
