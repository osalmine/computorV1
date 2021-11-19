/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_reduced_form.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:39 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:00:51 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"fmt"
	"sort"

	"github.com/osalmine/computorV1"
)

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *computorV1.Cell) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(cells []computorV1.Cell) {
	ps := &cellSorter{
		cells: cells,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type cellSorter struct {
	cells []computorV1.Cell
	by    func(p1, p2 *computorV1.Cell) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *cellSorter) Len() int {
	return len(s.cells)
}

// Swap is part of sort.Interface.
func (s *cellSorter) Swap(i, j int) {
	s.cells[i], s.cells[j] = s.cells[j], s.cells[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *cellSorter) Less(i, j int) bool {
	return s.by(&s.cells[i], &s.cells[j])
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func sortCells(cells []computorV1.Cell) []computorV1.Cell {
	exponent := func(c1, c2 *computorV1.Cell) bool {
		return c1.Exponent < c2.Exponent
	}

	By(exponent).Sort(cells)
	// fmt.Printf("Sorted: %+v\n", cells)
	return cells
}

func getSignedCoefficientString(input float64, index int) string {
	if input < 0 {
		return fmt.Sprintf("- %g ", -input)
	}
	if index == 0 {
		return fmt.Sprintf("%g ", input)
	}
	return fmt.Sprintf("+ %g ", input)
}

func getVariableAndExponent(cell computorV1.Cell, existingCellsString string) string {
	var cellsString string
	if len(existingCellsString) > 0 && existingCellsString[len(existingCellsString)-1] != ' ' {
		cellsString += " "
	}
	if cell.Visual.CapitalX {
		cellsString += "* X"
	} else {
		cellsString += "* x"
	}
	if cell.Visual.DisplayExponent {
		cellsString += fmt.Sprintf("^%d ", cell.Exponent)
	}
	if cellsString[len(cellsString)-1] != ' ' {
		cellsString += " "
	}
	return cellsString
}

func cellsToString(cells []computorV1.Cell) string {
	var cellsString string
	for i, cell := range cells {
		// fmt.Printf("%+v\n", cell)
		cellsString += getSignedCoefficientString(cell.Coefficient, i)
		if cell.Variable {
			cellsString += getVariableAndExponent(cell, cellsString)
		}
	}
	if len(cellsString) == 0 {
		cellsString = "0 "
	}
	cellsString += "= 0"
	return cellsString
}

func PrintReducedForm(cells []computorV1.Cell) {
	sortedCells := sortCells(cells)
	fmt.Printf("Reduced form: %+v\n", cellsToString(sortedCells))
}
