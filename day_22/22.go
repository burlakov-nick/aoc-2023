package day_22

import (
	. "aoc-2023/helpers"
	"sort"
	"strings"
)

type Segment struct {
	from, to Vec3
}

func (s Segment) iter() chan Vec3 {
	ch := make(chan Vec3)
	go func() {
		for x := s.from.X; x <= s.to.X; x++ {
			for y := s.from.Y; y <= s.to.Y; y++ {
				for z := s.from.Z; z <= s.to.Z; z++ {
					ch <- Vec3{x, y, z}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func parse(filepath string) []Segment {
	segments := make([]Segment, 0)
	for _, line := range ReadLines(filepath) {
		ints := ParseInts(strings.Replace(line, "~", ",", -1), ",")
		lx := min(ints[0], ints[3])
		rx := max(ints[0], ints[3])
		ly := min(ints[1], ints[4])
		ry := max(ints[1], ints[4])
		lz := min(ints[2], ints[5])
		rz := max(ints[2], ints[5])
		segments = append(segments, Segment{
			Vec3{lx, ly, lz},
			Vec3{rx, ry, rz},
		})
	}
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].from.Z < segments[j].from.Z
	})
	return segments
}

func canMoveDown(segment Segment, busy map[Vec3]int) (bool, []int) {
	if segment.from.Z == 1 {
		return false, nil
	}
	touched := NewSet[int]()
	segment.from.Z -= 1
	segment.to.Z -= 1
	for v := range segment.iter() {
		t, ok := busy[v]
		if ok {
			touched.Add(t)
		}
	}
	return touched.Count() == 0, touched.Items()
}

func moveDown(segments []Segment) ([]Set[int], []Set[int]) {
	busy := map[Vec3]int{}
	up := make([]Set[int], len(segments))
	down := make([]Set[int], len(segments))
	for i := 0; i < len(segments); i++ {
		up[i] = NewSet[int]()
		down[i] = NewSet[int]()
	}

	for i := 0; i < len(segments); i++ {
		for {
			canMove, touched := canMoveDown(segments[i], busy)
			if !canMove {
				for _, t := range touched {
					up[t].Add(i)
					down[i].Add(t)
				}
				break
			}
			segments[i].from.Z -= 1
			segments[i].to.Z -= 1
		}
		for v := range segments[i].iter() {
			busy[v] = i
		}
	}
	return up, down
}

func Solve(filepath string) {
	segments := parse(filepath)
	up, down := moveDown(segments)

	res := 0
	for i := 0; i < len(segments); i++ {
		if All(up[i].Items(), func(u int) bool {
			return down[u].Count() > 1
		}) {
			res += 1
		}
	}
	println(res)

	res2 := 0
	for i := 0; i < len(segments); i++ {
		fallen := NewSet[int](i)
		for j := i + 1; j < len(segments); j++ {
			if segments[j].from.Z == 1 {
				continue
			}
			if fallen.IsSuperSet(down[j]) {
				fallen.Add(j)
			}
		}
		res2 += fallen.Count() - 1
	}
	println(res2)
}
