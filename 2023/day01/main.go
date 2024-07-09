package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const INPUT = "https://adventofcode.com/2023/day/1/input"

// Soma: 55172
func main() {

	fmt.Println("2023 - Day 01 #lucaslimafernandes")

	var values []int
	items, _ := readFile()

	for _, v := range items {

		n, err := pega(v)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%v\n", n)

		values = append(values, n)
	}

	fmt.Printf("Soma: %v ", soma(values))
}

func readFile() ([]string, error) {

	var res []string

	r, err := os.Open("2023/day01/input.txt")
	// r, err := os.Open("2023/day01/test.txt")
	if err != nil {
		log.Println(err)
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		// DEBUG
		// if i < 2 {
		fmt.Println(scanner.Text())
		res = append(res, scanner.Text())
		// }

	}

	return res, nil

}

func reverse(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]

	}

	return string(runes)

}

func pega(s string) (int, error) {

	var n1 int
	var n2 int
	var n int

	reversed := reverse(s)

	for _, char := range s {
		if unicode.IsDigit(char) {
			n1, _ = strconv.Atoi(string(char))
			break
		}
	}

	for _, char2 := range reversed {
		if unicode.IsDigit(char2) {
			n2, _ = strconv.Atoi(string(char2))
			break
		}
	}

	n, _ = strconv.Atoi(fmt.Sprintf("%v%v", n1, n2))

	return n, nil

}

func soma(s []int) int {

	var sum int

	for _, v := range s {
		sum = sum + v
	}

	return sum
}
