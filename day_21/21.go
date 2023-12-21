package day_21

import (
	. "aoc-2023/helpers"
	"strings"
)

func Solve(filepath string) {
	field := ReadLines(filepath)
	var start Vec
	for x, line := range field {
		y := strings.Index(line, "S")
		if y >= 0 {
			start = Vec{X: x, Y: y}
			break
		}
	}

	sz := len(field)
	points := NewSet[Vec](start)
	numIters := 500
	println(sz)
	for iter := 0; iter < numIters; iter++ {
		if iter == 64 || iter%sz == (sz/2) {
			// use to extrapolate https://www.wolframalpha.com/input?i=+3916%2C+34870%2C+96644%2C+189238%2C+312652
			println(iter, points.Count())
		}
		newPoints := NewSet[Vec]()
		for _, from := range points.Items() {
			for next := range Neighbors4(from) {
				if field[Mod(next.X, sz)][Mod(next.Y, sz)] == '#' {
					continue
				}
				newPoints.Add(next)
			}
		}
		points = newPoints
	}
}
