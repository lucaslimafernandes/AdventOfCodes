package part2

import (
	"strconv"
	"strings"
)

var CONFIG = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Game2(s string) int {

	var r, g, b int

	end_nr := strings.Index(s, ":")
	ss := s[end_nr+2:]
	sub := strings.Split(ss, ";")

	for _, v := range sub {

		s2 := strings.Split(v, ",")

		for _, v2 := range s2 {

			v2 = strings.Trim(v2, " ")

			if strings.Contains(v2, "red") {
				i := strings.Index(v2, " ")
				rtemp, _ := strconv.Atoi(v2[0:i])

				if rtemp > r {
					r = rtemp
				}
			}

			if strings.Contains(v2, "green") {
				i := strings.Index(v2, " ")
				gtemp, _ := strconv.Atoi(v2[0:i])

				if gtemp > g {
					g = gtemp
				}
			}

			if strings.Contains(v2, "blue") {
				i := strings.Index(v2, " ")
				btemp, _ := strconv.Atoi(v2[0:i])

				if btemp > b {
					b = btemp
				}
			}

			// fmt.Println(v2)

		}

	}
	// fmt.Println(s, r, g, b)

	return r * g * b

}
