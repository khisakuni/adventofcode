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

	reLower := regexp.MustCompile("[a-z]")

	strs := strings.Split(string(input), "\n")
	var sum int32
	for _, str := range strs {
		if str == "" {
			continue
		}

		first := str[:len(str)/2]
		second := str[len(str)/2:]
		m := map[rune]struct{}{}
		for _, c := range first {
			m[c] = struct{}{}
		}

		for _, c := range second {
			if _, ok := m[c]; ok {
				if isLower := reLower.MatchString(string(c)); isLower {
					sum += c - 'a' + 1
					break
				}

				sum += c - 'A' + 27
				break
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
