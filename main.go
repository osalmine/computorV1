package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/osalmine/computorV1/utils"
)

func main() {
	str := "    5     *    X   ^0   + 4 *  X^ 1 -9.3*X^2   =  1* X ^ 0"

	s := strings.Join(strings.Fields(str), "")
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
	leftRight := strings.Split(s, "=")
	fmt.Println("leftRight:", leftRight)
	splitRight := strings.FieldsFunc(leftRight[0], utils.Splitter)
	fmt.Println("right side:", splitRight)
	splitLeft := strings.FieldsFunc(leftRight[1], utils.Splitter)
	fmt.Println("left side:", splitLeft)
	fmt.Println(s[len(s)-1])
}
