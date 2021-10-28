package main

import (
	"fmt"
)

type Cell struct {
	sign        rune
	coefficient float64
	exponent    int
	variable    bool
}

func init() {
	fmt.Println("INIT")
}

func main() {
	str := "    5     *    X   ^0   + 4 *  X^ 1 -9.3*X^2   =  1* X ^ 0"
	parse(str)
}
