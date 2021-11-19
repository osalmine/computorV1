/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parsing.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:47 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:16:08 by osalmine         ###   ########.fr       */
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

func Parse(input string) []Cell {
	s := utils.RemoveWhitespace(input)
	// fmt.Println(s)
	leftRight := utils.SplitByEqual(s)
	// fmt.Println("leftRight:", leftRight)
	splitLeft := utils.SplitByPlusMinus(leftRight[0])
	// fmt.Println("left side:", splitLeft)
	splitRight := utils.SplitByPlusMinus(leftRight[1])
	// fmt.Println("right side:", splitRight)
	leftSideRawCells := parseToRawCells(splitLeft)
	rightSideRawCells := parseToRawCells(splitRight)
	rightSideRawCells = negateRawCells(rightSideRawCells)
	// fmt.Printf("\nleftSideRawCells: %+v\n\n", leftSideRawCells)
	// fmt.Printf("rightSideRawCells: %+v\n", rightSideRawCells)
	combinedRawCells := combineRawCellSlices(leftSideRawCells, rightSideRawCells)
	// fmt.Printf("\ncombinedRawCells: %+v\n", combinedRawCells)
	cells := transformRawCellsToCells(combinedRawCells)
	return cells
}
