package day_11

import (
	. "aoc-2023/helpers"
	"strings"
)

func solve(filepath string, expand int) {
	mt := ReadLines(filepath)
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	for i := 0; i < len(mt); i++ {
		if !strings.Contains(mt[i], "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	mt = TransposeStrings(mt)
	for i := 0; i < len(mt); i++ {
		if !strings.Contains(mt[i], "#") {
			emptyCols = append(emptyCols, i)
		}
	}
	mt = TransposeStrings(mt)

	galaxies := make([]Vec, 0)
	for x := 0; x < len(mt); x++ {
		for y := 0; y < len(mt[0]); y++ {
			if mt[x][y] == '#' {
				galaxies = append(galaxies, Vec{x, y})
			}
		}
	}

	s := 0
	for _, left := range galaxies {
		for _, right := range galaxies {
			dist := left.ManhattanDist(right)
			dist += (expand - 1) * Count(emptyRows, func(x int) bool {
				return min(left.X, right.X) <= x && x <= max(left.X, right.X)
			})
			dist += (expand - 1) * Count(emptyCols, func(y int) bool {
				return min(left.Y, right.Y) <= y && y <= max(left.Y, right.Y)
			})
			s += dist
		}
	}
	println(s / 2)
}

func Solve1(filepath string) {
	solve(filepath, 2)
}

func Solve2(filepath string) {
	solve(filepath, 1000000)
}
