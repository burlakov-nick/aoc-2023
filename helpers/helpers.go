package helpers

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(bytes), "\n")
}

func ParseInts(line string, sep string) []int {
	tokens := strings.Split(line, sep)
	xs := []int{}
	for _, token := range tokens {
		x, err := strconv.Atoi(token)
		if err == nil {
			xs = append(xs, x)
		}
	}
	return xs
}

func Sum[T int | int64 | float64](xs []T) T {
	var s T
	for _, x := range xs {
		s += x
	}
	return s
}

func Distinct[T string | int](items []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func Cells[T any](m [][]T) chan T {
	ch := make(chan T)
	go func() {
		for _, row := range m {
			for _, cell := range row {
				ch <- cell
			}
		}
		close(ch)
	}()
	return ch
}

type Vec struct {
	X, Y int
}

func Neighbors8(v Vec, sz Vec) chan Vec {
	ch := make(chan Vec)
	go func() {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				t := Vec{v.X + dx, v.Y + dy}
				if 0 <= t.X && t.X < sz.X && 0 <= t.Y && t.Y < sz.Y {
					ch <- t
				}
			}
		}
		close(ch)
	}()
	return ch
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

type Set[T comparable] struct {
	items map[T]bool
}

func ToSet[T comparable](items []T) Set[T] {
	xs := Set[T]{}
	for _, x := range items {
		xs = xs.Add(x)
	}
	return xs
}

func (s Set[T]) Add(x T) Set[T] {
	if s.items == nil {
		s.items = make(map[T]bool)
	}
	s.items[x] = true
	return s
}

func (s Set[T]) Contains(x T) bool {
	return s.items[x]
}

func (s Set[T]) Count() int {
	return len(s.items)
}

func (s Set[T]) Items() []T {
	keys := make([]T, 0, len(s.items))
	for k := range s.items {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := Set[T]{}
	for _, x := range s.Items() {
		if other.Contains(x) {
			result = result.Add(x)
		}
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
