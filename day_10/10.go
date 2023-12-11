package day_10

import (
	. "aoc-2023/helpers"
)

type Shifts map[uint8][2]Vec

func Solve(filepath string) {
	shifts := Shifts{
		'|': {Vec{X: -1}, Vec{X: 1}},
		'-': {Vec{Y: -1}, Vec{Y: 1}},
		'L': {Vec{X: -1}, Vec{Y: 1}},
		'J': {Vec{X: -1}, Vec{Y: -1}},
		'7': {Vec{X: 1}, Vec{Y: -1}},
		'F': {Vec{X: 1}, Vec{Y: 1}},
	}
	input := ReadLines(filepath)
	s := Ints(input[0])
	start := Vec{X: s[0], Y: s[1]}
	field := input[1:]

	loop := NewSet(start)
	queue := []Vec{start}
	for i := 0; i < len(queue); i++ {
		from := queue[i]
		for _, shift := range shifts[field[from.X][from.Y]] {
			to := from.Add(shift)
			if !loop.Contains(to) {
				loop.Add(to)
				queue = append(queue, to)
			}
		}
	}
	println((len(queue) + 1) / 2)

	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field[0]); y++ {
			from := Vec{X: x, Y: y}
			if !loop.Contains(from) {
				field[x] = ReplaceStringAt(field[x], y, ".")
			}
		}
	}

	count := 0
	for _, line := range field {
		isInterior := 0
		line = RegexReplace(`(F-*7|L-*J)`, line, "")
		line = RegexReplace(`F-*J|L-*7`, line, "|")

		for _, c := range line {
			if c == '|' {
				isInterior ^= 1
			}
			if c == '.' && isInterior == 1 {
				count += 1
			}
		}
	}
	println(count)
}
