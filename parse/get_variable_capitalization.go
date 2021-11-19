/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_variable_capitalization.go                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:13:50 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:13:55 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

import "strings"

func getVariableCapitalization(str string) bool {
	if !strings.Contains(str, "x") && !strings.Contains(str, "X") {
		return false
	}
	return strings.Contains(str, "X")
}
