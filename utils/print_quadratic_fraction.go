/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_quadratic_fraction.go                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/03 22:18:35 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/03 22:26:07 by osalmine         ###   ########.fr       */
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

func getUniqueValues(arr []float64) map[float64]float64 {
	dict := make(map[float64]float64)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	// fmt.Println(dict)
	return dict
}

func getNumberPairs(input map[float64]float64) map[string]float64 {
	var rootCoefficients []float64
	for key, value := range input {
		fmt.Printf("%f: %f\n", key, value)
		// if value >= 2 {
		for ; value >= 2; value -= 2 {
			fmt.Println("VALUE;", value)
			rootCoefficients = append(rootCoefficients, key)
			input[key] -= 2
			if input[key] == 0 {
				delete(input, key)
			}
		}
		// }
		fmt.Printf("%f: %f\n", key, value)
	}
	fmt.Println("rootCoefficients", rootCoefficients)
	fmt.Println("INPUT AFTER:", input)
	return map[string]float64{"rootCoefficient": MultiplySlice(rootCoefficients), "rootNumber": MultiplyMap(input)}
}

func PrintIrreducibleFraction(equ equation, substract bool, complex bool) {
	if complex {
		equ.Discriminant = -equ.Discriminant
		fmt.Printf("EQU: %#v\n", equ)
	}
	factors := FactorNumber(equ.Discriminant)
	fmt.Println("FACTORS:", factors)
	uniqueFactors := getUniqueValues(factors)
	fmt.Println("UNIQUE FACTORS:", uniqueFactors)
	simplifiedRoot := getNumberPairs(uniqueFactors)
	fmt.Println("simplifiedRoot", simplifiedRoot)
	divisor := GCD(-equ.B, GCD(simplifiedRoot["rootCoefficient"], 2*equ.A))
	fmt.Println("DIVISOR:", divisor)
	displayA := (2 * equ.A) / divisor
	displayB := -equ.B / divisor
	displayRootCo := simplifiedRoot["rootCoefficient"] / divisor
	fmt.Println("displayA:", displayA, "displayB:", displayB, "disaplayRootCo:", displayRootCo)
	var negative bool
	var upper string
	fmt.Println("SUBSTRACT:", substract)
	if substract {
		if displayRootCo != 1 && displayRootCo != -1 && simplifiedRoot["rootNumber"] == 1 {
			fmt.Printf("%f - %f = %f\n", displayB, displayRootCo, displayB-displayRootCo)
			equUpper := displayB - displayRootCo
			newDivisor := GCD(equUpper, displayA)
			displayA = displayA / newDivisor
			equUpper = equUpper / newDivisor
			upper = fmt.Sprintf("%g", equUpper)
		} else if displayRootCo != 1 && displayRootCo != -1 && simplifiedRoot["rootNumber"] != 1 {
			if displayB < 0 {
				upper = fmt.Sprintf("%g + %g√%g", -displayB, displayRootCo, simplifiedRoot["rootNumber"])
				negative = true
			} else {
				upper = fmt.Sprintf("%g - %g√%g", displayB, displayRootCo, simplifiedRoot["rootNumber"])
			}
		} else {
			if displayB < 0 {
				upper = fmt.Sprintf("%g + √%g", -displayB, simplifiedRoot["rootNumber"])
				negative = true
			} else {
				upper = fmt.Sprintf("%g - √%g", displayB, simplifiedRoot["rootNumber"])
			}
		}
	} else {
		if displayRootCo != 1 && displayRootCo != -1 && simplifiedRoot["rootNumber"] == 1 { // 5x^2 + 3x - 2 = 0
			fmt.Printf("%f + %f = %f\n", displayB, displayRootCo, displayB+displayRootCo)
			equUpper := displayB + displayRootCo
			newDivisor := GCD(equUpper, displayA)
			displayA = displayA / newDivisor
			equUpper = equUpper / newDivisor
			upper = fmt.Sprintf("%g", equUpper)
		} else if displayRootCo != 1 && displayRootCo != -1 && simplifiedRoot["rootNumber"] != 1 {
			if displayB < 0 {
				upper = fmt.Sprintf("%g - %g√%g", -displayB, displayRootCo, simplifiedRoot["rootNumber"])
				negative = true
			} else {
				upper = fmt.Sprintf("%g + %g√%g", displayB, displayRootCo, simplifiedRoot["rootNumber"])
			}
		} else {
			if displayB < 0 {
				upper = fmt.Sprintf("%g - √%g", -displayB, simplifiedRoot["rootNumber"])
				negative = true
			} else {
				upper = fmt.Sprintf("%g + √%g", displayB, simplifiedRoot["rootNumber"])
			}
		}
	}
	fmt.Println("upper:", upper)
	if complex {
		upper += "i"
	}
	// fmt.Println("len(√):", len("√"))
	if displayA != 1 {
		var width int
		if len(upper)-2 > len(fmt.Sprintf("%g", displayA)) {
			fmt.Println("LENGTH OF UPPER:", len(upper))
			width = len(upper) - 2
		} else {
			fmt.Println("LENGTH OF A:", len(fmt.Sprintf("%g", displayA)))
			width = len(fmt.Sprintf("%g", displayA))
		}
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
	} else {
		log.Println(upper)
	}
}
