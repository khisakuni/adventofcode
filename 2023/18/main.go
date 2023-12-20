package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const gridSize = 1000

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	grid := make([][]byte, gridSize)
	for i := range grid {
		grid[i] = make([]byte, gridSize)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	var count int
	row := 500
	col := 500
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		d := parts[0]
		l, _ := strconv.Atoi(parts[1])

		switch d {
		case "R":
			for i := 0; i < l; i++ {
				grid[row][col] = '#'
				col++
				count++
			}
		case "D":
			for i := 0; i < l; i++ {
				grid[row][col] = '#'
				row++
				count++
			}
		case "L":
			for i := 0; i < l; i++ {
				grid[row][col] = '#'
				col--
				count++
			}
		case "U":
			for i := 0; i < l; i++ {
				grid[row][col] = '#'
				row--
				count++
			}
		}
	}

	queue := [][]int{
		{501, 501},
	}

	for len(queue) > 0 {
		var current []int
		current, queue = queue[0], queue[1:]
		row, col = current[0], current[1]
		if grid[row][col] == '#' {
			continue
		}
		if row < 0 || row >= len(grid) {
			continue
		}
		if col < 0 || col >= len(grid[0]) {
			continue
		}

		grid[row][col] = '#'
		count++
		queue = append(
			queue,
			[]int{row - 1, col},
			[]int{row + 1, col},
			[]int{row, col + 1},
			[]int{row, col - 1},
		)
	}

	// for _, l := range grid {
	// 	fmt.Printf("%v\n", string(l))
	// }

	fmt.Printf("count: %v\n", count)
}
