package part1

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/lucaslimafernandes/AdventOfCodes/2023/day01/utils"
)

func Pega(s string) (int, error) {

	var n1 int
	var n2 int
	var n int

	reversed := utils.Reverse(s)

	for _, char := range s {
		if unicode.IsDigit(char) {
			n1, _ = strconv.Atoi(string(char))
			break
		}
	}

	for _, char2 := range reversed {
		if unicode.IsDigit(char2) {
			n2, _ = strconv.Atoi(string(char2))
			break
		}
	}

	n, _ = strconv.Atoi(fmt.Sprintf("%v%v", n1, n2))

	return n, nil

}
