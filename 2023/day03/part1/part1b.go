package part1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Numbers(s [][]string) map[int][]int {

	dig := make(map[int][]int)

	for line, v := range s {
		var dline []int
		temporary := ""
		for _, v2 := range v[0] {
			if unicode.IsDigit(rune(v2)) {
				// fmt.Println(line, string(v2))
				temporary = fmt.Sprintf("%s%s", temporary, string(v2))
			}

			if !unicode.IsDigit(rune(v2)) {
				n, _ := strconv.Atoi(temporary)

				if n > 0 {
					dline = append(dline, n)
				}
				temporary = ""
			}

		}
		fmt.Println("dline", dline)
		dig[line] = dline
	}

	// fmt.Println(dig)
	return dig

}

func Checks(m map[int][]int, s [][]string) {

	sumOf := 0

	maxRow := len(s)
	maxCol := len(s[0][0])

	var sa []int

	for k := 0; k < len(m); k++ {

		tempText := s[k][0]

		for _, v := range m[k] {
			stop := false
			st := strconv.Itoa(v)

			numberOccur := strings.Count(s[k][0], st)

			for occur := 0; occur < numberOccur; occur++ {

				stop = false

				c := strings.Index(tempText, st)
				abytes := []byte(tempText)

				if c < 0 {
					continue
				}
				copy(abytes[c:c+len(st)], []byte(fmt.Sprintf("%v", strings.Repeat(".", len(st)))))

				tempText = string(abytes)

				if stop {
					continue
				}
				for char := 0; char < len(st); char++ {

					for row := -1; row <= 1; row++ {
						for col := -1; col <= 1; col++ {

							if k+row < 0 || c+col < 0 || k+row >= maxRow || c+col >= maxCol {
								continue
							}

							// if !unicode.IsDigit(rune(s[k+row][0][c+col])) && string(s[k+row][0][c+col]) != "." {
							if !unicode.IsDigit(rune(s[k+row][0][c+col+char])) && string(s[k+row][0][c+col+char]) != "." {

								if !stop {
									sumOf = sumOf + v
									sa = append(sa, v)
									// fmt.Println("VAI CONTAR AQUI", v, stop)
								}
								stop = true
							}
						}
					}
				}

			}

		}
		// fmt.Println(sumOf, k) // soma linha a linha
	}

	sort.Ints(sa)

	fmt.Println(sa)
	fmt.Println("FINAL:", sumOf)

}
