/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_exponents.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:13:20 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:47:05 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"strconv"
	"strings"

	"github.com/osalmine/computorV1/utils"
)

func getExponents(input []string) []int {
	var exponents []int
	for _, str := range input {
		if strings.Contains(str, "x") || strings.Contains(str, "X") {
			splittedByPower := utils.SplitByPower(str)
			if len(splittedByPower) > 1 {
				exponent, _ := strconv.Atoi(splittedByPower[1])
				exponents = append(exponents, exponent)
			} else {
				exponents = append(exponents, 1)
			}
		}
	}
	return exponents
}
