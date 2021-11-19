/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   check_poly_degree.go                               :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 21:16:49 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 21:16:56 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import "fmt"

func checkPolynomialDegree(polynomial int) bool {
	if polynomial > 2 {
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		return false
	}
	return true
}
