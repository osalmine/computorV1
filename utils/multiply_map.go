/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   multiply_map.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/13 23:09:07 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/13 23:20:11 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

func MultiplyMap(input map[float64]float64) float64 {
	var sum float64 = 1
	for key, value := range input {
		for ; value > 0; value-- {
			sum *= key
		}
	}
	return sum
}
