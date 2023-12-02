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

func getMaxColorsNeeded(line string) Colors {
	rounds := strings.Split(line, ";")
	mx := Colors{0, 0, 0}
	for _, round := range rounds {
		colors := parseColors(round)
		mx = Colors{max(mx.red, colors.red), max(mx.green, colors.green), max(mx.blue, colors.blue)}
	}
	return mx
}

func Solve1(filepath string) {
	available := Colors{12, 13, 14}
	sum := 0

	for i, line := range ReadLines(filepath) {
		mx := getMaxColorsNeeded(line)
		if available.red >= mx.red && available.green >= mx.green && available.blue >= mx.blue {
			sum += i + 1
		}
	}
	println(sum)
}

func Solve2(filepath string) {
	sum := 0
	for _, line := range ReadLines(filepath) {
		mx := getMaxColorsNeeded(line)
		sum += mx.red * mx.green * mx.blue
	}
	println(sum)
}
