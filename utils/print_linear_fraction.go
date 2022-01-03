/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_linear_fraction.go                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/03 22:19:20 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/03 23:42:29 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func getFactorCharWidth(c, b float64) int {
	var width int
	if len(fmt.Sprintf("%g", c)) > len(fmt.Sprintf("%g", b)) {
		width = len(fmt.Sprintf("%g", c))
	} else {
		width = len(fmt.Sprintf("%g", b))
	}
	return width
}

func getPrintPadding(width int, nb float64) int {
	return (width + len(fmt.Sprintf("%g", nb))) / 2
}

func PrintTwoVarFraction(upper, lower float64) {
	divisor := GCD(upper, lower)
	if divisor != 1 && divisor != -1 {
		upper /= divisor
		lower /= divisor
	}
	if upper < 0 && lower < 0 {
		width := getFactorCharWidth(math.Abs(upper), math.Abs(lower))
		log.Printf("%*g", getPrintPadding(width, math.Abs(upper)), math.Abs(upper))
		log.Printf("%s\n", strings.Repeat("-", width))
		log.Printf("%*g", getPrintPadding(width, math.Abs(lower)), math.Abs(lower))
	} else if upper < 0 || lower < 0 {
		width := getFactorCharWidth(math.Abs(upper), math.Abs(lower))
		log.Printf("  %*g", getPrintPadding(width, math.Abs(upper)), math.Abs(upper))
		log.Printf("- %s\n", strings.Repeat("-", width))
		log.Printf("  %*g", getPrintPadding(width, math.Abs(lower)), math.Abs(lower))
	} else {
		width := getFactorCharWidth(upper, lower)
		log.Printf("%*g", getPrintPadding(width, upper), upper)
		log.Printf("%s\n", strings.Repeat("-", width))
		log.Printf("%*g", getPrintPadding(width, lower), lower)
	}
}
