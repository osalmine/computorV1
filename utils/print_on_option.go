/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_on_option.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 13:29:50 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 13:34:56 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "fmt"

func PrintOnOption(option bool, text string, rest ...interface{}) {
	if option {
		fmt.Print(text)
		fmt.Print(" ")
		fmt.Print(rest...)
		fmt.Print("\n")
	}
}
