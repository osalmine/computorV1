/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split_by_equal.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:57 by osalmine          #+#    #+#             */
/*   Updated: 2021/10/30 23:47:57 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "strings"

func SplitByEqual(stringToSplit string) []string {
	splittedString := strings.Split(stringToSplit, "=")
	return splittedString
}