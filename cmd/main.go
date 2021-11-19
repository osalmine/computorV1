/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:32 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:11:59 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"

	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/parse"
	"github.com/osalmine/computorV1/solve"
	"github.com/osalmine/computorV1/utils"
)

type Computor = computorV1.Computor

func printPolynomialDegree(computor Computor) {
	utils.PrintOnOption(!computor.Options.Silent, "Polynomial degree:", computor.PolynomialDegree)
}

func printHelp() {
	fmt.Println("   ___                            _                    _ ")
	fmt.Println("  / __\\___  _ __ ___  _ __  _   _| |_ ___  _ __/\\   /\\/ |")
	fmt.Println(" / /  / _ \\| '_ ` _ \\| '_ \\| | | | __/ _ \\| '__\\ \\ / /| |")
	fmt.Println("/ /__| (_) | | | | | | |_) | |_| | || (_) | |   \\ V / | |")
	fmt.Println("\\____/\\___/|_| |_| |_| .__/ \\__,_|\\__\\___/|_|    \\_/  |_|")
	fmt.Println("                     |_|                                 ")
	fmt.Printf("\n\n")
	fmt.Printf("ComputorV1 solves polynomial equations up to the second degree\n\n")
	fmt.Println("For example: 5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	fmt.Println("You can also input it in natural form: 5 + 4x + x^2 = 0")
	fmt.Println("And use capital or non-capital X as variable")
	fmt.Printf("\n\n")
	fmt.Println("Options:")
	fmt.Println("\n-h:\tdisplay help")
	fmt.Println("-v:\tverbose")
	fmt.Println("-c:\tdisplay complex solutions")
	fmt.Println("-s:\tdisplay only the solution(s)")
}

func main() {
	args := os.Args[1:]
	str, options := parse.ParseArgs(args)
	if options.Help {
		printHelp()
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
	cells := parse.Parse(str)
	// fmt.Println("Parsing done")
	cells = parse.CombineCells(cells)
	// fmt.Println("Combining done")
	// fmt.Printf("CELLS: %+v\n", cells)
	computor := utils.ComposeComputor(cells, options)
	if !computor.Options.Silent {
		utils.PrintReducedForm(cells)
	}
	// fmt.Printf("COMPUTOR: %+v\n", computor)
	printPolynomialDegree(computor)
	solve.Solve(computor)
}
