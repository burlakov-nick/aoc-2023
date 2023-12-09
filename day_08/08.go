package day_08

import (
	. "aoc-2023/helpers"
	"strings"
)

type Network struct {
	instructions string
	edges        map[string][2]string
}

func parse(filepath string) Network {
	lines := ReadLines(filepath)
	instructions := lines[0]
	edges := map[string][2]string{}
	for _, line := range lines[2:] {
		tokens := strings.Fields(Remove(line, "=", "(", ",", ")"))
		edges[tokens[0]] = [2]string{tokens[1], tokens[2]}
	}
	return Network{instructions, edges}
}

func (n Network) CountSteps(start string, isFinish func(string) bool) int {
	pos, step := start, 0
	for ; !isFinish(pos); step += 1 {
		dir := rune(n.instructions[step%len(n.instructions)])
		if dir == 'L' {
			pos = n.edges[pos][0]
		} else {
			pos = n.edges[pos][1]
		}
	}
	return step
}

func Solve1(filepath string) {
	network := parse(filepath)
	isFinish := func(s string) bool { return s == "ZZZ" }

	println(network.CountSteps("AAA", isFinish))
}

func Solve2(filepath string) {
	network := parse(filepath)
	var start []string
	for pos, _ := range network.edges {
		if strings.HasSuffix(pos, "A") {
			start = append(start, pos)
		}
	}
	isFinish := func(s string) bool { return strings.HasSuffix(s, "Z") }
	res := 1
	for _, pos := range start {
		res = LCM(res, network.CountSteps(pos, isFinish))
	}
	println(res)
}
