/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse_args.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 18:03:01 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 20:56:20 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import (
	"bufio"
	"fmt"
	"os"

	"github.com/osalmine/computorV1"
)

type Options = computorV1.Options

func ParseArgs(args []string) (string, Options) {
	var options Options
	var input string
	for _, arg := range args {
		if arg == "-v" {
			options.Verbose = true
		} else if arg == "-h" {
			options.Help = true
		} else if arg == "-s" {
			options.Silent = true
		} else if arg == "-c" {
			options.Complex = true
		} else {
			input = arg
		}
	}
	// fmt.Println("len(input)", len(input))
	// if len(input) == 0 && !options.help {
	// 	// fmt.Scanln(&input)
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	for scanner.Scan() {
	// 		input = scanner.Text()
	// 		break
	// 	}
	// }
	if len(input) == 0 && !options.Help {
		// fmt.Scanln(&input)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a polynomial equation: ")

		readerInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("\n")
		input = readerInput
		// input = scanner.Text()
		// for scanner.Scan() {
		// 	input = scanner.Text()
		// 	break
		// }
	}
	// fmt.Println(input)
	return input, options
}
