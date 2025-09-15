package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// var input = `47|53
// 97|13
// 97|61
// 97|47
// 75|29
// 61|13
// 75|53
// 29|13
// 97|29
// 53|29
// 61|53
// 97|53
// 61|29
// 47|13
// 75|47
// 97|75
// 47|61
// 75|61
// 47|29
// 75|13
// 53|13

// 75,47,61,53,29
// 97,61,53,29,13
// 75,29,13
// 75,97,47,61,53
// 61,13,29
// 97,13,75,29,47`
//

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

func part1() {
	var count int

	lines := strings.Split(input, "\n")
	rules := map[string]map[string]bool{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		if rule := strings.Split(line, "|"); len(rule) >= 2 {
			r, ok := rules[rule[0]]
			if !ok {
				r = map[string]bool{}
				rules[rule[0]] = r
			}

			r[rule[1]] = true
			continue
		}

		seq := strings.Split(line, ",")
		if len(seq) == 0 {
			continue
		}

		slices.SortFunc(seq, func(a, b string) int {
			ruleA := rules[a]
			if ruleA[b] {
				return -1
			}

			ruleB := rules[b]
			if ruleB[a] {
				return 1
			}

			return 0
		})

		if strings.Join(seq, ",") == line {
			num, _ := strconv.Atoi(seq[len(seq)/2])
			count += num
		}

	}

	fmt.Printf("part 1: %v\n", count)
}

func part2() {
	var count int

	lines := strings.Split(input, "\n")
	rules := map[string]map[string]bool{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		if rule := strings.Split(line, "|"); len(rule) >= 2 {
			r, ok := rules[rule[0]]
			if !ok {
				r = map[string]bool{}
				rules[rule[0]] = r
			}

			r[rule[1]] = true
			continue
		}

		seq := strings.Split(line, ",")
		if len(seq) == 0 {
			continue
		}

		slices.SortFunc(seq, func(a, b string) int {
			ruleA := rules[a]
			if ruleA[b] {
				return -1
			}

			ruleB := rules[b]
			if ruleB[a] {
				return 1
			}

			return 0
		})

		if strings.Join(seq, ",") != line {
			num, _ := strconv.Atoi(seq[len(seq)/2])
			count += num
		}

	}

	fmt.Printf("part 2: %v\n", count)
}
