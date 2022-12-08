package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(input), "\n")
	for _, str := range strs {
		if str == "" {
			continue
		}

		var startIndex int
		m := map[string]int{}
		for i, c := range str {
			lastIndex, ok := m[string(c)]
			if ok && lastIndex+1 > startIndex {
				startIndex = lastIndex + 1
			}

			m[string(c)] = i

			if i-startIndex == 3 {
				fmt.Printf("Pos: %v\n", i+1)
				break
			}
		}
	}
}
