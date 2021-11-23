/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:32 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 15:42:13 by osalmine         ###   ########.fr       */
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
	cells := parse.Parse(str, options)
	cells = parse.CombineCells(cells, options)
	computor := utils.ComposeComputor(cells, options)
	utils.PrintReducedForm(computor)
	utils.PrintPolynomialDegree(computor)
	solve.Solve(computor)
}
