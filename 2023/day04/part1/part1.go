package part1

import (
	"fmt"
	"strings"
)

func Sep(s string) {

	// var cardNumber int

	cnTemp := s[len("Card "):strings.Index(s, ":")]

	fmt.Println(cnTemp)
}
