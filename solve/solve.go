/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   solve.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 20:48:24 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:20:03 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/utils"
)

type Cell = computorV1.Cell
type Options = computorV1.Options
type Computor = computorV1.Computor

func Solve(computor Computor) {
	if !checkPolynomialDegree(computor.PolynomialDegree) {
		return
	}
	utils.PrintOnOption(computor.Options.Verbose, "Discriminant:", calculateDiscriminant(computor.Cells))
	// fmt.Println("SQRT discriminant: ", ft_sqrt(discriminant))
	if computor.PolynomialDegree == 2 {
		calculateQuadraticEquation(computor)
	} else if computor.PolynomialDegree == 1 {
		calculateLinearEquation(computor)
	} else {
		utils.PrintOnOption(!computor.Options.Silent, "All real numbers are the solution")
	}
}
