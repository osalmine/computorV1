/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_quadratic_fraction.go                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/03 22:18:35 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/08 18:06:25 by osalmine         ###   ########.fr       */
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
	sqrtSymbolExtraLen := 2
	if !strings.Contains(upper, "√") {
		sqrtSymbolExtraLen = 0
	}
	if len(upper)-sqrtSymbolExtraLen > len(fmt.Sprintf("%g", displayA)) {
		return len(upper) - sqrtSymbolExtraLen
	} else {
		return len(fmt.Sprintf("%g", displayA))
	}
}

func getPadding(upper string, displayA float64, width int) (int, int) {
	var upperPadding int
	var lowerPadding int
	sqrtSymbolExtraLen := 2
	if !strings.Contains(upper, "√") {
		sqrtSymbolExtraLen = 0
	}
	if len(upper)-sqrtSymbolExtraLen > len(fmt.Sprintf("%g", displayA)) {
		upperPadding = 0
		lowerPadding = (width + len(fmt.Sprintf("%g", displayA))) / 2
		return upperPadding, lowerPadding
	} else {
		upperPadding = (width + len(upper)) / 2
		lowerPadding = 0
		return upperPadding, lowerPadding
	}
}

func printEquation(upper string, displayA float64, negative bool) {
	width := getWidth(upper, displayA)
	upperPadding, lowerPadding := getPadding(upper, displayA, width)
	if negative {
		log.Printf("  %*s", upperPadding, upper)
		log.Println("-", strings.Repeat("-", width))
		log.Printf("  %*g\n", lowerPadding, displayA)
	} else {
		log.Printf("%*s", upperPadding, upper)
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
	if rootNumber == 1 {
		equUpper := displayB + displayRootCo
		if substract {
			equUpper = displayB - displayRootCo
		}
		negative := equUpper < 0
		if negative {
			equUpper *= -1
		}
		newDivisor := GCD(equUpper, displayA)
		elements.displayA = displayA / newDivisor
		equUpper = equUpper / newDivisor
		return fmt.Sprintf("%g", equUpper), negative
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
