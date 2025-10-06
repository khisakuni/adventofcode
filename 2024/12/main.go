package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// var input = `AAAA
// BBCD
// BBCC
// EEEC`

// var input = `RRRRIICCFF
// RRRRIICCCF
// VVRRRCCFFF
// VVRCCCJFFF
// VVVVCJJCFE
// VVIVCCJJEE
// VVIIICJJEE
// MIIIIIJJEE
// MIIISIJEEE
// MMMISSJEEE`

// var input = `EEEEE
// EXXXX
// EEEEE
// EXXXX
// EEEEE`

// var input = `AAAAAA
// AAABBA
// AAABBA
// ABBAAA
// ABBAAA
// AAAAAA`

func main() {
	part1()
	part2()
	// solve(input)
}

func part1() {
	var grid [][]byte
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(line))
	}

	visited := map[string]bool{}

	var current byte
	var currentArea int
	var currentPerimeter int

	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || y < 0 {
			return false
		}

		if x >= len(grid[0]) || y >= len(grid) {
			return false
		}

		if grid[y][x] != current {
			return false
		}

		key := fmt.Sprintf("%v-%v", y, x)
		if visited[key] {
			return true
		}

		visited[key] = true

		currentArea++
		// Top
		if !dfs(x, y-1) {
			currentPerimeter++
		}
		// Right
		if !dfs(x+1, y) {
			currentPerimeter++
		}
		// Down
		if !dfs(x, y+1) {
			currentPerimeter++
		}
		// Left
		if !dfs(x-1, y) {
			currentPerimeter++
		}

		return true
	}

	var sum int
	for y, line := range grid {
		for x, c := range line {
			key := fmt.Sprintf("%v-%v", y, x)
			if visited[key] {
				continue
			}

			current = c
			currentArea = 0
			currentPerimeter = 0
			dfs(x, y)

			sum += currentArea * currentPerimeter
		}
	}

	fmt.Printf("part 1: %v\n", sum)
}

func part2() {
	var grid [][]byte
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(line))
	}

	visited := map[string]bool{}

	var current byte
	var currentArea int
	var currentSides int
	var horizontal map[string]string
	var vertical map[string]string

	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || y < 0 {
			return false
		}

		if x >= len(grid[0]) || y >= len(grid) {
			return false
		}

		if grid[y][x] != current {
			return false
		}

		key := fmt.Sprintf("%v-%v", y, x)
		if visited[key] {
			return true
		}

		visited[key] = true

		currentArea++
		// Top
		if !dfs(x, y-1) {
			k := fmt.Sprintf("%v,%v", y, x)
			horizontal[k] = "top"
		}
		// Right
		if !dfs(x+1, y) {
			k := fmt.Sprintf("%v,%v", y, x+1)
			vertical[k] = "right"
		}
		// Down
		if !dfs(x, y+1) {
			k := fmt.Sprintf("%v,%v", y+1, x)
			horizontal[k] = "down"
		}
		// Left
		if !dfs(x-1, y) {
			k := fmt.Sprintf("%v,%v", y, x)
			vertical[k] = "left"
		}

		return true
	}

	var sum int
	for y, line := range grid {
		for x, c := range line {
			key := fmt.Sprintf("%v-%v", y, x)
			if visited[key] {
				continue
			}

			current = c
			currentArea = 0
			currentSides = 0
			horizontal = map[string]string{}
			vertical = map[string]string{}
			dfs(x, y)

			for k, dir := range horizontal {
				var y, x int
				fmt.Sscanf(k, "%v,%v", &y, &x)

				if val, ok := horizontal[fmt.Sprintf("%v,%v", y, x-1)]; !ok || val != dir {
					currentSides++
				}
			}

			for k, dir := range vertical {
				var y, x int
				fmt.Sscanf(k, "%v,%v", &y, &x)

				if val, ok := vertical[fmt.Sprintf("%v,%v", y-1, x)]; !ok || val != dir {
					currentSides++
				}
			}

			sum += currentArea * currentSides
		}
	}

	fmt.Printf("part 2: %v\n", sum)
}
