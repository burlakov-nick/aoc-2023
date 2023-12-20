package day_20

import (
	. "aoc-2023/helpers"
	"strings"
)

type Pulse int

const (
	No   Pulse = 0
	Low  Pulse = 1
	High Pulse = 2
)

const broadcaster = "broadcaster"

type Connections []string

type ModulePulse struct {
	Module string
	Pulse  Pulse
}

type FlipFlop struct {
	IsOn bool
}

type Conjunction struct {
	LastPulse map[string]Pulse
}

type Broadcaster struct {
	Pulse Pulse
}

type Module interface {
	Receive(from string, pulse Pulse) Pulse
}

func (m *FlipFlop) Receive(_ string, pulse Pulse) Pulse {
	if pulse != Low {
		return No
	}
	m.IsOn = !m.IsOn
	if m.IsOn {
		return High
	} else {
		return Low
	}
}

func (m *Conjunction) Receive(from string, pulse Pulse) Pulse {
	m.LastPulse[from] = pulse
	if All(Values(m.LastPulse), func(p Pulse) bool {
		return p == High
	}) {
		return Low
	} else {
		return High
	}
}

func (m *Broadcaster) Receive(_ string, _ Pulse) Pulse {
	return Low
}

func parse(filepath string) (map[string]Module, map[string]Connections, map[string]Connections) {
	modules := map[string]Module{}
	prev := map[string]Connections{}
	next := map[string]Connections{}

	lines := ReadLines(filepath)
	conjunctions := make([]string, 0)
	for _, line := range lines {
		tokens := strings.Split(Remove(line, " ->", ","), " ")
		from := tokens[0]
		if from[0] == '%' {
			from = from[1:]
			modules[from] = &FlipFlop{IsOn: false}
		} else if from[0] == '&' {
			from = from[1:]
			conjunctions = append(conjunctions, from)
		} else if from == broadcaster {
			modules[from] = &Broadcaster{}
		}
		next[from] = tokens[1:]
		for _, n := range tokens[1:] {
			prev[n] = append(prev[n], from)
		}
	}

	for _, conj := range conjunctions {
		last := map[string]Pulse{}
		for _, p := range prev[conj] {
			last[p] = Low
		}
		modules[conj] = &Conjunction{last}
	}

	return modules, prev, next
}

func run(modules map[string]Module, next map[string]Connections, onPulse func(mp ModulePulse)) {
	queue := make([]ModulePulse, 0)
	queue = append(queue, ModulePulse{Module: broadcaster, Pulse: Low})
	for i := 0; i < len(queue); i++ {
		from, pulse := queue[i].Module, queue[i].Pulse
		for _, to := range next[from] {
			onPulse(ModulePulse{Pulse: pulse, Module: to})
			newModule, ok := modules[to]
			if ok {
				newPulse := newModule.Receive(from, pulse)
				if newPulse != No {
					queue = append(queue, ModulePulse{Module: to, Pulse: newPulse})
				}
			}
		}
	}
}

func Solve1(filepath string) {
	modules, _, next := parse(filepath)

	counts := map[Pulse]int{Low: 1000}
	for iter := 0; iter < 1000; iter++ {
		run(modules, next, func(mp ModulePulse) {
			pulse, _ := mp.Pulse, mp.Module
			counts[pulse] = counts[pulse] + 1
		})
	}

	println(counts[Low] * counts[High])
}

func Solve2(filepath string) {
	modules, prev, next := parse(filepath)
	predFinal := prev["rx"][0]
	finals := NewSet[string](prev[predFinal]...)

	res := 1
	for iter := 1; finals.Count() > 0; iter++ {
		run(modules, next, func(mp ModulePulse) {
			pulse, to := mp.Pulse, mp.Module
			if pulse == Low && finals.Contains(to) {
				res *= iter
				finals.Remove(to)
			}
		})
	}
	println(res)
}
