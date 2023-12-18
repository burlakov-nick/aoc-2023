package day_18

import (
	. "aoc-2023/helpers"
	"strings"
)

func Solve1(filepath string) {
	cur := Vec{}
	directions := map[byte]Vec{
		'R': {0, 1},
		'L': {0, -1},
		'U': {-1, 0},
		'D': {1, 0},
	}

	q, b := 0, 0
	for _, line := range ReadLines(filepath) {
		tokens := strings.Split(line, " ")
		dir, count := tokens[0][0], Int(tokens[1])
		next := cur.Add(directions[dir].Mul(count))
		q += (cur.Y + next.Y) * (cur.X - next.X)
		b += count
		cur = next
	}
	println(Abs(q)/2 + b/2 + 1)
}

func Solve2(filepath string) {
	directions := [4]Vec{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cur := Vec{}
	q, b := 0, 0
	for _, line := range ReadLines(filepath) {
		tokens := strings.Split(line, " ")
		count := HexToInt(tokens[2][2:7])
		dir := int(tokens[2][7] - '0')
		next := cur.Add(directions[dir].Mul(count))
		q += (cur.Y + next.Y) * (cur.X - next.X)
		b += count
		cur = next
	}
	println(Abs(q)/2 + b/2 + 1)
}
