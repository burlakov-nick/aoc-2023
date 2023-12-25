package day_25

import (
	. "aoc-2023/helpers"
	"strings"
	"sync"
)

func Solve(filepath string) {
	ids := map[string]int{}
	n := 0
	lines := ReadLines(filepath)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		for _, t := range tokens {
			if _, ok := ids[t]; !ok {
				ids[t] = n
				n += 1
			}
		}
	}

	next := make([][]int, n)
	edges := make([]Edge, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		from := ids[tokens[0]]
		for _, name := range tokens[1:] {
			to := ids[name]
			edge := Edge{from, to}
			edges = append(edges, edge)
			next[from] = append(next[from], to)
			next[to] = append(next[to], from)
		}
	}

	var wg sync.WaitGroup
	jobQueue := make(chan []Edge)
	workers := 16
	for i := 0; i < workers; i++ {
		go func() {
			for banned := range jobQueue {
				tryFindBridge(n, next, banned)
				wg.Done()
			}
		}()
	}

	wg.Add(1)
	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			wg.Add(1)
			jobQueue <- []Edge{edges[i], edges[j]}
		}
	}
	wg.Done()
}

type Edge struct {
	from, to int
}

func isBanned(v, to int, banned []Edge) bool {
	for _, ban := range banned {
		if ban.from == v && ban.to == to || ban.from == to && ban.to == v {
			return true
		}
	}
	return false
}

func tryFindBridge(n int, next [][]int, banned []Edge) {
	visited := make([]bool, n)
	time := make([]int, n)
	minUp := make([]int, n)
	s := State{
		timer:   0,
		time:    time,
		minUp:   minUp,
		visited: visited,
		next:    next,
		banned:  banned,
	}
	ok, from, to := findBridge(0, -1, &s)
	if ok {
		bannedNew := []Edge{banned[0], banned[1], {from, to}}
		res := 1
		used := make([]bool, n)
		for i := 0; i < n; i++ {
			if !used[i] {
				res *= count(i, used, next, bannedNew)
			}
		}
		println(res)
		panic("bridge found")
	}

}

func count(v int, visited []bool, next [][]int, banned []Edge) int {
	res := 1
	visited[v] = true
	for _, to := range next[v] {
		if !visited[to] && !isBanned(v, to, banned) {
			res += count(to, visited, next, banned)
		}
	}
	return res
}

type State struct {
	timer       int
	time, minUp []int
	visited     []bool
	next        [][]int
	banned      []Edge
}

func findBridge(v, prev int, s *State) (bool, int, int) {
	s.visited[v] = true
	s.time[v] = s.timer
	s.minUp[v] = s.timer
	s.timer += 1
	for _, to := range s.next[v] {
		if to == prev || isBanned(v, to, s.banned) {
			continue
		}
		if s.visited[to] {
			s.minUp[v] = min(s.minUp[v], s.time[to])
		} else {
			ok, r1, r2 := findBridge(to, v, s)
			if ok {
				return ok, r1, r2
			}
			s.minUp[v] = min(s.minUp[v], s.minUp[to])
			if s.minUp[to] > s.time[v] {
				return true, v, to
			}
		}
	}
	return false, -1, -1
}
