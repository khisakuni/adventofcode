package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
}

func part1() {
	lines := strings.Split(input, "\n")
	grid := [][]byte{}
	var x, y int
	for i, line := range lines {
		grid = append(grid, []byte(line))

		for j, c := range line {
			if c == '^' || c == '>' || c == 'v' || c == '<' {
				x = j
				y = i
			}

		}
	}

	visited := map[string]bool{}
	for (x > 0 && x < len(grid[0])-1) && (y > 0 && y < len(grid)-1) {
		visited[fmt.Sprintf("%d-%d", x, y)] = true
		c := grid[y][x]
		nextX := x
		nextY := y
		switch c {
		case '^':
			nextY--
			if grid[nextY][nextX] == '#' {
				grid[y][x] = '>'
				continue
			}
		case '>':
			nextX++
			if grid[nextY][nextX] == '#' {
				grid[y][x] = 'v'
				continue
			}
		case 'v':
			nextY++
			if grid[nextY][nextX] == '#' {
				grid[y][x] = '<'
				continue
			}
		case '<':
			nextX--
			if grid[nextY][nextX] == '#' {
				grid[y][x] = '^'
				continue
			}
		}

		grid[y][x] = 'X'
		x = nextX
		y = nextY
		grid[y][x] = c

	}

	fmt.Printf("part 1: %d", len(visited)+1)
}
