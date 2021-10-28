package main

import (
	"fmt"
	"regexp"

	"github.com/osalmine/computorV1/utils"
)

type Coefficient struct {
	sign  string
	value float64
}

type RawCell struct {
	variableCount int
	sign          rune
	exponents     []int
	coefficients  []Coefficient
}

func countVariables(input string) int {
	var count int
	for _, char := range input {
		if char == 'x' || char == 'X' {
			count++
		}
	}
	fmt.Println("VARIABLE COUNT", count)
	return count
}

func getSign(input string) rune {
	if input[0] == '-' {
		fmt.Println("NEGATIVE SIGN")
		return '-'
	}
	fmt.Println("POSITIVE SIGN")
	return '+'
}

func parseToRawCells(input []string, negateCell bool) {
	fmt.Println("INPUT", input)
	var RawCells []RawCell
	for _, str := range input {
		fmt.Println("STR", str)
		RawCells = append(RawCells, RawCell{
			variableCount: countVariables(str),
			sign:          getSign(str),
		})
	}
}

func parse(input string) {
	s := utils.RemoveWhitespace(input)
	fmt.Println(s)
	content := []byte(s)
	re := regexp.MustCompile(`\^`).FindAllIndex(content, -1)
	fmt.Println("POWERS:", re)
	for i, val := range re {
		fmt.Println("INDEX:", i, "strIndex 0:", string(s[val[0]]), "strIndex 1:", string(s[val[1]]))
	}
	// for i, strIndex := range re {
	// 	fmt.Println("INDEX:", i, "strIndex:", string(s[strIndex]), "strIndex + 1:", string(s[strIndex]+1))
	// }
	// powers := strings.Index(s, "^")
	// fmt.Println("POWERS:", powers)
	leftRight := utils.SplitByEqual(s)
	fmt.Println("leftRight:", leftRight)
	splitRight := utils.SplitByPlusMinus(leftRight[0])
	fmt.Println("right side:", splitRight)
	splitLeft := utils.SplitByPlusMinus(leftRight[1])
	fmt.Println("left side:", splitLeft)
	parseToRawCells(splitRight, false)
}
