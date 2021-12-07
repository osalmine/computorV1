/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_help.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 22:37:20 by osalmine          #+#    #+#             */
/*   Updated: 2021/12/07 15:15:57 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "fmt"

func PrintHelp() {
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
	fmt.Println("and use capital or non-capital X as variable")
	fmt.Printf("\n\n")
	fmt.Println("Options:")
	fmt.Println("\n-h:\tdisplay help")
	fmt.Println("-v:\tverbose")
	fmt.Println("-c:\tdisplay complex solutions")
	fmt.Println("-s:\tdisplay only the solution(s)")
	fmt.Println("-d:\tdisplay cells")
	fmt.Println("-n:\tuse scientific notation")
}
