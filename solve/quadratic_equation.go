/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   quadratic_equation.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:18:15 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:27:00 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"fmt"
	"math/cmplx"

	"github.com/osalmine/computorV1/utils"
)

type equation struct {
	a, b, c, discriminant float64
}

func solveNegativeDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly negative, no real solutions")
	if !computor.Options.Complex {
		utils.PrintOnOption(!computor.Options.Silent, "To show complex solutions, use the -c flag")
	} else {
		utils.PrintOnOption(!computor.Options.Silent, "Complex solutions:")
		x1 := (complex(-equ.b, 0) - 1i*cmplx.Sqrt(complex(-equ.discriminant, 0))) / (2 * complex(equ.a, 0))
		x2 := (complex(-equ.b, 0) + 1i*cmplx.Sqrt(complex(-equ.discriminant, 0))) / (2 * complex(equ.a, 0))
		fmt.Printf("%.6g\n%.6g\n", x1, x2)
	}
}

func solveZeroDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly zero, the one solution is:")
	x := -equ.b / (2 * equ.a)
	fmt.Printf("%.6g\n", x)
}

func solvePositiveDiscriminant(computor Computor, equ equation) {
	utils.PrintOnOption(!computor.Options.Silent, "Discriminant is strictly positive, the two solutions are:")
	x1 := (-equ.b - ft_sqrt(equ.discriminant)) / (2 * equ.a)
	x2 := (-equ.b + ft_sqrt(equ.discriminant)) / (2 * equ.a)
	fmt.Printf("%.6g\n%.6g\n", x1, x2)
}

func calculateQuadraticEquation(computor Computor) {
	a, b, c := getABC(computor.Cells)
	discriminant := b*b - 4*a*c
	equation := equation{a, b, c, discriminant}
	if discriminant < 0 {
		solveNegativeDiscriminant(computor, equation)
	} else if discriminant == 0 {
		solveZeroDiscriminant(computor, equation)
	} else {
		solvePositiveDiscriminant(computor, equation)
	}
}
