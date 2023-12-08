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

	nodeLines := lines[2 : len(lines)-1]
	graph := map[string][]string{}
	re := regexp.MustCompile(`([A-Z]{3})`)
	for _, l := range nodeLines {
		matches := re.FindAllString(l, -1)
		n := matches[0]
		left := matches[1]
		right := matches[2]
		graph[n] = []string{left, right}
	}

	next := "AAA"
	var count int
	for next != "ZZZ" {
		direction := instructions[count%len(instructions)]
		// fmt.Printf("current: %v, dir: %s\n", next, string(direction))
		nodes := graph[next]
		if direction == 'L' {
			next = nodes[0]
		} else {
			next = nodes[1]
		}
		count++
	}

	fmt.Printf("count: %d\n", count)

}
