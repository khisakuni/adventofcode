package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Credit to https://github.com/macos-fuse-t/aoc/blob/main/2023/12/main.go

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

		var unfoldedSprings []string
		var unfoldedGroups []int
		for i := 0; i < 5; i++ {
			unfoldedSprings = append(unfoldedSprings, springs)
			unfoldedGroups = append(unfoldedGroups, groups...)
		}

		groups = unfoldedGroups

		springs = strings.Join(unfoldedSprings, "?")

		// fmt.Printf("springs: %v\n", springs)
		// fmt.Printf("groups: %v\n", unfoldedGroups)

		cache := map[string]int{}
		// total += countVariations([]byte(springs), groups, 0, cache)
		// fmt.Printf("total: %v\n", total)
		total += countArrangements(springs, groups, cache)

	}

	fmt.Printf("res: %v\n", total)
}

// validate checks if the given variation is valid given the remaining groups of broken springs.
func validate(s string, groups []int) (bool, bool) {
	questions := strings.Count(s, "?")
	brokens := strings.Count(s, "#")

	// We've exhausted the groups but there are still more brokens.
	// We've reached the end but it's invalid.
	if brokens > 0 && len(groups) == 0 {
		return false, true
	}

	// Replace all questions with operational springs.
	s = strings.ReplaceAll(s, "?", ".")

	// We've reached the end if there are no more questions.
	end := questions == 0

	// Get the broken groups.
	brokenGroups := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	})

	// If the number of broken groups doesn't match the group count, it's invalid.
	if len(brokenGroups) != len(groups) {
		return false, end
	}

	// Check that the broken count matches the group.
	for i, c := range groups {
		if c != len(brokenGroups[i]) {
			return false, end
		}
	}

	// All broken groups match group. This sequence is valid.
	return true, true
}

// eat consumes the string for the given amount of brokens and returns the index.
func eat(s string, group int) (int, bool) {
	// There are less springs than there are broken ones for the group.
	// No need to check.
	if len(s) < group {
		return 0, false
	}

	// If there are any operational springs between the beginning and
	// the size of the group, it's invalid.
	for i := 0; i < group; i++ {
		if s[i] == '.' {
			return 0, false
		}
	}

	// Remaining number of springs exactly matches the number of broken ones in the group.
	if len(s) == group {
		return group, true
	}

	// The end of the group is a delimeter.
	if s[group] == '.' || s[group] == '?' {
		return group + 1, true
	}

	// Couldn't consume any springs.
	return 0, false
}

func countArrangements(s string, groups []int, visited map[string]int) int {
	key := s + fmt.Sprintf("%v", groups)
	// fmt.Printf("s: %v\n", s)
	if n, ok := visited[key]; ok {
		return n
	}

	visited[key] = 0
	if valid, end := validate(s, groups); end {
		if valid {
			// Found a valid variation.
			visited[key] = 1
			return 1
		}

		// Found invalid variation: don't count towards total.
		return 0
	}

	if s[0] == '.' {
		n := countArrangements(s[1:], groups, visited)
		visited[key] = n
		return n
	}

	n := 0
	if s[0] == '?' {
		n += countArrangements(s[1:], groups, visited)
	}

	cnt, ok := eat(s, groups[0])
	if !ok {
		// Couldn't consume anymore: we're done here.
		visited[key] = n
		return n
	}

	// Check the permutations for the rest of the string.
	n += countArrangements(s[cnt:], groups[1:], visited)
	visited[key] = n
	return n
}
