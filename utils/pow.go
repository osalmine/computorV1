/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   pow.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:11:34 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/30 12:10:34 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

func Pow(x float64, y int) float64 {
	var output float64 = 1
	for i := 1; i <= y; i++ {
		output *= x
	}
	return output
}
