/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   linear_equation.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:17:56 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/11 12:43:27 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"log"

	"github.com/osalmine/computorV1/utils"
)

func calculateLinearEquation(computor Computor) {
	utils.PrintOnOption(computor.Options.Verbose, "Linear equation")
	utils.PrintOnOption(!computor.Options.Silent, "The solution is:")
	_, b, c := getABC(computor.Cells)
	x := -c / b
	if x == -0 {
		x = 0
	}
	log.SetFlags(0)
	if computor.Options.SciNotation {
		log.Printf("%.6e\n", x)
	} else {
		log.Printf("%.6g\n", x)
	}
}
