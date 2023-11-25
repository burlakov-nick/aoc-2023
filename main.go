package main

import (
	"aoc-2023/day_01"
	"fmt"
	"os"
	"slices"
	"time"
)

type Solver struct {
	Solve1 func(string)
	Solve2 func(string)
}

func main() {
	days := map[string]Solver{
		"01": {day_01.Solve1, day_01.Solve2},
	}
	currentDay := fmt.Sprintf("%02d", time.Now().Day())
	if len(os.Args) > 1 {
		currentDay = os.Args[1]
	}

	solve := days[currentDay].Solve1
	part := "1"
	if slices.Contains(os.Args, "--hard") {
		solve = days[currentDay].Solve2
		part = "2"
	}

	fmt.Printf("/------------ SAMPLE (%s) ------------/\n", part)
	solve("./day_" + currentDay + "/sample.txt")

	fmt.Printf("/------------ INPUT  (%s) ------------/\n", part)
	solve("./day_" + currentDay + "/input.txt")
}
