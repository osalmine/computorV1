/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   negate_raw_cells.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:14:11 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:14:20 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

func negateRawCells(rawCells []RawCell) []RawCell {
	for i, cell := range rawCells {
		for j, coefficient := range cell.Coefficients {
			rawCells[i].Coefficients[j] = -coefficient
		}
		if len(cell.Coefficients) == 0 {
			rawCells[i].Coefficients = append(rawCells[i].Coefficients, -1)
		}
	}
	return rawCells
}
