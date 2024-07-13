package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(file_path string) ([]string, error) {

	var res []string

	r, err := os.Open(file_path)
	if err != nil {
		log.Println(err)
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		// DEBUG
		// if i < 2 {
		res = append(res, scanner.Text())
		// }

	}

	return res, nil

}

func Reverse(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]

	}

	return string(runes)

}

func Soma(s []int) int {

	var sum int

	for _, v := range s {
		sum = sum + v
	}

	return sum
}
