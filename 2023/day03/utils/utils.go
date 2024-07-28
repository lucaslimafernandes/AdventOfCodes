package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fp string) [][]string {

	var res [][]string

	r, err := os.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// temp := []string{scanner.Text()}
		res = append(res, []string{scanner.Text()})
	}

	return res
}
