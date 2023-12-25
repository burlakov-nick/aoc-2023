package day_25

import (
	. "aoc-2023/helpers"
	"strings"
	"sync"
)

func Solve(filepath string) {
	next := map[string][]string{}
	edges := make([]Edge, 0)
	nodes := make([]string, 0)
	for _, line := range ReadLines(filepath) {
		tokens := strings.Split(line, " ")
		from := tokens[0]
		nodes = append(nodes, from)
		for _, to := range tokens[1:] {
			edge := Edge{from, to}
			edges = append(edges, edge)
			next[from] = append(next[from], to)
			next[to] = append(next[to], from)
		}
	}
	start := edges[0].from

	var wg sync.WaitGroup
	jobQueue := make(chan []Edge)
	workers := 16
	for i := 0; i < workers; i++ {
		go func() {
			for banned := range jobQueue {
				tryFindBridge(start, nodes, next, banned)
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
	from, to string
}

func isBanned(v, to string, banned []Edge) bool {
	for _, ban := range banned {
		if ban.from == v && ban.to == to || ban.from == to && ban.to == v {
			return true
		}
	}
	return false
}

func tryFindBridge(start string, nodes []string, next map[string][]string, banned []Edge) {
	s := State{
		timer:   0,
		time:    map[string]int{},
		minUp:   map[string]int{},
		visited: NewSet[string](),
		next:    next,
		banned:  banned,
	}
	ok, from, to := findBridge(start, "", &s)
	if ok {
		bannedNew := []Edge{banned[0], banned[1], {from, to}}
		res := 1
		visited := NewSet[string]()
		for _, v := range nodes {
			if !visited.Contains(v) {
				res *= count(v, visited, next, bannedNew)
			}
		}
		println(res)
		panic("bridge found")
	}

}

func count(v string, visited Set[string], next map[string][]string, banned []Edge) int {
	res := 1
	visited.Add(v)
	for _, to := range next[v] {
		if !visited.Contains(to) && !isBanned(v, to, banned) {
			res += count(to, visited, next, banned)
		}
	}
	return res
}

type State struct {
	timer       int
	time, minUp map[string]int
	visited     Set[string]
	next        map[string][]string
	banned      []Edge
}

func findBridge(v, prev string, s *State) (bool, string, string) {
	s.visited.Add(v)
	s.time[v] = s.timer
	s.minUp[v] = s.timer
	s.timer += 1
	for _, to := range s.next[v] {
		if to == prev || isBanned(v, to, s.banned) {
			continue
		}
		if s.visited.Contains(to) {
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
	return false, "", ""
}
