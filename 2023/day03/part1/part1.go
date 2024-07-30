package part1

import (
	"fmt"
	"unicode"
)

// Para verificar linha a linha e posição a posição,
// se é um simbolo, ponto ou digito
// var symb map[int]map[int]string

type sym struct {
	Row  int
	Col  int
	Type string
}

type poss struct {
	Row int
	Col int
}

var symb []sym
var sposs []poss

func SymbolsIdentifier(s [][]string) []sym {

	// symb = make(map[int]map[int]string)

	fmt.Println(s)

	fmt.Println("Rows:", len(s))

	for i, v := range s {
		fmt.Println("Row", i, "Len:", len(v[0]), v)

		for i2, v2 := range v[0] {
			// fmt.Println(i2, string(v2))
			if unicode.IsDigit(rune(v2)) {
				// fmt.Println("Row:", i, "Col:", i2, string(v2), "IsDigit")
				// symb[i] = map[int]string{i2: "IsDigit"}
				symb = append(symb, sym{
					Row:  i,
					Col:  i2,
					Type: "IsDigit",
				})
			} else if string(v2) == "." {
				// fmt.Println("Row:", i, "Col:", i2, string(v2), "IsPunct")
				// symb[i] = map[int]string{i2: "IsDigit"}
				symb = append(symb, sym{
					Row:  i,
					Col:  i2,
					Type: "IsPunct",
				})
			} else {
				// fmt.Println("Row:", i, "Col:", i2, string(v2), "IsSymbol")
				// symb[i] = map[int]string{i2: "IsDigit"}
				symb = append(symb, sym{
					Row:  i,
					Col:  i2,
					Type: "IsSymbol",
				})
			}
		}
	}
	fmt.Println("SYMB", symb)
	return symb

}

// Ver se as posições acima são digitos
func PossibilitiesIdentifier(s []sym, ss [][]string) []poss {

	// var r, c int

	for _, v := range s {

		// fmt.Println(v)

		if v.Type == "IsSymbol" {
			// fmt.Println("R:", v.Row, "C:", v.Col)
			// fmt.Println(ss[v.Row][v.Col])
			// fmt.Println(string(ss[v.Row][0][v.Col]))
			// fmt.Println(r, c)
			for tr := -1; tr <= 1; tr++ {
				inRow := false
				for tc := -1; tc <= 1; tc++ {

					fmt.Println(v.Row)
					if v.Row+tr < 0 || v.Col+tc < 0 {
						if inRow {
							sposs = append(sposs, poss{v.Row + tr, v.Col + tc})
						} else {
							continue
						}
					}

					temp := ss[v.Row+tr][0][v.Col+tc+1]
					if unicode.IsDigit(rune(temp)) {
						fmt.Println("CONTINUE:", v.Row, v.Col)
						inRow = true
						continue
					}

					temp2 := ss[v.Row+tr][0][v.Col+tc]
					if unicode.IsDigit(rune(temp2)) {

						// fmt.Println("TR:", v.Row+tr, "TC:", v.Col+tc, "\t", string(temp))
						sposs = append(sposs, poss{v.Row + tr, v.Col + tc})
					}

				}

			}
		}
	}

	// fmt.Println(sposs)
	return sposs

}

func Game3(s []poss, ss [][]string) {

	fmt.Println(s)

	var k []string
	subs := ss[s[0].Row][0][s[0].Col]
	fmt.Println(string(subs))

	for item := range s {
		srow := s[item].Row
		scol := s[item].Col
		// Verificar anteriores e posteriores
		for c := 0; c < len(ss[srow][0]); c++ {
			// fmt.Println(c)
			if unicode.IsDigit(rune(ss[srow][0][c])) {
				k = append(k, string(ss[srow][0][c]))
			}

			if c == scol {
				fmt.Println("BREKA")
				break
			}

			if string(ss[srow][0][c]) == "." {
				k = nil
			}

		}
		// quase, mudar de slice para outra estrutura
		fmt.Println("K:", k)
	}

}
