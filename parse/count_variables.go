/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   count_variables.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:12:18 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:12:30 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse

func countVariables(input string) int {
	var count int
	for _, char := range input {
		if char == 'x' || char == 'X' {
			count++
		}
	}
	return count
}
