package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var entires []int
	var count int
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line == "" {
			entires = append(entires, count)
			count = 0
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		count += num
	}

	sort.Ints(entires)

	var top int
	for _, num := range entires[len(entires)-3 : len(entires)] {
		top += num
	}

	fmt.Printf("Num: %d\n", top)
}
