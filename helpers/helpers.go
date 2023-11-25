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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
