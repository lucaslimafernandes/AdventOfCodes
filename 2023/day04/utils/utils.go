package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fp string) []string {

	var res []string

	f, err := os.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res

}
