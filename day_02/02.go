package day_02

import (
	. "aoc-2023/helpers"
	"regexp"
	"strconv"
	"strings"
)

type Colors struct {
	red, green, blue int
}

func parseColors(round string) Colors {
	parse := func(color string) int {
		re := regexp.MustCompile(`(\d+) ` + color)
		matched := re.FindStringSubmatch(round)
		if len(matched) == 0 {
			return 0
		}
		parsed, _ := strconv.Atoi(matched[1])
		return parsed
	}

	return Colors{parse("red"), parse("green"), parse("blue")}
}

func Solve1(filepath string) {
	mx := Colors{12, 13, 14}
	sum := 0

	for i, line := range ReadLines(filepath) {
		isPossible := true
		rounds := strings.Split(line, ";")
		for _, round := range rounds {
			colors := parseColors(round)
			if colors.red > mx.red || colors.green > mx.green || colors.blue > mx.blue {
				isPossible = false
			}
		}
		if isPossible {
			sum += i + 1
		}
	}

	println(sum)
}

func Solve2(filepath string) {
	sum := 0
	for _, line := range ReadLines(filepath) {
		mx := Colors{0, 0, 0}
		rounds := strings.Split(line, ";")
		for _, round := range rounds {
			colors := parseColors(round)
			mx = Colors{max(mx.red, colors.red), max(mx.green, colors.green), max(mx.blue, colors.blue)}
		}
		sum += mx.red * mx.green * mx.blue
	}
	println(sum)
}
