package day_17

import (
	. "aoc-2023/helpers"
)

type State struct {
	pos, dir Vec
}

func update(state State, value int, q Heap[State], dist map[State]int) {
	oldDist, ok := dist[state]
	if !ok || oldDist > value {
		dist[state] = value
		q.Push(state, value)
	}
}

func solve(mt []string, minSteps, maxSteps int) int {
	init1 := State{Vec{}, Vec{X: 1}}
	init2 := State{Vec{}, Vec{Y: 1}}

	q := NewHeap[State]()
	dist := map[State]int{}
	sz := Vec{len(mt), len(mt[0])}

	update(init1, 0, q, dist)
	update(init2, 0, q, dist)

	for !q.Empty() {
		state, value := q.Pop()
		pos, dir := state.pos, state.dir
		if value > dist[state] {
			continue
		}

		newDist := dist[state]
		for i := 1; i <= maxSteps; i++ {
			newPos := pos.Add(dir.Mul(i))
			if !newPos.Inside(sz) {
				continue
			}
			newDist += int(mt[newPos.X][newPos.Y] - '0')
			if i >= minSteps {
				update(State{newPos, dir.Rotate()}, newDist, q, dist)
				update(State{newPos, dir.RotateClockwise()}, newDist, q, dist)
			}
		}
	}

	finish := Vec{len(mt) - 1, len(mt[0]) - 1}
	return min(dist[State{finish, Vec{0, 1}}], dist[State{finish, Vec{1, 0}}])
}

func Solve1(filepath string) {
	mt := ReadLines(filepath)
	println(solve(mt, 1, 3))
}

func Solve2(filepath string) {
	mt := ReadLines(filepath)
	println(solve(mt, 4, 10))
}
