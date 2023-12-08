package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	instructions := lines[0]

	var start []string

	nodeLines := lines[2 : len(lines)-1]
	graph := map[string][]string{}
	re := regexp.MustCompile(`([A-Z]{3})`)
	for _, l := range nodeLines {
		matches := re.FindAllString(l, -1)
		n := matches[0]
		if n[2] == 'A' {
			start = append(start, n)
		}
		left := matches[1]
		right := matches[2]
		graph[n] = []string{left, right}
	}

	fmt.Printf("start: %v\n", start)

	var counts []int
	for _, s := range start {
		next := s
		var count int
		for next[2] != 'Z' {
			direction := instructions[count%len(instructions)]
			nodes := graph[next]
			if direction == 'L' {
				next = nodes[0]
			} else {
				next = nodes[1]
			}
			count++
		}

		counts = append(counts, count)
	}

	var steps int
	for i, c := range counts {
		if steps == 0 {
			steps = c
			continue
		}

		steps = lcm(steps, counts[i])
	}

	fmt.Printf("steps %d\n", steps)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a < b {
		tmp := b
		b = a
		a = tmp
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}

	return b
}
