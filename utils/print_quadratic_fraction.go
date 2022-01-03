/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_quadratic_fraction.go                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/03 22:18:35 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/03 23:24:27 by osalmine         ###   ########.fr       */
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
		fmt.Printf("%f: %f\n", key, value)
		for ; value >= 2; value -= 2 {
			fmt.Println("VALUE;", value)
			rootCoefficients = append(rootCoefficients, key)
			input[key] -= 2
			if input[key] == 0 {
				delete(input, key)
			}
		}
		fmt.Printf("%f: %f\n", key, value)
	}
	fmt.Println("rootCoefficients", rootCoefficients)
	fmt.Println("INPUT AFTER:", input)
	return simplifiedRoot{rootCoefficient: MultiplySlice(rootCoefficients), rootNumber: MultiplyMap(input)}
}

func transformToDisplay(equ equation) equationElements {
	factors := FactorNumber(equ.Discriminant)
	fmt.Println("FACTORS:", factors)
	uniqueFactors := getUniqueValues(factors)
	fmt.Println("UNIQUE FACTORS:", uniqueFactors)
	simplifiedRoot := getNumberPairs(uniqueFactors)
	fmt.Println("simplifiedRoot", simplifiedRoot)
	divisor := GCD(-equ.B, GCD(simplifiedRoot.rootCoefficient, 2*equ.A))
	fmt.Println("DIVISOR:", divisor)
	displayA := (2 * equ.A) / divisor
	displayB := -equ.B / divisor
	displayRootCo := simplifiedRoot.rootCoefficient / divisor
	return equationElements{displayA, displayB, displayRootCo, simplifiedRoot.rootNumber}
}

func getWidth(upper string, displayA float64) int {
	if len(upper)-2 > len(fmt.Sprintf("%g", displayA)) {
		fmt.Println("LENGTH OF UPPER:", len(upper))
		return len(upper) - 2
	} else {
		fmt.Println("LENGTH OF A:", len(fmt.Sprintf("%g", displayA)))
		return len(fmt.Sprintf("%g", displayA))
	}
}

func printEquation(upper string, displayA float64, negative bool) {
	width := getWidth(upper, displayA)
	fmt.Println("width", width)
	fmt.Println("(width + len(fmt.Sprintf(displayA))) / 2", (width+len(fmt.Sprintf("%g", displayA)))/2)
	log.Printf("(%s) / %g\n", upper, displayA)
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

func constructUpperPart(elements *equationElements, substract bool) (string, bool) {
	displayA := elements.displayA
	displayB := elements.displayB
	displayRootCo := elements.displayRootCo
	rootNumber := elements.rootNumber
	if substract {
		if displayRootCo != 1 && displayRootCo != -1 && rootNumber == 1 {
			fmt.Printf("%f - %f = %f\n", displayB, displayRootCo, displayB-displayRootCo)
			equUpper := displayB - displayRootCo
			newDivisor := GCD(equUpper, displayA)
			elements.displayA = displayA / newDivisor
			equUpper = equUpper / newDivisor
			return fmt.Sprintf("%g", equUpper), false
		} else if displayRootCo != 1 && displayRootCo != -1 && rootNumber != 1 {
			if displayB < 0 {
				return fmt.Sprintf("%g + %g√%g", -displayB, displayRootCo, rootNumber), true
			} else {
				return fmt.Sprintf("%g - %g√%g", displayB, displayRootCo, rootNumber), false
			}
		} else {
			if displayB < 0 {
				return fmt.Sprintf("%g + √%g", -displayB, rootNumber), true
			} else {
				return fmt.Sprintf("%g - √%g", displayB, rootNumber), false
			}
		}
	} else {
		if displayRootCo != 1 && displayRootCo != -1 && rootNumber == 1 { // 5x^2 + 3x - 2 = 0
			fmt.Printf("%f + %f = %f\n", displayB, displayRootCo, displayB+displayRootCo)
			equUpper := displayB + displayRootCo
			newDivisor := GCD(equUpper, displayA)
			elements.displayA = displayA / newDivisor
			equUpper = equUpper / newDivisor
			return fmt.Sprintf("%g", equUpper), false
		} else if displayRootCo != 1 && displayRootCo != -1 && rootNumber != 1 {
			if displayB < 0 {
				return fmt.Sprintf("%g - %g√%g", -displayB, displayRootCo, rootNumber), true
			} else {
				return fmt.Sprintf("%g + %g√%g", displayB, displayRootCo, rootNumber), false
			}
		} else {
			if displayB < 0 {
				return fmt.Sprintf("%g - √%g", -displayB, rootNumber), true
			} else {
				return fmt.Sprintf("%g + √%g", displayB, rootNumber), false
			}
		}
	}
}

func PrintIrreducibleFraction(equ equation, substract bool, complex bool) {
	if complex {
		equ.Discriminant = -equ.Discriminant
		fmt.Printf("EQU: %#v\n", equ)
	}
	elements := transformToDisplay(equ)
	fmt.Println("displayA:", elements.displayA, "displayB:", elements.displayB, "disaplayRootCo:", elements.displayRootCo)
	var negative bool
	var upper string
	fmt.Println("SUBSTRACT:", substract)
	upper, negative = constructUpperPart(&elements, substract)
	fmt.Println("upper:", upper)
	if complex {
		upper += "i"
	}
	if elements.displayA != 1 {
		printEquation(upper, elements.displayA, negative)
	} else {
		log.Println(upper)
	}
}
