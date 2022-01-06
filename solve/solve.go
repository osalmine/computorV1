/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   solve.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 20:48:24 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/06 22:39:54 by osalmine         ###   ########.fr       */
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
	if computor.PolynomialDegree == 2 {
		utils.PrintOnOption(computor.Options.Verbose, "Discriminant:", calculateDiscriminant(computor.Cells))
		calculateQuadraticEquation(computor)
	} else if computor.PolynomialDegree == 1 {
		calculateLinearEquation(computor)
	} else {
		calculateNoVariables(computor)
	}
}
