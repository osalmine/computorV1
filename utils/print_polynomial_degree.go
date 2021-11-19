/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print_polynomial_degree.go                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 22:36:29 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/19 22:37:06 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import "github.com/osalmine/computorV1"

func PrintPolynomialDegree(computor computorV1.Computor) {
	PrintOnOption(!computor.Options.Silent, "Polynomial degree:", computor.PolynomialDegree)
}
