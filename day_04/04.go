package day_04

import (
	. "aoc-2023/helpers"
	"strings"
)

func parseCards(line string) (Set[int], Set[int]) {
	parts := strings.Split(line, ":")
	sets := strings.Split(parts[1], "|")
	return ToSet(ParseInts(sets[0], " ")), ToSet(ParseInts(sets[1], " "))
}

func Solve1(filepath string) {
	sum := 0
	for _, line := range ReadLines(filepath) {
		left, right := parseCards(line)
		win := left.Intersect(right).Count()
		if win > 0 {
			sum += 1 << (win - 1)
		}
	}
	println(sum)
}

func Solve2(filepath string) {
	lines := ReadLines(filepath)
	counts := make([]int64, len(lines))
	for i := 0; i < len(counts); i++ {
		left, right := parseCards(lines[i])
		win := left.Intersect(right).Count()
		counts[i] += 1
		for j := i + 1; j < min(len(counts), i+win+1); j++ {
			counts[j] += counts[i]
		}
	}
	println(Sum(counts))
}
