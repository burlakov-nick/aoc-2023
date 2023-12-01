package day_01

import (
	. "aoc-2023/helpers"
	"fmt"
	"strings"
)

func solve(filepath string, words map[string]int) {
	for i := 0; i < 10; i++ {
		words[string(rune('0'+i))] = i
	}
	s := 0
	for _, line := range ReadLines(filepath) {
		var digits []int
		for i := 0; i < len(line); i++ {
			for w, val := range words {
				if strings.HasPrefix(line[i:], w) {
					digits = append(digits, val)
					break
				}
			}
		}
		s += digits[0]*10 + digits[len(digits)-1]
	}
	fmt.Println(s)
}

func Solve1(filepath string) {
	solve(filepath, map[string]int{})
}

func Solve2(filepath string) {
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	solve(filepath, words)
}
