package part1

import (
	"fmt"
	"strings"
)

func Sep(s []string) {

	// var cardNumber []int

	// Iter Card numbers
	for _, cn := range s {

		cnTemp := cn[len("Card "):strings.Index(cn, ":")]
		fmt.Println(cnTemp)
	}

}
