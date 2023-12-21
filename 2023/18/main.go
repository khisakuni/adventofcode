package main

import (
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
)

const (
	directionDown = iota
	directionUp
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var coords [][]int64
	var x, y int64
	// m[0] = []int64{col}
	var lineLen int64
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		instruction := parts[2]
		l, _ := strconv.ParseInt(instruction[2:7], 16, 64)
		d := instruction[7]

		// fmt.Printf(">> %v, %v\n", d, l)

		lineLen += l
		switch d {
		// Right
		case '0':
			// case "R":
			x += l
		// Down
		case '1':
			// case "D":
			coords = append(coords, []int64{x, y}, []int64{x, y + l})
			y += l
		// Left
		case '2':
			// case "L":
			x -= l
		// Up
		case '3':
			// case "U":
			coords = append(coords, []int64{x, y}, []int64{x, y - l})
			y -= l
		}
	}

	var left, right int64
	for i := 0; i < len(coords)-1; i++ {
		current := coords[i]
		next := coords[i+1]

		left += current[0] * next[1]
		right += current[1] * next[0]
	}

	area := (left - right) / 2

	fmt.Printf("count: %v\n", area+(lineLen/2)+1)
}
