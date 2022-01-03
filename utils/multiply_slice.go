/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   multiply_slice.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/13 23:06:42 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/13 23:20:23 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

func MultiplySlice(input []float64) float64 {
	var sum float64 = 1
	for _, value := range input {
		sum *= value
	}
	return sum
}
