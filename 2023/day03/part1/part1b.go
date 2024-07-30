package part1

import "fmt"

func Numbers(s [][]string) {

	for i, v := range s {
		fmt.Println("Row", i, "Len:", len(v[0]), v)
	}
}
