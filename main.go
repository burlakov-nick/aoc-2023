package main

import (
	"aoc-2023/day_01"
	"aoc-2023/day_02"
	"aoc-2023/day_03"
	"aoc-2023/day_04"
	"aoc-2023/day_05"
	"aoc-2023/day_06"
	"aoc-2023/day_07"
	"aoc-2023/day_08"
	"aoc-2023/day_09"
	"aoc-2023/day_10"
	"aoc-2023/day_11"
	"aoc-2023/day_12"
	"aoc-2023/day_13"
	"aoc-2023/day_14"
	"fmt"
	"os"
	"slices"
)

type Solver struct {
	Solve1 func(string)
	Solve2 func(string)
}

func main() {
	days := map[string]Solver{
		"01": {day_01.Solve1, day_01.Solve2},
		"02": {day_02.Solve1, day_02.Solve2},
		"03": {day_03.Solve1, day_03.Solve2},
		"04": {day_04.Solve1, day_04.Solve2},
		"05": {day_05.Solve1, day_05.Solve2},
		"06": {day_06.Solve1, day_06.Solve2},
		"07": {day_07.Solve1, day_07.Solve2},
		"08": {day_08.Solve1, day_08.Solve2},
		"09": {day_09.Solve, day_09.Solve},
		"10": {day_10.Solve, day_10.Solve},
		"11": {day_11.Solve1, day_11.Solve2},
		"12": {day_12.Solve1, day_12.Solve2},
		"13": {day_13.Solve, day_13.Solve},
		"14": {day_14.Solve1, day_14.Solve2},
	}
	currentDay := os.Args[1]

	solve := days[currentDay].Solve1
	part := "1"
	if slices.Contains(os.Args, "--hard") {
		solve = days[currentDay].Solve2
		part = "2"
	}

	if !slices.Contains(os.Args, "--input") {
		fmt.Printf("/------------ SAMPLE (%s) ------------/\n", part)
		solve("./day_" + currentDay + "/sample.txt")
	}

	if !slices.Contains(os.Args, "--sample") {
		fmt.Printf("/------------ INPUT  (%s) ------------/\n", part)
		solve("./day_" + currentDay + "/input.txt")
	}
}
