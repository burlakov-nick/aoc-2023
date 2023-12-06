package day_06

import (
	. "aoc-2023/helpers"
	"unicode"
)

func solve(time, dist int) int {
	var cnt int
	for speed := 1; speed < time; speed++ {
		if speed*(time-speed) > dist {
			cnt += 1
		}
	}
	return cnt
}

func Solve1(filepath string) {
	lines := ReadLines(filepath)
	time := ParseInts(lines[0], " ", "Time:")
	dist := ParseInts(lines[1], " ", "Distance:")
	res := 1
	for i := 0; i < len(time); i++ {
		res *= solve(time[i], dist[i])
	}
	println(res)
}

func parseIntWithKerning(line string) int {
	var x int
	for _, c := range line {
		if unicode.IsDigit(c) {
			x = x*10 + int(c-'0')
		}
	}
	return x
}

func Solve2(filepath string) {
	lines := ReadLines(filepath)
	time := parseIntWithKerning(lines[0])
	dist := parseIntWithKerning(lines[1])
	println(solve(time, dist))
}
