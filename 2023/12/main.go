package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Credit to https://github.com/rumkugel13/AdventOfCode2023/blob/main/day12.go

*/

func main() {
	data, _ := os.ReadFile("input.txt")

	var total int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		var groups []int
		groupStrs := strings.Split(parts[1], ",")
		for _, str := range groupStrs {
			num, _ := strconv.Atoi(str)
			groups = append(groups, num)
		}

		springs := parts[0]
		// res := dp(0, 0, record, groups, cache)
		total += countVariations([]byte(springs), groups, 0)

	}

	fmt.Printf("res: %v\n", total)
}

func countVariations(springs []byte, groups []int, start int) int {
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			// Encountered '?'; generate two generations:
			// One with an opeartion spring and the other with a broken one.
			// A valid configuration will be added to the total.

			springs[i] = '.'
			count := countVariations(springs, groups, i+1)

			springs[i] = '#'
			count += countVariations(springs, groups, i+1)

			springs[i] = '?'
			return count
		}
	}

	// Found valid variation: count towards total.
	if isValidVariation(springs, groups) {
		return 1
	}

	// Not a valid variation: don't count towards total.
	return 0
}

func isValidVariation(springs []byte, groups []int) bool {
	// fmt.Printf("springs: %v\n", string(springs))
	group := 0

	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' {
			// Found broken spring; count the number of broke  springs and compare to current group.
			start := i
			for i < len(springs) && springs[i] == '#' {
				i++
			}

			// Number of broken springs doesn't match the current group.
			// This is an invalid variation.
			if group < len(groups) && i-start != groups[group] {
				return false
			}

			// Number of broken springs matches the number for the current group.
			// Advance to the next group.
			group++
		}
	}

	// The variation is valid if all the groups have been verified.
	return group == len(groups)
}
