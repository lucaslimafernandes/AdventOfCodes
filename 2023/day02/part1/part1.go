package part1

import (
	"strconv"
	"strings"
)

var CONFIG = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// game checks if a game is possible based on the available quantities of
// red, green and blue cubes
// Parameters:
// - s: a string in the format "Game n: x red, y green, z blue; ..." representig the game's moves.
// Returns:
// - int: the game number.
// - bool: true if the game is possible, false otherwise

func Game(s string) (int, bool) {

	end_nr := strings.Index(s, ":")
	nr_game, _ := strconv.Atoi(s[5:end_nr])
	ss := s[end_nr+2:]
	sub := strings.Split(ss, ";")

	for _, v := range sub {

		s2 := strings.Split(v, ",")

		for _, v2 := range s2 {

			var r, g, b int
			v2 := strings.Trim(v2, " ")

			if strings.Contains(v2, "red") {
				i := strings.Index(v2, " ")
				r, _ = strconv.Atoi(v2[0:i])
			}

			if strings.Contains(v2, "green") {
				i := strings.Index(v2, " ")
				g, _ = strconv.Atoi(v2[0:i])
			}

			if strings.Contains(v2, "blue") {
				i := strings.Index(v2, " ")
				b, _ = strconv.Atoi(v2[0:i])
			}

			if r > CONFIG["red"] || g > CONFIG["green"] || b > CONFIG["blue"] {
				return nr_game, false
			}
		}
	}

	return nr_game, true

}

// sums IDs of the possible games from a result map
// Paramenters:
// -m: a map of integers to booleans where the key is the game ID and the value is true if the game is possible, false otherwise.
// Returns:
// - int: the sum of the IDs of the possible games.
func SumIds(m map[int]bool) int {

	res := 0

	for i := range m {

		if m[i] {
			res = res + i
		}
	}

	return res
}
