package main

import (
	"fmt"
	"log"

	"github.com/lucaslimafernandes/AdventOfCodes/2023/day01/part1"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day01/part2"
	"github.com/lucaslimafernandes/AdventOfCodes/2023/day01/utils"
)

// Soma: 55172
func main() {

	fmt.Println("2023 - Day 01 #lucaslimafernandes")

	var values []int
	var values2 []int
	items, _ := utils.ReadFile("input.txt")

	// Part 1

	for _, v := range items {

		n, err := part1.Pega(v)
		if err != nil {
			log.Fatalln(err)
		}

		values = append(values, n)
	}

	fmt.Println("Part 1")
	fmt.Printf("Soma: %v \n", utils.Soma(values))

	// Part 2
	test, _ := utils.ReadFile("input.txt")

	for _, v := range test {

		n2, err := part2.Pega2(v)
		if err != nil {
			log.Fatalln(err)
		}

		values2 = append(values2, n2)

	}

	fmt.Println("Part 2")
	fmt.Printf("Soma: %v \n", utils.Soma(values2))

}
