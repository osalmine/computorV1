/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   gcd.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/13 23:44:48 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/04 19:32:55 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "math"

// GCD returns the greatest common divisor of two numbers
// can be combined GCD(a, GCD(b, c)) to get the GCD of three numbers etc.
func GCD(a, b float64) float64 {
	if a != math.Trunc(a) || b != math.Trunc(b) {
		return 1
	}
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}
	return a
}
