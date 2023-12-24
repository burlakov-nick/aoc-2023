package day_23

import (
	. "aoc-2023/helpers"
	"strings"
)

func isSlope(c byte) bool {
	return strings.ContainsRune("<>v^", rune(c))
}

func isHub(pos Vec, field []string) bool {
	slopes := 0
	for _, n := range Neighbors4(pos) {
		if isSlope(field[n.X][n.Y]) {
			slopes += 1
		}
	}
	return slopes > 1
}

func walkNextHub(steps, steps2 int, prev, pos Vec, field []string, hubs map[Vec]int) Edge {
	hubId, ok := hubs[pos]
	if ok {
		return Edge{hubId, steps}
	}
	for _, next := range Neighbors4(pos) {
		if next == prev || field[next.X][next.Y] == '#' {
			continue
		}
		return walkNextHub(steps+1, steps2+1, pos, next, field, hubs)
	}
	panic("")
}

type Edge struct {
	to, dist int
}

func parse(filepath string, ignoreSlope bool) [][]Edge {
	field := ReadLines(filepath)
	hubs := map[Vec]int{}
	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field); y++ {
			if field[x][y] != '.' {
				continue
			}
			v := Vec{x, y}
			if x == 0 || x == len(field)-1 || isHub(v, field) {
				id := len(hubs)
				hubs[v] = id
			}
		}
	}

	next := make([][]Edge, len(hubs))
	sz := Vec{len(field), len(field[0])}
	dir := map[byte]Vec{
		'<': {0, -1},
		'>': {0, 1},
		'v': {1, 0},
		'^': {-1, 0},
	}
	directions := [4]Vec{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for pos, id := range hubs {
		for _, d := range directions {
			n := pos.Add(d)
			if !n.Inside(sz) {
				continue
			}
			cell := field[n.X][n.Y]
			if cell == '#' {
				continue
			}
			if ignoreSlope || cell == '.' || d == dir[cell] {
				hub := walkNextHub(1, 1, pos, n, field, hubs)
				next[id] = append(next[id], hub)
			}
		}
	}

	return next
}

func dfs(v, fin int, visited []bool, next [][]Edge) int {
	if v == fin {
		return 0
	}
	visited[v] = true
	best := -1
	for _, edge := range next[v] {
		to, dist := edge.to, edge.dist
		if visited[to] {
			continue
		}
		path := dfs(to, fin, visited, next)
		if path >= 0 {
			best = max(best, path+dist)
		}
	}
	visited[v] = false
	return best
}

func Solve1(filepath string) {
	next := parse(filepath, false)
	visited := make([]bool, len(next))
	println(dfs(0, len(next)-1, visited, next))
}

func Solve2(filepath string) {
	next := parse(filepath, true)
	visited := make([]bool, len(next))
	println(dfs(0, len(next)-1, visited, next))
}
