package day_05

import (
	. "aoc-2023/helpers"
)

func parseMapping(block []string) [][]int {
	return Map(block[1:], Ints)
}

func convert(id int, mappings [][][]int) int {
	for _, mapping := range mappings {
		for _, desc := range mapping {
			dest, src, length := desc[0], desc[1], desc[2]
			if src <= id && id < src+length {
				id = dest + (id - src)
				break
			}
		}
	}
	return id
}

func Solve1(filepath string) {
	blocks := ReadBlocks(filepath)
	seeds := ParseInts(blocks[0][0], " ", "seeds: ")
	mappings := Map(blocks[1:], parseMapping)

	converted := Map(seeds, func(s int) int {
		return convert(s, mappings)
	})
	mn := Min(converted)
	println(mn)
}

type Segment struct {
	left, right int
}

func cut(xs []Segment, left, right int) []Segment {
	var res []Segment
	for _, x := range xs {
		if left <= x.left && x.right <= right {
		} else if x.right < left || x.left > right {
			res = append(res, x)
		} else if x.left < left && right < x.right {
			res = append(res, Segment{x.left, left - 1})
			res = append(res, Segment{right + 1, x.right})
		} else if left > x.left {
			res = append(res, Segment{x.left, left - 1})
		} else if right < x.right {
			res = append(res, Segment{right + 1, x.right})
		}
	}
	return res
}

func convert2(segments []Segment, mappings [][][]int) []Segment {
	for _, mapping := range mappings {
		var newSegments []Segment
		var rest []Segment
		rest = append(rest, segments...)
		for _, desc := range mapping {
			dest, src, length := desc[0], desc[1], desc[2]
			for _, x := range segments {
				if x.right < src || x.left >= src+length {
				} else {
					oldLeft := max(x.left, src)
					oldRight := min(x.right, src+length-1)
					newLeft := dest + oldLeft - src
					newRight := dest + oldRight - src
					rest = cut(rest, oldLeft, oldRight)
					newSegments = append(newSegments, Segment{newLeft, newRight})
				}
			}
		}
		newSegments = append(newSegments, rest...)
		segments = newSegments
	}
	return segments
}

func Solve2(filepath string) {
	blocks := ReadBlocks(filepath)
	seeds := ParseInts(blocks[0][0], " ", "seeds: ")
	mappings := Map(blocks[1:], parseMapping)

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
