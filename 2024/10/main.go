package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// var input = `89010123
// 78121874
// 87430965
// 96549874
// 45678903
// 32019012
// 01329801
// 10456732`

// var input = `..90..9
// ...1.98
// ...2..7
// 6543456
// 765.987
// 876....
// 987....`

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	var score int
	var dfs func(x, y int, num rune, visited map[string]bool) int
	dfs = func(x, y int, num rune, visited map[string]bool) int {
		if x < 0 || y < 0 {
			return 0
		}

		if x >= len(lines[0]) || y >= len(lines) {
			return 0
		}

		if grid[y][x] != num {
			return 0
		}

		if num == '9' {
			key := fmt.Sprintf("%v-%v", x, y)
			if visited[key] {
				return 0
			}
			visited[key] = true
			return 1
		}

		// Up
		localScore := dfs(x, y-1, num+1, visited)

		// Right
		localScore += dfs(x+1, y, num+1, visited)

		// Down
		localScore += dfs(x, y+1, num+1, visited)

		// Left
		localScore += dfs(x-1, y, num+1, visited)

		return localScore
	}

	for y, line := range grid {
		for x := range line {
			score += dfs(y, x, '0', map[string]bool{})
		}
	}
	fmt.Printf("part 1: %v\n", score)
}

func part2() {
	lines := strings.Split(input, "\n")
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	var score int
	var dfs func(x, y int, num rune) int
	dfs = func(x, y int, num rune) int {
		if x < 0 || y < 0 {
			return 0
		}

		if x >= len(lines[0]) || y >= len(lines) {
			return 0
		}

		if grid[y][x] != num {
			return 0
		}

		if num == '9' {
			return 1
		}

		// Up
		localScore := dfs(x, y-1, num+1)

		// Right
		localScore += dfs(x+1, y, num+1)

		// Down
		localScore += dfs(x, y+1, num+1)

		// Left
		localScore += dfs(x-1, y, num+1)

		return localScore
	}

	for y, line := range grid {
		for x := range line {
			score += dfs(y, x, '0')
		}
	}
	fmt.Printf("part 2: %v\n", score)
}
