package day_24

import (
	. "aoc-2023/helpers"
	"strings"
)

type Point struct {
	x, y, z int
}

const eps = 1e-9

func eq(a, b float64) bool {
	return a-b < eps
}

func inside(left, right, x float64) bool {
	return left <= x && x <= right
}

func line2d(p1, v1 Point) (float64, float64) {
	a := float64(v1.y) / float64(v1.x)
	b := float64(p1.y) - a*float64(p1.x)
	return a, b
}

func intersect2d(p1, v1, p2, v2 Point) (bool, float64, float64) {
	a1, b1 := line2d(p1, v1)
	a2, b2 := line2d(p2, v2)
	if eq(a1, a2) && eq(b1, b2) {
		return false, 0, 0
	}
	cx := (b2 - b1) / (a1 - a2)
	cy := cx*a1 + b1
	inFuture := (cx > float64(p1.x)) == (v1.x > 0) && (cx > float64(p2.x)) == (v2.x > 0)
	return inFuture, cx, cy
}

func parse(filepath string) ([]Point, []Point) {
	point := make([]Point, 0)
	velocity := make([]Point, 0)
	for _, line := range ReadLines(filepath) {
		x := ParseInts(line, " ", ",", "@")
		point = append(point, Point{x[0], x[1], x[2]})
		velocity = append(velocity, Point{x[3], x[4], x[5]})
	}
	return point, velocity
}

func Solve1(filepath string) {
	point, velocity := parse(filepath)

	var left, right float64
	if strings.HasSuffix(filepath, "input.txt") {
		left = 200000000000000.0
		right = 400000000000000.0
	} else {
		left = 7.0
		right = 27.0
	}

	res := 0
	for i := 0; i < len(point); i++ {
		for j := i + 1; j < len(point); j++ {
			ok, cx, cy := intersect2d(point[i], velocity[i], point[j], velocity[j])
			if ok && inside(left, right, cx) && inside(left, right, cy) {
				res += 1
			}
		}
	}
	println(res)
}

func Solve2(filepath string) {
	//
}
