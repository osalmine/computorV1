/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sqrt.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 20:48:51 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 20:48:54 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

func ft_sqrt(input float64) float64 {
	const PRECISION = 0.000000000000001
	var output float64
	var bigger float64
	var smaller float64
	output = input / 2
	for i := 0; i < 10; i++ {
		if input/output >= output {
			bigger = input / output
		} else {
			bigger = output
		}
		if input/output < output {
			smaller = input / output
		} else {
			smaller = output
		}
		if (bigger - smaller) < PRECISION {
			return output
		}
		output = (output + input/output) / 2
	}
	return output
}
