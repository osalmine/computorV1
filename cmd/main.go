/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:32 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/22 17:24:22 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"

	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/parse"
	"github.com/osalmine/computorV1/solve"
	"github.com/osalmine/computorV1/utils"
)

type Computor = computorV1.Computor

func main() {
	str, options := parse.ParseArgs(os.Args[1:])
	if options.Help {
		utils.PrintHelp()
		return
	}
	// str := "    -5*2     *    X   ^0   + 4 *  X^ 1 -9.3*X^2 + X * X - 3^2   =  1* X ^ 0"
	// str := "5 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
	// str := "5 + 4 * X^1 - 9 * X^2 = 1 * X^0"
	// str := "4 * X^1 - 9.3 * X^2 + 7*X^42 = 4*X^42+3*X^42"
	// str := "5 * X^0 + 4 * X^1 = 4 * X^0"
	// str := "5 + 4 * X = 4"
	// str := "42 * X^0 = 42 * X^0"
	// str := "5 + 4 * X + X^2= X^2"
	// str := "x^2 - x + 5 * x + 9 = 0" //    DISCRIMINANT 0
	cells := parse.Parse(str, options)
	// fmt.Println("Parsing done")
	cells = parse.CombineCells(cells)
	// fmt.Println("Combining done")
	// fmt.Printf("CELLS: %+v\n", cells)
	computor := utils.ComposeComputor(cells, options)
	if !computor.Options.Silent {
		utils.PrintReducedForm(cells)
	}
	// fmt.Printf("COMPUTOR: %+v\n", computor)
	utils.PrintPolynomialDegree(computor)
	solve.Solve(computor)
}
