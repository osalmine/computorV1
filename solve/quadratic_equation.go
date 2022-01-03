/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   quadratic_equation.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:18:15 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/03 22:26:40 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"

	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/utils"
)

type equation = computorV1.Equation

func printSolveNegativeDiscriminant(computor Computor, equ equation, x complex128, substract bool) {
	log.SetFlags(0)
	// fmt.Println("X INITIAL SOLUTION:", x)
	if computor.Options.IrrFraction {
		fmt.Printf("EQU: %#v\n", equ)
		utils.PrintIrreducibleFraction(equ, substract, true)
	} else if computor.Options.SciNotation {
		log.Printf("%.6e\n", x)
	} else {
		log.Printf("%.6g\n", x)
	}
}

func solveNegativeDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly negative, no real solutions")
	if !computor.Options.Complex {
		utils.PrintOnOption(!computor.Options.Silent, "To show complex solutions, use the -c flag")
	} else {
		utils.PrintOnOption(!computor.Options.Silent, "Complex solutions:")
		x1 := (complex(-equ.B, 0) - 1i*cmplx.Sqrt(complex(-equ.Discriminant, 0))) / (2 * complex(equ.A, 0))
		x2 := (complex(-equ.B, 0) + 1i*cmplx.Sqrt(complex(-equ.Discriminant, 0))) / (2 * complex(equ.A, 0))
		printSolveNegativeDiscriminant(computor, equ, x1, true)
		printSolveNegativeDiscriminant(computor, equ, x2, false)
	}
}

func solveZeroDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly zero, the one solution is:")
	x := -equ.B / (2 * equ.A)
	if x == -0 {
		x = 0
	}
	log.SetFlags(0)
	if computor.Options.IrrFraction && x != math.Trunc(x) {
		utils.PrintTwoVarFraction(-equ.B, 2*equ.A)
	} else if computor.Options.SciNotation {
		log.Printf("%.6e\n", x)
	} else {
		log.Printf("%.6g\n", x)
	}
}

func printSolvePositiveDiscriminant(computor Computor, equ equation, x float64, substract bool) {
	log.SetFlags(0)
	// fmt.Println("X INITIAL SOLUTION:", x)
	if computor.Options.IrrFraction && x != math.Trunc(x) {
		fmt.Printf("EQU: %#v\n", equ)
		utils.PrintIrreducibleFraction(equ, substract, false)
	} else if computor.Options.SciNotation {
		log.Printf("%.6e\n", x)
	} else {
		log.Printf("%.6g\n", x)
	}
}

func solvePositiveDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly positive, the two solutions are:")
	x1 := (-equ.B - ft_sqrt(equ.Discriminant)) / (2 * equ.A)
	x2 := (-equ.B + ft_sqrt(equ.Discriminant)) / (2 * equ.A)
	printSolvePositiveDiscriminant(computor, equ, x1, true)
	printSolvePositiveDiscriminant(computor, equ, x2, false)
}

func calculateQuadraticEquation(computor Computor) {
	a, b, c := getABC(computor.Cells)
	discriminant := b*b - 4*a*c
	equation := equation{A: a, B: b, C: c, Discriminant: discriminant}
	utils.PrintOnOption(computor.Options.Verbose, "Quadratic equation")
	if discriminant < 0 {
		solveNegativeDiscriminant(computor, equation)
	} else if discriminant == 0 {
		solveZeroDiscriminant(computor, equation)
	} else {
		solvePositiveDiscriminant(computor, equation)
	}
}
