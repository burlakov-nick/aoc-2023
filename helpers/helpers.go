package helpers

import (
	"cmp"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(bytes), "\n")
}

func ReadBlocks(filename string) [][]string {
	lines := ReadLines(filename)
	blocks := [][]string{}
	block := []string{}
	for _, line := range lines {
		if line == "" {
			blocks = append(blocks, block)
			block = []string{}
		} else {
			block = append(block, line)
		}
	}
	if len(block) > 0 {
		blocks = append(blocks, block)
	}
	return blocks
}

func Remove(line string, tokens ...string) string {
	for _, t := range tokens {
		line = strings.ReplaceAll(line, t, "")
	}
	return line
}

func ParseInts(line string, sep string, remove ...string) []int {
	tokens := strings.Split(Remove(line, remove...), sep)
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

func Max[T cmp.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T cmp.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
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

func Map[T1 any, T2 any](items []T1, mp func(T1) T2) []T2 {
	res := []T2{}
	for _, x := range items {
		res = append(res, mp(x))
	}
	return res
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
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

func (p Vec) Inside(sz Vec) bool {
	return 0 <= p.X && p.X < sz.X && 0 <= p.Y && p.Y < sz.Y
}

func (p Vec) Add(other Vec) Vec {
	return Vec{p.X + other.X, p.Y + other.Y}
}

func (p Vec) Equals(other Vec) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Vec) Less(other Vec) bool {
	if p.X != other.X {
		return p.X < other.X
	}
	return p.Y < other.Y
}

func (p Vec) Rotate() Vec {
	return Vec{-p.Y, p.X}
}

func (p Vec) RotateClockwise() Vec {
	return Vec{p.Y, -p.X}
}

func Neighbors4(v Vec) chan Vec {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	ch := make(chan Vec)
	go func() {
		for i := 0; i < 4; i++ {
			t := Vec{v.X + dx[i], v.Y + dy[i]}
			ch <- t
		}
		close(ch)
	}()
	return ch
}

func Neighbors4Boxed(v Vec, sz Vec) chan Vec {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	ch := make(chan Vec)
	go func() {
		for i := 0; i < 4; i++ {
			t := Vec{v.X + dx[i], v.Y + dy[i]}
			if t.Inside(sz) {
				ch <- t
			}
		}
		close(ch)
	}()
	return ch
}

func Neighbors8Boxed(v Vec, sz Vec) chan Vec {
	ch := make(chan Vec)
	go func() {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				t := Vec{v.X + dx, v.Y + dy}
				if t.Inside(sz) {
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

func Ints(s string) []int {
	return ParseInts(s, " ")
}

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable](items ...T) Set[T] {
	xs := Set[T]{}
	xs.items = make(map[T]bool)
	for _, x := range items {
		xs.Add(x)
	}
	return xs
}

func ToSet[T comparable](items []T) Set[T] {
	xs := NewSet[T]()
	for _, x := range items {
		xs.Add(x)
	}
	return xs
}

func (s Set[T]) Add(x T) {
	s.items[x] = true
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
	result := NewSet[T]()
	for _, x := range s.Items() {
		if other.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

func (s Set[T]) Extend(other Set[T]) {
	for _, x := range other.Items() {
		s.Add(x)
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

func RegexReplace(pattern, src, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(src, repl)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
