package day_13

import (
	. "aoc-2023/helpers"
)

func reflectionDiff(mt []string, col int) int {
	rowLen := len(mt[0])
	l := min(col, rowLen-col)
	diff := 0
	for i := 0; i < len(mt); i++ {
		row := mt[i]
		left := Reverse(row[col-l : col])
		right := row[col : col+l]
		for j := 0; j < len(left); j++ {
			if left[j] != right[j] {
				diff += 1
			}
		}
	}
	return diff
}

func findReflection(mt, mt2 []string, targetDiff int) int {
	m := [2][]string{mt, mt2}
	multiplier := [2]int{1, 100}
	for dir := 0; dir < 2; dir++ {
		for col := 1; col < len(m[dir][0]); col++ {
			if reflectionDiff(m[dir], col) == targetDiff {
				return multiplier[dir] * col
			}
		}
	}
	panic("")
}

func Solve(filepath string) {
	s1, s2 := 0, 0
	for _, mt := range ReadBlocks(filepath) {
		mt2 := TransposeStrings(mt)
		s1 += findReflection(mt, mt2, 0)
		s2 += findReflection(mt, mt2, 1)
	}
	println(s1)
	println(s2)
}
