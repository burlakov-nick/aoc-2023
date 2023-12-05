package day_05

import (
	. "aoc-2023/helpers"
)

// Segment inclusive
type Segment struct {
	left, right int
}

type MapInfo struct {
	source Segment
	shift  int
}

func (x Segment) IsIntersecting(other Segment) bool {
	return x.left <= other.right && other.left <= x.right
}

func (x Segment) Intersect(other Segment) Segment {
	left := max(x.left, other.left)
	right := min(x.right, other.right)
	return Segment{left, right}
}

func (x Segment) Shift(shift int) Segment {
	return Segment{x.left + shift, x.right + shift}
}

func (x Segment) Inside(p int) bool {
	return x.left <= p && p <= x.right
}

func remove(xs []Segment, c Segment) []Segment {
	var res []Segment
	for _, x := range xs {
		if !x.IsIntersecting(c) {
			res = append(res, x)
			continue
		}
		if x.left < c.left && c.left <= x.right {
			res = append(res, Segment{x.left, c.left - 1})
		}
		if x.left <= c.right && c.right < x.right {
			res = append(res, Segment{c.right + 1, x.right})
		}
	}
	return res
}

func parseInput(filepath string) ([]int, [][]MapInfo) {
	blocks := ReadBlocks(filepath)
	seeds := ParseInts(blocks[0][0], " ", "seeds: ")
	var mappings [][]MapInfo
	for _, block := range blocks[1:] {
		var mapping []MapInfo
		for _, line := range block[1:] {
			ints := Ints(line)
			dest, source, length := ints[0], ints[1], ints[2]
			mapping = append(mapping, MapInfo{Segment{source, source + length - 1}, dest - source})
		}
		mappings = append(mappings, mapping)
	}
	return seeds, mappings
}

func convert(x int, mappings [][]MapInfo) int {
	for _, mapping := range mappings {
		for _, m := range mapping {
			if m.source.Inside(x) {
				x += m.shift
				break
			}
		}
	}
	return x
}

func Solve1(filepath string) {
	seeds, mappings := parseInput(filepath)

	converted := Map(seeds, func(s int) int {
		return convert(s, mappings)
	})
	mn := Min(converted)
	println(mn)
}

func convert2(segments []Segment, mappings [][]MapInfo) []Segment {
	for _, mapping := range mappings {
		var newSegments []Segment
		var unmapped []Segment
		unmapped = append(unmapped, segments...)
		for _, m := range mapping {
			for _, x := range segments {
				if x.IsIntersecting(m.source) {
					old := x.Intersect(m.source)
					unmapped = remove(unmapped, old)
					newSegments = append(newSegments, old.Shift(m.shift))
				}
			}
		}
		newSegments = append(newSegments, unmapped...)
		segments = newSegments
	}
	return segments
}

func Solve2(filepath string) {
	seeds, mappings := parseInput(filepath)

	var segments []Segment
	for i := 0; i < len(seeds); i += 2 {
		segments = append(segments, Segment{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	segments = convert2(segments, mappings)

	mn := Min(Map(segments, func(x Segment) int {
		return x.left
	}))
	println(mn)
}