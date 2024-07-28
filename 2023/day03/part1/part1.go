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

func Game1(s [][]string) []sym {

	// symb = make(map[int]map[int]string)

	fmt.Println(symb)

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

	fmt.Println(symb)

	return symb

}

// Ver se as posições acima são digitos
func Game2(s []sym, ss [][]string) []poss {

	// var r, c int

	for _, v := range s {

		// fmt.Println(v)

		if v.Type == "IsSymbol" {
			// fmt.Println("R:", v.Row, "C:", v.Col)
			// fmt.Println(ss[v.Row][v.Col])
			// fmt.Println(string(ss[v.Row][0][v.Col]))
			// r = v.Row
			// c = v.Col
			// fmt.Println(r, c)
			for tr := -1; tr <= 1; tr++ {
				for tc := -1; tc <= 1; tc++ {

					if v.Row+tr < 0 || v.Col+tc < 0 {
						continue
					}
					temp := ss[v.Row+tr][0][v.Col+tc]
					if unicode.IsDigit(rune(temp)) {

						// fmt.Println("TR:", r+tr, "TC:", c+tc, "\t", string(temp))
						sposs = append(sposs, poss{v.Row + tr, v.Col + tc})
					}

				}

			}
		}
	}

	return sposs

}

func Game3(s []poss, ss [][]string) {

	fmt.Println(s)

	subs := ss[s[0].Row][0][s[0].Col]
	fmt.Println(string(subs))

	// Verificar anteriores e posteriores
}
