/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   calculate_discriminant.go                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:17:32 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:47:47 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

func calculateDiscriminant(cells []Cell) float64 {
	a, b, c := getABC(cells)
	return b*b - 4*a*c
}
