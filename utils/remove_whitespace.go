/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   remove_whitespace.go                               :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/10/30 23:47:51 by osalmine          #+#    #+#             */
/*   Updated: 2021/10/30 23:47:52 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "strings"

func RemoveWhitespace(inputString string) string {
	output := strings.Join(strings.Fields(inputString), "")
	return output
}
