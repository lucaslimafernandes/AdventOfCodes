package part1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// func Numbers(s [][]string) {

// 	// var sumies []int

// 	for i, v := range s {
// 		fmt.Println("Row", i, "Len:", len(v[0]), v)

// 		var lineDigit map[int]int
// 		for i2, v2 := range v[0] {

// 			if unicode.IsDigit(rune(v2)) {}
// 			fmt.Println(string(v2))
// 		}

// 	}
// }

func Numbers(s [][]string) map[int][]int {

	// var dig map[int][]int
	dig := make(map[int][]int)

	for line, v := range s {
		var dline []int
		temporary := ""
		for _, v2 := range v[0] {
			if unicode.IsDigit(rune(v2)) {
				// fmt.Println(line, string(v2))
				temporary = fmt.Sprintf("%s%s", temporary, string(v2))
			}

			if string(v2) == "." {
				n, _ := strconv.Atoi(temporary)

				if n > 0 {
					dline = append(dline, n)
				}
				temporary = ""
			}

			// fmt.Println("temporary", temporary)
		}
		// fmt.Println("dline", dline)
		dig[line] = dline
	}

	fmt.Println(dig)
	return dig

}

func Checks(m map[int][]int, s [][]string) {

	maxRow := len(s)
	maxCol := len(s[0][0])
	for k := range m {
		fmt.Println(k, m[k])

		fmt.Println("s", s[k][0])

		// mk := m[k]
		for _, v := range m[k] {
			// fmt.Println("SSS", s[k][0])
			st := strconv.Itoa(v)
			c := strings.Index(s[k][0], st)
			fmt.Println("mk", v, c, s[k][0], len(st))

			for row := -1; row <= 1; row++ {
				for col := -1; col <= 1; col++ {
					// Problema aqui
					if k > 0 && c > 0 && k < maxRow && c < maxCol {

						if !unicode.IsDigit(rune(s[k+row][0][c+col])) && string(s[k+row][0][c+col]) != "." {
							fmt.Println("VAI CONTAR AQUI")
						}

					}
				}
			}

		}

	}

}
