/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   factor_number.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/13 23:48:18 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/14 21:50:15 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "math"

func FactorNumber(input float64) []float64 {
	var factors []float64
	if input == 1 {
		return []float64{1}
	}
	var div float64 = 2
	for input > 1 {
		if math.Mod(input, div) != 0 {
			div++
		} else {
			input /= div
			factors = append(factors, div)
		}
	}
	return factors
}
