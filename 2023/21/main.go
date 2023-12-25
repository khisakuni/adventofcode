package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	grid := bytes.Split(input, []byte{'\n'})

	var startRow int
	var startCol int
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				startRow = row
				startCol = col
			}
		}
	}

	// fmt.Printf("%v, %v\n", startRow, startCol)

	const target = 64
	queue := [][]int{{startRow, startCol}, nil}
	var steps int
	visited := map[string]bool{}
	for len(queue) > 0 {
		var current []int
		current, queue = queue[0], queue[1:]
		if current == nil {
			steps++
			if steps == target {
				break
			}
			queue = append(queue, nil)
			visited = map[string]bool{}
			continue
		}

		row, col := current[0], current[1]
		if visited[fmt.Sprintf("%d-%d", row, col)] {
			continue
		}

		visited[fmt.Sprintf("%d-%d", row, col)] = true

		// Up
		if row > 0 && grid[row-1][col] != '#' {
			queue = append(queue, []int{row - 1, col})
		}

		// Right
		if col < len(grid[0])-1 && grid[row][col+1] != '#' {
			queue = append(queue, []int{row, col + 1})
		}

		// Down
		if row < len(grid)-1 && grid[row+1][col] != '#' {
			queue = append(queue, []int{row + 1, col})
		}

		// Left
		if col > 0 && grid[row][col-1] != '#' {
			queue = append(queue, []int{row, col - 1})
		}
	}

	m := map[string]bool{}
	for _, coord := range queue {
		row, col := coord[0], coord[1]
		grid[row][col] = 'O'
		m[fmt.Sprintf("%d-%d", row, col)] = true
	}

	// for _, row := range grid {
	// 	fmt.Printf("%v\n", string(row))
	// }

	fmt.Printf(">> %v\n", len(m))
}
