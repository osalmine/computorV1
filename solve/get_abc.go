/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_abc.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:17:07 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:17:13 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

func getABC(cells []Cell) (float64, float64, float64) {
	var a, b, c float64
	for _, cell := range cells {
		if cell.Exponent == 2 {
			a = cell.Coefficient
		} else if cell.Exponent == 1 {
			b = cell.Coefficient
		} else {
			c = cell.Coefficient
		}
	}
	return a, b, c
}
