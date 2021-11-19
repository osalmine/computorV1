/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_power.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:48:09 by osalmine          #+#    #+#             */
/*   Updated: 2021/10/30 23:48:09 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "strings"

func SplitByPower(stringToSplit string) []string {
	splittedString := strings.Split(stringToSplit, "^")
	return splittedString
}
