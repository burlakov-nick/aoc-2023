package helpers

import (
	"cmp"
	"container/heap"
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

func Reverse(s string) string {
	res := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[len(s)-1-i]
	}
	return string(res)
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

func RemoveAt[T any](xs []T, i int) []T {
	return append(xs[:i], xs[i+1:]...)
}

func Values[K comparable, V any](m map[K]V) []V {
	vs := make([]V, 0)
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
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

func Copy[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
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

func Count[T any](xs []T, f func(T) bool) int {
	c := 0
	for _, x := range xs {
		if f(x) {
			c += 1
		}
	}
	return c
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

func Repeat[T any](value T, n int) []T {
	xs := make([]T, n)
	for i := 0; i < n; i++ {
		xs[i] = value
	}
	return xs
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

func (p Vec) Mul(k int) Vec {
	return Vec{p.X * k, p.Y * k}
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

func (p Vec) ManhattanDist(other Vec) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func Neighbors4(v Vec) []Vec {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	res := make([]Vec, 0)
	for i := 0; i < 4; i++ {
		t := Vec{v.X + dx[i], v.Y + dy[i]}
		res = append(res, t)
	}
	return res
}

func Neighbors4Boxed(v, sz Vec) []Vec {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	res := make([]Vec, 0)
	for i := 0; i < 4; i++ {
		t := Vec{v.X + dx[i], v.Y + dy[i]}
		if t.Inside(sz) {
			res = append(res, t)
		}
	}
	return res
}

func Neighbors8Boxed(v, sz Vec) []Vec {
	res := make([]Vec, 0)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			t := Vec{v.X + dx, v.Y + dy}
			if t.Inside(sz) {
				res = append(res, t)
			}
		}
	}
	return res
}

type Vec3 struct {
	X, Y, Z int
}

func (p Vec3) Add(other Vec3) Vec3 {
	return Vec3{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func HexToInt(s string) int {
	i, err := strconv.ParseInt(s, 16, 0)
	check(err)
	return int(i)
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

func (s Set[T]) Remove(x T) {
	delete(s.items, x)
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

func (s Set[T]) IsSuperSet(other Set[T]) bool {
	for _, x := range other.Items() {
		if !s.Contains(x) {
			return false
		}
	}
	return true
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

func ReplaceStringAt(src string, index int, repl string) string {
	return src[:index] + repl + src[index+1:]
}

func TransposeStrings(src []string) []string {
	result := make([]string, len(src[0]))
	for _, line := range src {
		for y, c := range line {
			result[y] = result[y] + string(c)
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	return x / Abs(x)
}

type HeapPair[T any] struct {
	value   int
	payload T
}
type HeapList[T any] []HeapPair[T]

func (h HeapList[T]) Len() int           { return len(h) }
func (h HeapList[T]) Less(i, j int) bool { return h[i].value < h[j].value }
func (h HeapList[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HeapList[T]) Push(x any) {
	*h = append(*h, x.(HeapPair[T]))
}
func (h *HeapList[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Heap[T any] struct {
	list *HeapList[T]
}

func NewHeap[T any]() Heap[T] {
	list := make(HeapList[T], 0)
	return Heap[T]{list: &list}
}

func (h Heap[T]) Push(payload T, value int) {
	heap.Push(h.list, HeapPair[T]{value, payload})
}

func (h Heap[T]) Pop() (T, int) {
	hp := heap.Pop(h.list).(HeapPair[T])
	return hp.payload, hp.value
}

func (h Heap[T]) Empty() bool {
	return len(*h.list) == 0
}

func Mod(a, b int) int {
	return ((a % b) + b) % b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
