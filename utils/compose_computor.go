/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   compose_computor.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:26 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 20:59:59 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "github.com/osalmine/computorV1"

func calculatePolynomialDegree(cells []computorV1.Cell) int {
	var polyDegree int
	for _, cell := range cells {
		if cell.Exponent > polyDegree {
			polyDegree = cell.Exponent
		}
	}
	return polyDegree
}

func ComposeComputor(cells []computorV1.Cell, options computorV1.Options) computorV1.Computor {
	return computorV1.Computor{
		Cells:            cells,
		PolynomialDegree: calculatePolynomialDegree(cells),
		Options:          options,
	}
}
