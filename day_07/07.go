package day_07

import (
	. "aoc-2023/helpers"
	"slices"
	"sort"
	"strings"
)

type Hand struct {
	frequencies []int
	cards       string
	value       int
}

func getFrequencies(cards string, useJokers bool) []int {
	jokers := 0
	if useJokers {
		jokers = strings.Count(cards, "J")
		cards = strings.ReplaceAll(cards, "J", "")
	}

	frequency := make(map[rune]int)
	for _, c := range cards {
		frequency[c] = frequency[c] + 1
	}
	var res []int
	for _, v := range frequency {
		res = append(res, v)
	}
	slices.Sort(res)
	slices.Reverse(res)

	if len(res) > 0 {
		res[0] += jokers
	} else {
		res = append(res, 5)
	}
	return res
}

func less(left, right Hand, cardsOrder string) bool {
	for k := 0; k < min(len(left.frequencies), len(right.frequencies)); k++ {
		if left.frequencies[k] != right.frequencies[k] {
			return left.frequencies[k] < right.frequencies[k]
		}
	}
	for k := 0; k < 5; k++ {
		if left.cards[k] != right.cards[k] {
			left := strings.IndexRune(cardsOrder, rune(left.cards[k]))
			right := strings.IndexRune(cardsOrder, rune(right.cards[k]))
			return left > right
		}
	}
	return false
}

func parseHands(filepath string, useJokers bool) []Hand {
	var hands []Hand
	for _, line := range ReadLines(filepath) {
		tokens := strings.Split(line, " ")
		cards, value := tokens[0], Int(tokens[1])
		frequencies := getFrequencies(cards, useJokers)
		hands = append(hands, Hand{frequencies, cards, value})
	}
	return hands
}

func solve(filepath string, useJoker bool, cardsOrder string) {
	hands := parseHands(filepath, useJoker)

	sort.Slice(hands, func(i, j int) bool {
		return less(hands[i], hands[j], cardsOrder)
	})

	s := 0
	for i, hand := range hands {
		s += (i + 1) * hand.value
	}
	println(s)
}

func Solve1(filepath string) {
	solve(filepath, false, "AKQJT98765432")
}

func Solve2(filepath string) {
	solve(filepath, true, "AKQT98765432J")
}
