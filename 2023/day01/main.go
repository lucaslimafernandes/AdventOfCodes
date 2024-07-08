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

func main() {

	fmt.Println("2023 - Day 01 #lucaslimafernandes")

	items, _ := readFile()

	for _, v := range items {

		n1, n2, err := pega(v)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%v %v\n", n1, n2)
		fmt.Println(n1 + n2)
	}

}

func readFile() ([]string, error) {

	var res []string

	r, err := os.Open("2023/day01/input.txt")
	if err != nil {
		log.Println(err)
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		// DEBUG
		if i < 2 {
			fmt.Println(scanner.Text())
			res = append(res, scanner.Text())
		}

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

func pega(s string) (int, int, error) {

	var n1 int
	var n2 int

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

	return n1, n2, nil

}
