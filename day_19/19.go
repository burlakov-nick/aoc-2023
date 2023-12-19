package day_19

import (
	. "aoc-2023/helpers"
	"strings"
)

type Part map[byte]int

type Rule struct {
	cord, cmp byte
	value     int
	ifMatch   string
}

type Workflow struct {
	workflowId     string
	rules          []Rule
	nextWorkflowId string
}
type Workflows map[string]Workflow

func parse(filepath string) (Workflows, []Part) {
	blocks := ReadBlocks(filepath)
	ws, ps := blocks[0], blocks[1]
	workflows := make(Workflows)
	for _, w := range ws {
		workflowId := w[:strings.Index(w, "{")]
		w = w[strings.Index(w, "{")+1 : len(w)-1]
		rs := strings.Split(w, ",")
		rules := make([]Rule, len(rs)-1)
		for j := 0; j < len(rs)-1; j++ {
			tokens := strings.Split(rs[j], ":")
			rules[j] = Rule{cord: tokens[0][0], cmp: tokens[0][1], value: Int(tokens[0][2:]), ifMatch: tokens[1]}
		}
		workflows[workflowId] = Workflow{workflowId: workflowId, rules: rules, nextWorkflowId: rs[len(rs)-1]}
	}
	parts := make([]Part, len(ps))
	for i, p := range ps {
		xs := strings.Split(p[1:len(p)-1], ",")
		values := make(Part)
		for _, x := range xs {
			values[x[0]] = Int(x[2:])
		}
		parts[i] = values
	}
	return workflows, parts
}

func Solve1(filepath string) {
	workflows, parts := parse(filepath)
	s := 0
	for _, part := range parts {
		if verdict(workflows, part) == "A" {
			s += Sum(Values(part))
		}
	}
	println(s)
}

func (rule Rule) matched(part Part) bool {
	if rule.cmp == '<' {
		return part[rule.cord] < rule.value
	} else {
		return part[rule.cord] > rule.value
	}
}

func verdict(workflows Workflows, part Part) string {
	cur := "in"
	for cur != "A" && cur != "R" {
		next := ""
		for _, rule := range workflows[cur].rules {
			if rule.matched(part) {
				next = rule.ifMatch
				break
			}
		}
		if next == "" {
			next = workflows[cur].nextWorkflowId
		}
		cur = next
	}
	return cur
}

func Solve2(filepath string) {
	workflows, _ := parse(filepath)
	maxRange := Range{1, 4000}
	println(traverse(workflows, "in", PartRanges{'x': maxRange, 'm': maxRange, 'a': maxRange, 's': maxRange}))
}

func traverse(workflows Workflows, cur string, rs PartRanges) int {
	if cur == "R" || rs.combinations() == 0 {
		return 0
	}
	if cur == "A" {
		return rs.combinations()
	}
	s := 0
	for _, rule := range workflows[cur].rules {
		if rule.cmp == '<' {
			s += traverse(workflows, rule.ifMatch, rs.cutRight(rule.cord, rule.value-1))
			rs = rs.cutLeft(rule.cord, rule.value)
		} else {
			s += traverse(workflows, rule.ifMatch, rs.cutLeft(rule.cord, rule.value+1))
			rs = rs.cutRight(rule.cord, rule.value)
		}
	}
	s += traverse(workflows, workflows[cur].nextWorkflowId, rs)
	return s
}

type Range struct {
	left, right int
}
type PartRanges map[byte]Range

func (r Range) cutLeft(until int) Range {
	return Range{max(r.left, until), r.right}
}

func (r Range) cutRight(until int) Range {
	return Range{r.left, min(r.right, until)}
}

func (rs PartRanges) cutLeft(cord byte, until int) PartRanges {
	nrs := rs.copy()
	nrs[cord] = nrs[cord].cutLeft(until)
	return nrs
}

func (rs PartRanges) cutRight(cord byte, until int) PartRanges {
	nrs := rs.copy()
	nrs[cord] = nrs[cord].cutRight(until)
	return nrs
}

func (rs PartRanges) combinations() int {
	c := 1
	for _, r := range rs {
		c *= max(0, r.right-r.left+1)
	}
	return c
}

func (rs PartRanges) copy() PartRanges {
	nrs := make(PartRanges)
	for k, v := range rs {
		nrs[k] = v
	}
	return nrs
}
