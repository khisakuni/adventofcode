package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// var input = `....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#..^.....
// ........#.
// #.........
// ......#...`

func main() {
	part1()
	fmt.Println("")
	part2()
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

func part2() {
	lines := strings.Split(input, "\n")
	grid := [][]byte{}
	var x, y int
	var startX, startY int
	var current, startCurrent byte
	for i, line := range lines {
		grid = append(grid, []byte(line))

		for j, c := range line {
			if c == '^' || c == '>' || c == 'v' || c == '<' {
				x = j
				y = i
				startX = x
				startY = y
				current = byte(c)
				startCurrent = current
			}

		}
	}

	const (
		up = 1 << iota
		right
		down
		left
	)

	path := map[string]struct{}{}
	for (x > 0 && x < len(grid[0])-1) && (y > 0 && y < len(grid)-1) {
		nextX := x
		nextY := y
		switch current {
		case '^':
			nextY--
			switch grid[nextY][nextX] {
			case '#':
				current = '>'
				continue
			}

		case '>':
			nextX++
			switch grid[nextY][nextX] {
			case '#':
				current = 'v'
				continue
			}
		case 'v':
			nextY++
			switch grid[nextY][nextX] {
			case '#':
				current = '<'
				continue
			}
		case '<':
			nextX--

			switch grid[nextY][nextX] {
			case '#':
				current = '^'
				continue
			}
		}

		x = nextX
		y = nextY
		if x == startX && y == startY {
			continue
		}

		path[fmt.Sprintf("%v-%v", x, y)] = struct{}{}
	}

	var count int
	for k := range path {
		parts := strings.Split(k, "-")
		blockX, _ := strconv.Atoi(parts[0])
		blockY, _ := strconv.Atoi(parts[1])

		grid[blockY][blockX] = '#'

		x := startX
		y := startY

		current = startCurrent
		visited := map[string]bool{}
		for (x > 0 && x < len(grid[0])-1) && (y > 0 && y < len(grid)-1) {
			nextX := x
			nextY := y
			switch current {
			case '^':
				nextY--
				switch grid[nextY][nextX] {
				case '#':
					current = '>'
					continue
				}

			case '>':
				nextX++
				switch grid[nextY][nextX] {
				case '#':
					current = 'v'
					continue
				}
			case 'v':
				nextY++
				switch grid[nextY][nextX] {
				case '#':
					current = '<'
					continue
				}
			case '<':
				nextX--
				switch grid[nextY][nextX] {
				case '#':
					current = '^'
					continue
				}
			}

			x = nextX
			y = nextY
			key := fmt.Sprintf("%v-%v-%v", x, y, current)
			didVisit := visited[key]
			if didVisit {
				count++
				break
			}

			visited[key] = true
		}

		grid[blockY][blockX] = '.'
	}

	fmt.Printf("part 2: %d\n", count)
}
