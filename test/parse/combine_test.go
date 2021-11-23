/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   combine_test.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/23 16:15:24 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 16:25:29 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse_test

import (
	"reflect"
	"testing"

	"github.com/osalmine/computorV1/parse"
	"github.com/osalmine/computorV1/utils"
)

func TestCombine(t *testing.T) {
	var tests = []struct {
		name    string
		input   []Cell
		options Options
		want    []Cell
	}{
		{
			"Input string:    -5*2     *    X   ^0   + 4 *  X^ 1 -9.3*X^2 + X * X - 3^2   =  1* X ^ 0",
			[]Cell{
				{Coefficient: 9, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 1, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: false}},
				{Coefficient: -9.3, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -10, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -1, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
			Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: -2, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -8.3, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			ans := parse.CombineCells(test.input, test.options)
			if !reflect.DeepEqual(ans, test.want) {
				t.Errorf("got %s, want %s", utils.PrettyStringCells(ans), utils.PrettyStringCells(test.want))
			}
		})
	}
}
