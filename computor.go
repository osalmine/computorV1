/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   computor.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/19 20:46:54 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 13:39:00 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package computorV1

type Computor struct {
	Cells            []Cell
	PolynomialDegree int
	Options          Options
}

type Cell struct {
	Coefficient float64
	Exponent    int
	Variable    bool
	Visual      VisualRepresentation
}

type VisualRepresentation struct {
	CapitalX        bool
	DisplayExponent bool
}

type Options struct {
	Verbose   bool
	Help      bool
	Silent    bool
	Complex   bool
	ShowCells bool
}

type RawCell struct {
	VariableCount int
	Exponents     []int
	Coefficients  []float64
	Visual        VisualRepresentation
}
