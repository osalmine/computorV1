/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_quadratic_fraction.go                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/03 22:18:35 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/04 19:35:40 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/osalmine/computorV1"
)

type equation = computorV1.Equation
type simplifiedRoot struct {
	rootCoefficient float64
	rootNumber      float64
}
type equationElements struct {
	displayA, displayB, displayRootCo, rootNumber float64
}

func getUniqueValues(arr []float64) map[float64]float64 {
	dict := make(map[float64]float64)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}

func getNumberPairs(input map[float64]float64) simplifiedRoot {
	var rootCoefficients []float64
	for key, value := range input {
		for ; value >= 2; value -= 2 {
			rootCoefficients = append(rootCoefficients, key)
			input[key] -= 2
			if input[key] == 0 {
				delete(input, key)
			}
		}
	}
	return simplifiedRoot{rootCoefficient: MultiplySlice(rootCoefficients), rootNumber: MultiplyMap(input)}
}

func transformToDisplay(equ equation) equationElements {
	factors := FactorNumber(equ.Discriminant)
	uniqueFactors := getUniqueValues(factors)
	simplifiedRoot := getNumberPairs(uniqueFactors)
	divisor := GCD(-equ.B, GCD(simplifiedRoot.rootCoefficient, 2*equ.A))
	displayA := (2 * equ.A) / divisor
	displayB := -equ.B / divisor
	displayRootCo := simplifiedRoot.rootCoefficient / divisor
	return equationElements{displayA, displayB, displayRootCo, simplifiedRoot.rootNumber}
}

func getWidth(upper string, displayA float64) int {
	if len(upper)-2 > len(fmt.Sprintf("%g", displayA)) {
		return len(upper) - 2
	} else {
		return len(fmt.Sprintf("%g", displayA))
	}
}

func printEquation(upper string, displayA float64, negative bool) {
	width := getWidth(upper, displayA)
	lowerPadding := (width + len(fmt.Sprintf("%g", displayA))) / 2
	if negative {
		log.Println(" ", upper)
		log.Println("-", strings.Repeat("-", width))
		log.Printf("  %*g\n", lowerPadding, displayA)
	} else {
		log.Println(upper)
		log.Println(strings.Repeat("-", width))
		log.Printf("%*g\n", lowerPadding, displayA)
	}
}

func ternary(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

func constructUpperPart(elements *equationElements, substract bool) (string, bool) {
	displayA := elements.displayA
	displayB := elements.displayB
	displayRootCo := elements.displayRootCo
	rootNumber := elements.rootNumber
	if displayRootCo != 1 && displayRootCo != -1 && rootNumber == 1 { // 5x^2 + 3x - 2 = 0
		equUpper := displayB + displayRootCo
		if substract {
			equUpper = displayB - displayRootCo
		}
		newDivisor := GCD(equUpper, displayA)
		elements.displayA = displayA / newDivisor
		equUpper = equUpper / newDivisor
		return fmt.Sprintf("%g", equUpper), false
	} else if displayRootCo != 1 && displayRootCo != -1 && rootNumber != 1 {
		if displayB < 0 {
			return fmt.Sprintf("%g %s %g√%g", -displayB, ternary(substract, "+", "-"), displayRootCo, rootNumber), true
		} else {
			return fmt.Sprintf("%g %s %g√%g", displayB, ternary(substract, "-", "+"), displayRootCo, rootNumber), false
		}
	} else {
		if displayB < 0 {
			return fmt.Sprintf("%g %s √%g", -displayB, ternary(substract, "+", "-"), rootNumber), true
		} else {
			return fmt.Sprintf("%g %s √%g", displayB, ternary(substract, "-", "+"), rootNumber), false
		}
	}
}

func PrintIrreducibleFraction(equ equation, substract bool, complex bool) {
	if complex {
		equ.Discriminant = -equ.Discriminant
	}
	elements := transformToDisplay(equ)
	upper, negative := constructUpperPart(&elements, substract)
	if complex {
		upper += "i"
	}
	if elements.displayA != 1 {
		printEquation(upper, elements.displayA, negative)
	} else {
		log.Println(upper)
	}
}
