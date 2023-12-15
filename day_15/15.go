package day_15

import (
	. "aoc-2023/helpers"
	"os"
	"slices"
	"strings"
)

func hash(s string) int {
	hash := 0
	for _, c := range s {
		hash = (hash + int(c)) * 17 % 256
	}
	return hash
}

func Solve1(filepath string) {
	line, _ := os.ReadFile(filepath)
	tokens := strings.Split(string(line), ",")
	s := 0
	for _, token := range tokens {
		s += hash(token)
	}
	println(s)
}

func Solve2(filepath string) {
	line, _ := os.ReadFile(filepath)
	tokens := strings.Split(string(line), ",")
	labels := [256][]string{}
	values := map[string]int{}

	for _, token := range tokens {
		if token[len(token)-1] == '-' {
			label := token[:len(token)-1]
			boxId := hash(label)
			index := slices.Index(labels[boxId], label)
			if index >= 0 {
				labels[boxId] = RemoveAt(labels[boxId], index)
			}
		} else {
			label := token[:len(token)-2]
			boxId := hash(label)
			index := slices.Index(labels[boxId], label)
			if index < 0 {
				labels[boxId] = append(labels[boxId], label)
			}
			values[label] = int(token[len(token)-1] - '0')
		}
	}

	s := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < len(labels[i]); j++ {
			s += (i + 1) * (j + 1) * values[labels[i][j]]
		}
	}
	println(s)
}
