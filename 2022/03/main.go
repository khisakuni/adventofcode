package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(input), "\n")
	var sum int32
	m := map[rune]bool{}
	for i, str := range strs {
		if str == "" {
			continue
		}

		switch i % 3 {
		case 0:
			// Reset
			m = map[rune]bool{}

			// Populate
			for _, c := range str {
				m[c] = false
			}
		case 1:
			// Set val to true if duplicate
			for _, c := range str {
				if _, ok := m[c]; ok {
					m[c] = true
				}
			}
		case 2:
			// If entry exists and value is true, it's in all three.
			for _, c := range str {
				if val, ok := m[c]; ok && val {
					sum += priority(c)
					break
				}
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}

func priority(r rune) int32 {
	reLower := regexp.MustCompile("[a-z]")
	if isLower := reLower.MatchString(string(r)); isLower {
		return r - 'a' + 1
	}

	return r - 'A' + 27
}
