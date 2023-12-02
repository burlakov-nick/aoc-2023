package main

import (
	"aoc-2023/day_01"
	"aoc-2023/day_02"
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
	}
	currentDay := os.Args[1]

	solve := days[currentDay].Solve1
	part := "1"
	if slices.Contains(os.Args, "--hard") {
		solve = days[currentDay].Solve2
		part = "2"
	}

	if !slices.Contains(os.Args, "--only-input") {
		fmt.Printf("/------------ SAMPLE (%s) ------------/\n", part)
		solve("./day_" + currentDay + "/sample.txt")
	}

	if !slices.Contains(os.Args, "--only-sample") {
		fmt.Printf("/------------ INPUT  (%s) ------------/\n", part)
		solve("./day_" + currentDay + "/input.txt")
	}
}
