/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   linear_equation.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:17:56 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:18:01 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import (
	"fmt"

	"github.com/osalmine/computorV1/utils"
)

func calculateLinearEquation(computor Computor) {
	// fmt.Println("Linear equation")
	utils.PrintOnOption(computor.Options.Verbose, "Linear equation")
	_, b, c := getABC(computor.Cells)
	x := -c / b
	fmt.Printf("%.6g\n", x)
}
