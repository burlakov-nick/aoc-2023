package day_03

import (
	. "aoc-2023/helpers"
	"regexp"
	"unicode"
)

func Solve1(filepath string) {
	// use solve2
}

func Solve2(filepath string) {
	re := regexp.MustCompile(`(\d)+`)
	field := ReadLines(filepath)
	sz := Vec{len(field), len(field[0])}

	sum := 0
	gears := make(map[Vec][]int)
	for x, line := range field {
		matched := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matched {
			l, r := match[0], match[1]
			number := Int(line[l:r])
			hasSymbolAround := false
			for y := l; y < r; y++ {
				for n := range Neighbors8Boxed(Vec{x, y}, sz) {
					ch := rune(field[n.X][n.Y])
					if ch != '.' && !unicode.IsDigit(ch) {
						hasSymbolAround = true
					}
					if ch == '*' {
						gears[n] = append(gears[n], number)
					}
				}
			}
			if hasSymbolAround {
				sum += number
			}
		}
	}
	println(sum)

	sumGears := 0
	for _, v := range gears {
		v := Distinct(v)
		if len(v) == 2 {
			sumGears += v[0] * v[1]
		}
	}
	println(sumGears)
}
