package main

import (
	"fmt"

	"github.com/lucaslimafernandes/AdventOfCodes/2023/day02/part1"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day02/part2"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day02/utils"
)

var CONFIG = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// 2023 - Day 02 #lucaslimafernandes
// Soma: 2169
func main() {

	fmt.Println("2023 - Day 02 #lucaslimafernandes")

	// Part 1
	games := make(map[int]bool)

	a := utils.ReadFile("input.txt")

	for _, v := range a {

		i, v := part1.Game(v)
		games[i] = v

	}

	soma := part1.SumIds(games)

	// for k := range games {
	// 	fmt.Println(k, games[k])
	// }

	fmt.Println("Soma:", soma)

	// Part 2
	var res int

	for _, v2 := range a {

		res = res + part2.Game2(v2)

	}

	fmt.Println("Soma part2:", res)

}
