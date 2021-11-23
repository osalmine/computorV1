/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   linear_equation.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:17:56 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/22 17:22:37 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"log"

	"github.com/osalmine/computorV1/utils"
)

func calculateLinearEquation(computor Computor) {
	utils.PrintOnOption(computor.Options.Verbose, "Linear equation")
	_, b, c := getABC(computor.Cells)
	x := -c / b
	log.SetFlags(0)
	log.Printf("%.6g\n", x)
}
