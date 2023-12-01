package helpers

import (
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(bytes), "\n")
}

func Sum[T int | int64 | float64](xs []T) T {
	var s T
	for _, x := range xs {
		s += x
	}
	return s
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
