package day_12

import (
	. "aoc-2023/helpers"
	"strings"
)

func solve(filepath string, unfold bool) {
	s := 0
	for _, line := range ReadLines(filepath) {
		tokens := strings.Split(line, " ")
		if unfold {
			tokens[0] = strings.Join(Repeat(tokens[0], 5), "?")
			tokens[1] = strings.Join(Repeat(tokens[1], 5), ",")
		}

		springs := tokens[0]
		groups := ParseInts(tokens[1], ",")

		dp := make([][]int, len(springs)+1)
		for i := 0; i <= len(springs); i++ {
			dp[i] = make([]int, len(groups)+1)
		}

		dp[0][0] = 1
		for i := 0; i < len(springs); i++ {
			for j := 0; j <= len(groups); j++ {
				if springs[i] == '.' || springs[i] == '?' {
					dp[i+1][j] += dp[i][j]
				}

				if j == len(groups) || i+groups[j] > len(springs) {
					continue
				}
				group := groups[j]
				if strings.Contains(springs[i:i+group], ".") {
					continue
				}

				if i+group == len(springs) {
					dp[i+group][j+1] += dp[i][j]
				} else if springs[i+group] == '.' || springs[i+group] == '?' {
					dp[i+group+1][j+1] += dp[i][j]
				}
			}
		}
		s += dp[len(springs)][len(groups)]
	}
	println(s)
}

func Solve1(filepath string) {
	solve(filepath, false)
}

func Solve2(filepath string) {
	solve(filepath, true)
}
