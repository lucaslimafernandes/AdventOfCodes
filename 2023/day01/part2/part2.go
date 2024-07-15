package part2

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/lucaslimafernandes/AdventOfCodes/2023/day01/utils"
)

var m = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var mr = map[string]int{
	"eno":   1,
	"owt":   2,
	"eerht": 3,
	"ruof":  4,
	"evif":  5,
	"xis":   6,
	"neves": 7,
	"thgie": 8,
	"enin":  9,
}

func Pega2(s string) (int, error) {

	var n1 int
	var n2 int
	var n int

	reversed := utils.Reverse(s)

	for i1, char1 := range s {

		if unicode.IsDigit(char1) {
			v, _ := strconv.Atoi(string(char1))
			// fmt.Printf("é numb: %v \t %v\n", s, v)
			n1 = v

		} else {

			for key := range m {

				if len(s[i1:]) < len(key) {
					continue
				}

				v, ok := m[s[i1:len(key)+i1]]
				if ok {
					// fmt.Printf("é texto: %v \t %v\n", s, v)
					n1 = v
					break
				}
			}

		}

		if n1 > 0 {
			break
		}

	}

	for i2, char2 := range reversed {

		if unicode.IsDigit(char2) {
			v, _ := strconv.Atoi(string(char2))
			n2 = v

		} else {

			for key2 := range mr {

				if len(reversed[i2:]) < len(key2) {
					continue
				}

				v, ok := mr[reversed[i2:len(key2)+i2]]
				if ok {
					n2 = v
					break
				}
			}

		}

		if n2 > 0 {
			break
		}

	}

	n, _ = strconv.Atoi(fmt.Sprintf("%v%v", n1, n2))

	return n, nil

}
