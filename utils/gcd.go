/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   gcd.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/13 23:44:48 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/15 21:55:38 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "math"

// GCD returns the greatest common divisor of two numbers
// can be combined GCD(a, GCD(b, c)) to get the GCD of three numbers etc.
func GCD(a, b float64) float64 {
	for b != 0 {
		t := b
		b = math.Mod(a, b)
		a = t
	}
	return a
}
