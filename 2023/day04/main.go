package main

import (
	"fmt"

	"github.com/lucaslimafernandes/AdventOfCodes/2023/day04/part1"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day04/utils"
)

func main() {

	s := utils.ReadFile("test.txt")

	fmt.Println(s)

	for _, v := range s {

		part1.Sep(v)
	}

}
