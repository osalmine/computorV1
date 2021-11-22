/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   solve_test.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 22:40:05 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/22 17:21:26 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve_test

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/solve"
)

type Computor = computorV1.Computor
type Options = computorV1.Options
type Cell = computorV1.Cell
type VisualRepresentation = computorV1.VisualRepresentation

// func readOutput

func getCells(secondDegree, firstDegree, zeroDegree float64) []Cell {
	var cells []Cell
	if secondDegree != 0 {
		cells = append(cells, Cell{
			Coefficient: secondDegree,
			Exponent:    2,
			Variable:    true,
			Visual: VisualRepresentation{
				CapitalX:        false,
				DisplayExponent: false,
			},
		})
	}
	if firstDegree != 0 {
		cells = append(cells, Cell{
			Coefficient: firstDegree,
			Exponent:    1,
			Variable:    true,
			Visual: VisualRepresentation{
				CapitalX:        false,
				DisplayExponent: false,
			},
		})
	}
	if zeroDegree != 0 {
		cells = append(cells, Cell{
			Coefficient: zeroDegree,
			Exponent:    0,
			Variable:    false,
			Visual: VisualRepresentation{
				CapitalX:        false,
				DisplayExponent: false,
			},
		})
	}
	return cells
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}

func TestSolveLinear(t *testing.T) {
	computorT1 := Computor{
		PolynomialDegree: 1,
		Options: Options{
			Verbose: false,
			Help:    false,
			Silent:  true,
			Complex: false,
		},
		Cells: getCells(0, 3, 5),
	}
	output := captureOutput(func() {
		solve.Solve(computorT1)
	})
	output = strings.TrimSuffix(output, "\n")
	// fmt.Printf("OUTPUT: |%s|\n", output)
	if output != "-1.66667" {
		t.Error("Failure!")
	}
}

func TestSolveQuadratic(t *testing.T) {
	computorT2 := Computor{
		PolynomialDegree: 2,
		Options: Options{
			Verbose: false,
			Help:    false,
			Silent:  true,
			Complex: false,
		},
		Cells: getCells(1, 5, 5),
	}
	output := captureOutput(func() {
		solve.Solve(computorT2)
	})
	output = strings.TrimSuffix(output, "\n")
	// fmt.Printf("OUTPUT: |%s|\n", output)
	if output != "-3.61803\n-1.38197" {
		t.Error("Failure!")
	}
}

func TestSolveComplex(t *testing.T) {
	computorT2 := Computor{
		PolynomialDegree: 2,
		Options: Options{
			Verbose: false,
			Help:    false,
			Silent:  true,
			Complex: true,
		},
		Cells: getCells(1, 3, 5),
	}
	output := captureOutput(func() {
		solve.Solve(computorT2)
	})
	output = strings.TrimSuffix(output, "\n")
	// fmt.Printf("OUTPUT: |%s|\n", output)
	if output != "(-1.5-1.65831i)\n(-1.5+1.65831i)" {
		t.Error("Failure!")
	}
}
