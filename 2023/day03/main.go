package main

import (
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day03/part1"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day03/utils"
)

func main() {

	s := utils.ReadFile("test.txt")
	// res1 := part1.SymbolsIdentifier(s)

	// res2 := part1.PossibilitiesIdentifier(res1, s)

	// part1.Game3(res2, s)

	res1 := part1.Numbers(s)

	part1.Checks(res1, s)

}
