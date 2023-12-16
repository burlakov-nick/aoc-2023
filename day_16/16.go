package day_16

import (
	. "aoc-2023/helpers"
)

type State struct {
	field     []string
	n, m      int
	visited   Set[[4]int]
	energized Set[[2]int]
}

func dfs(x, y, dx, dy int, s State) {
	if x < 0 || x >= s.n || y < 0 || y >= s.m {
		return
	}
	cell := [2]int{x, y}
	cellDir := [4]int{x, y, dx, dy}
	if s.visited.Contains(cellDir) {
		return
	}
	s.visited.Add(cellDir)
	s.energized.Add(cell)
	switch s.field[x][y] {
	case '.':
		dfs(x+dx, y+dy, dx, dy, s)
	case '-':
		if dx != 0 {
			dfs(x, y-1, 0, -1, s)
			dfs(x, y+1, 0, 1, s)
		} else {
			dfs(x+dx, y+dy, dx, dy, s)
		}
	case '|':
		if dy != 0 {
			dfs(x-1, y, -1, 0, s)
			dfs(x+1, y, 1, 0, s)
		} else {
			dfs(x+dx, y+dy, dx, dy, s)
		}
	case '\\':
		dx, dy = dy, dx
		dfs(x+dx, y+dy, dx, dy, s)
	case '/':
		dx, dy = -dy, -dx
		dfs(x+dx, y+dy, dx, dy, s)
	}
}

func countEnergized(x, y, dx, dy int, field []string) int {
	state := State{field: field, n: len(field), m: len(field[0]), visited: NewSet[[4]int](), energized: NewSet[[2]int]()}
	dfs(x, y, dx, dy, state)
	return state.energized.Count()
}

func Solve1(filepath string) {
	field := ReadLines(filepath)
	println(countEnergized(0, 0, 0, 1, field))
}

func Solve2(filepath string) {
	field := ReadLines(filepath)
	best := 0
	for i := 0; i < len(field); i++ {
		best = max(best, countEnergized(i, 0, 0, 1, field))
		best = max(best, countEnergized(i, len(field[0])-1, 0, -1, field))
	}
	for i := 0; i < len(field[0]); i++ {
		best = max(best, countEnergized(0, i, 1, 0, field))
		best = max(best, countEnergized(len(field)-1, i, -1, 0, field))
	}
	// 8134 -- too low
	println(best)
}
