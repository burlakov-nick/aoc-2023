package day_09

import (
	. "aoc-2023/helpers"
	"slices"
)

func extrapolate(xs []int) []int {
	if All(xs, func(x int) bool { return x == 0 }) {
		slices.Insert(xs, 0, 0)
		xs = append(xs, 0)
		return xs
	}

	var diffs []int
	for i := 0; i < len(xs)-1; i++ {
		diffs = append(diffs, xs[i+1]-xs[i])
	}
	diffs = extrapolate(diffs)

	xs = slices.Insert(xs, 0, xs[0]-diffs[0])
	xs = append(xs, diffs[len(diffs)-1]+xs[len(xs)-1])
	return xs
}

func Solve(filepath string) {
	s1, s2 := 0, 0
	for _, line := range ReadLines(filepath) {
		xs := extrapolate(Ints(line))
		s1 += xs[len(xs)-1]
		s2 += xs[0]
	}
	println(s1)
	println(s2)
}
