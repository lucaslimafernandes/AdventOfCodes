package part1

import (
	"fmt"
	"sort"
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

			if !unicode.IsDigit(rune(v2)) {
				// if string(v2) == "." {
				n, _ := strconv.Atoi(temporary)

				// if n > 0 && !slices.Contains(dline, n) {
				if n > 0 {
					dline = append(dline, n)
				}
				temporary = ""
			}

			// fmt.Println("temporary", temporary)
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

	// fmt.Println(len(m))
	//range m
	for k := 0; k < len(m); k++ {

		// fmt.Println(k, m[k])

		// fmt.Println("s", s[k][0])
		tempText := s[k][0]

		// mk := m[k]
		for _, v := range m[k] {
			stop := false
			// fmt.Println("INI", stop, v)
			// fmt.Println("SSS", s[k][0])
			st := strconv.Itoa(v)
			// tempText := s[k][0]

			numberOccur := strings.Count(s[k][0], st)
			// lastIndex := 0

			for occur := 0; occur < numberOccur; occur++ {

				// está duplcando
				stop = false

				// cN := strings.Index(tempText[lastIndex:], st)
				// lastIndex = cN + len(st)
				// tempText = tempText[lastIndex:]

				// c := lastIndex
				c := strings.Index(tempText, st)
				abytes := []byte(tempText)
				// fmt.Println("C", c, "len st", len(st))

				if c < 0 {
					continue
				}
				// fmt.Println(tempText)

				copy(abytes[c:c+len(st)], []byte(fmt.Sprintf("%v", strings.Repeat(".", len(st)))))

				tempText = string(abytes)
				// fmt.Println(tempText, st, stop)
				// fmt.Println("mk", v, c, s[k][0], len(st))

				if stop {
					// fmt.Println("CONTINUE")
					continue
				}
				for char := 0; char < len(st); char++ {

					for row := -1; row <= 1; row++ {
						for col := -1; col <= 1; col++ {
							// Problema aqui
							// fmt.Println("maxRow", maxRow, "maxCol", maxCol)
							// fmt.Println("row", row, "col", col)
							// fmt.Println("k+row", k+row, "c+col", c+col)

							if k+row < 0 || c+col < 0 || k+row >= maxRow || c+col >= maxCol {
								continue
							}
							// if k > 0 && c > 0 && k+row < maxRow && c+col < maxCol {

							// if !unicode.IsDigit(rune(s[k+row][0][c+col])) && string(s[k+row][0][c+col]) != "." {
							if !unicode.IsDigit(rune(s[k+row][0][c+col+char])) && string(s[k+row][0][c+col+char]) != "." {
								if !stop {

									sumOf = sumOf + v
									sa = append(sa, v)
									// fmt.Println("VAI CONTAR AQUI", v, stop)
								}
								stop = true
							}

							// }
						}
					}
				}

			}

		}
		// fmt.Println(sumOf, k)
	}

	// ssa :=
	sort.Ints(sa)

	fmt.Println(sa)
	fmt.Println("FINAL:", sumOf)

}

func Checks2(m map[int][]int, s [][]string) {

	sumOf := 0

	maxRow := len(s)
	maxCol := len(s[0][0])
	for k := range m {

		// fmt.Println(k, m[k])

		// fmt.Println("s", s[k][0])
		tempText := s[k][0]

		// mk := m[k]
		for _, v := range m[k] {
			stop := false
			// fmt.Println("SSS", s[k][0])
			st := strconv.Itoa(v)
			// tempText := s[k][0]

			numberOccur := strings.Count(s[k][0], st)
			// lastIndex := 0

			for occur := 0; occur < numberOccur; occur++ {

				// está duplcando

				// cN := strings.Index(tempText[lastIndex:], st)
				// lastIndex = cN + len(st)
				// tempText = tempText[lastIndex:]

				// c := lastIndex
				c := strings.Index(tempText, st)
				abytes := []byte(tempText)
				// fmt.Println("C", c, "len st", len(st))

				if c < 0 {
					continue
				}
				copy(abytes[c:c+len(st)], []byte(fmt.Sprintf("%v", strings.Repeat(".", len(st)))))

				tempText = string(abytes)
				// fmt.Println("mk", v, c, s[k][0], len(st))

				if stop {
					// fmt.Println("CONTINUE")
					continue
				}
				for char := 0; char < len(st); char++ {

					for row := -1; row <= 1; row++ {
						for col := -1; col <= 1; col++ {
							// Problema aqui
							// fmt.Println("maxRow", maxRow, "maxCol", maxCol)
							// fmt.Println("row", row, "col", col)
							// fmt.Println("k+row", k+row, "c+col", c+col)

							if k+row < 0 || c+col < 0 || k+row >= maxRow || c+col >= maxCol {
								continue
							}
							// if k > 0 && c > 0 && k+row < maxRow && c+col < maxCol {

							// if !unicode.IsDigit(rune(s[k+row][0][c+col])) && string(s[k+row][0][c+col]) != "." {
							if !unicode.IsDigit(rune(s[k+row][0][c+col+char])) && string(s[k+row][0][c+col+char]) != "." {
								if !stop {

									sumOf = sumOf + v
									fmt.Println("VAI CONTAR AQUI", v, stop)
								}
								stop = true
							}

							// }
						}
					}
				}

			}

		}
	}

	fmt.Println("FINAL:", sumOf)

}
