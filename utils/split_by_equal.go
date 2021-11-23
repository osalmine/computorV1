/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_equal.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:57 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 15:58:03 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "strings"

func SplitByEqual(stringToSplit string) []string {
	splittedString := strings.Split(stringToSplit, "=")
	if len(splittedString) != 2 {
		panic("Equation does not contain single =")
	}
	return splittedString
}
