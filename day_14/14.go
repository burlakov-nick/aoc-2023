package day_14

import (
	. "aoc-2023/helpers"
	"fmt"
	"strings"
)

type Field struct {
	field [][]byte
	n, m  int
}

func (f Field) inside(x, y int) bool {
	return 0 <= x && x < f.n && 0 <= y && y < f.m
}

func (f Field) move(sx, sy, dx, dy int) {
	for f.inside(sx+dx, sy+dy) && f.field[sx+dx][sy+dy] == '.' {
		f.field[sx+dx][sy+dy] = f.field[sx][sy]
		f.field[sx][sy] = '.'
		sx += dx
		sy += dy
	}
}

func (f Field) tilt(lx, rx, ly, ry, dx, dy int) {
	for x := lx; x != rx; x += Sign(rx - lx) {
		for y := ly; y != ry; y += Sign(ry - ly) {
			if f.field[x][y] == 'O' {
				f.move(x, y, dx, dy)
			}
		}
	}
}

func (f Field) tiltNorth() {
	for x := 0; x < f.n; x++ {
		for y := 0; y < f.m; y++ {
			if f.field[x][y] == 'O' {
				f.move(x, y, -1, 0)
			}
		}
	}
}

func (f Field) tiltWest() {
	for y := 0; y < f.m; y++ {
		for x := 0; x < f.n; x++ {
			if f.field[x][y] == 'O' {
				f.move(x, y, 0, -1)
			}
		}
	}
}

func (f Field) tiltSouth() {
	for x := f.n - 1; x >= 0; x-- {
		for y := 0; y < f.m; y++ {
			if f.field[x][y] == 'O' {
				f.move(x, y, 1, 0)
			}
		}
	}
}

func (f Field) tiltEast() {
	for y := f.m - 1; y >= 0; y-- {
		for x := 0; x < f.n; x++ {
			if f.field[x][y] == 'O' {
				f.move(x, y, 0, 1)
			}
		}
	}
}

func (f Field) totalLoad() int {
	s := 0
	for i := 0; i < f.n; i++ {
		for j := 0; j < f.m; j++ {
			if f.field[i][j] == 'O' {
				s += f.n - i
			}
		}
	}
	return s
}

func (f Field) toString() string {
	xs := make([]string, f.n)
	for i := 0; i < f.n; i++ {
		xs[i] = string(f.field[i])
	}
	return strings.Join(xs, "\n") + "\n"
}

func parse(filepath string) Field {
	lines := ReadLines(filepath)
	field := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		field[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			field[i][j] = lines[i][j]
		}
	}
	return Field{field: field, n: len(field), m: len(field[0])}
}

func Solve1(filepath string) {
	field := parse(filepath)
	field.tiltNorth()
	println(field.totalLoad())
}

func Solve2(filepath string) {
	field := parse(filepath)
	visited := map[string]int{}
	visited[field.toString()] = 0
	for i := 1; i <= 1000000000; i++ {
		field.tiltNorth()
		field.tiltWest()
		field.tiltSouth()
		field.tiltEast()
		key := field.toString()
		firstVisit, ok := visited[key]
		if ok {
			fmt.Printf("First visit %d\n", firstVisit)
			fmt.Printf("After iteration %d\n", i)
			cycle := i - firstVisit
			left := (1000000000 - i) % cycle
			for j := 0; j < left; j++ {
				field.tiltNorth()
				field.tiltWest()
				field.tiltSouth()
				field.tiltEast()
			}
			break
		}
		visited[key] = i
	}
	println(field.totalLoad())
}
