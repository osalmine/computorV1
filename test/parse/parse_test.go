/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse_test.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/23 13:49:52 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/30 12:09:38 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package parse_test

import (
	"reflect"
	"testing"

	"github.com/osalmine/computorV1"
	"github.com/osalmine/computorV1/parse"
	"github.com/osalmine/computorV1/utils"
)

type Options = computorV1.Options
type Cell = computorV1.Cell
type V = computorV1.VisualRepresentation

func TestParse(t *testing.T) {
	var tests = []struct {
		input   string
		options Options
		want    []Cell
	}{
		{"3x^2 + 5*x - 4 = 0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: -4, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 5, Exponent: 1, Variable: true, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 3, Exponent: 2, Variable: true, Visual: V{CapitalX: false, DisplayExponent: true}},
				{Coefficient: -0, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
			},
		},
		{"    -5*2     *    X   ^0   + 4 *  X^ 1 -9.3*X^2 + X * X - 3^2   =  1* X ^ 0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: -9, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 1, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: false}},
				{Coefficient: -9.3, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -10, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -1, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"5 + 4 * X^1 - 9.3 * X^2 = 1 * X^0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: -9.3, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 5, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -1, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"5 + 4 * X^1 - 9 * X^2 = 1 * X^0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: -9, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 5, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -1, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"4 * X^1 - 9.3 * X^2 + 7*X^42 = 4*X^42+3*X^42", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 7, Exponent: 42, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -9.3, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -3, Exponent: 42, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -4, Exponent: 42, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"5 * X^0 + 4 * X^1 = 4 * X^0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 5, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -4, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"5 + 4 * X = 4", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: false}},
				{Coefficient: 5, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -4, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
			},
		},
		{"42 * X^0 = 42 * X^0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 42, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: -42, Exponent: 0, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"5 + 4 * X + X^2= X^2", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 1, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
				{Coefficient: 4, Exponent: 1, Variable: true, Visual: V{CapitalX: true, DisplayExponent: false}},
				{Coefficient: 5, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -1, Exponent: 2, Variable: true, Visual: V{CapitalX: true, DisplayExponent: true}},
			},
		},
		{"x^2 - x + 5 * x + 9 = 0", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 9, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 5, Exponent: 1, Variable: true, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -1, Exponent: 1, Variable: true, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 1, Exponent: 2, Variable: true, Visual: V{CapitalX: false, DisplayExponent: true}},
				{Coefficient: -0, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
			},
		},
		{"(4*6*-5) + 4*2*8x +(-5)^2 + 8^3 - (-2)*8x^2 = x^4 - x^4", Options{Verbose: false, Help: false, Silent: false, Complex: false, ShowCells: false},
			[]Cell{
				{Coefficient: 16, Exponent: 2, Variable: true, Visual: V{CapitalX: false, DisplayExponent: true}},
				{Coefficient: 512, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 25, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 64, Exponent: 1, Variable: true, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: -120, Exponent: 0, Variable: false, Visual: V{CapitalX: false, DisplayExponent: false}},
				{Coefficient: 1, Exponent: 4, Variable: true, Visual: V{CapitalX: false, DisplayExponent: true}},
				{Coefficient: -1, Exponent: 4, Variable: true, Visual: V{CapitalX: false, DisplayExponent: true}},
			},
		},
	}
	for _, test := range tests {

		testname := test.input
		t.Run(testname, func(t *testing.T) {
			ans := parse.Parse(test.input, test.options)
			if !reflect.DeepEqual(ans, test.want) {
				t.Errorf("got %s, want %s", utils.PrettyStringCells(ans), utils.PrettyStringCells(test.want))
			}
		})
	}
}
