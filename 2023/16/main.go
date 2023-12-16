package main

import (
	"bytes"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

const (
	north = iota
	east
	south
	west
)

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := bytes.Split(data, []byte{'\n'})
	grid = grid[:len(grid)-1]

	var max int

	// max = sum(grid, []int{0, 3, south})

	for i := 0; i < len(grid); i++ {
		c := copyGrid(grid)

		// for _, row := range c {
		// 	fmt.Printf("%v\n", string(row))
		// }
		s := sum(c, []int{i, 0, east})
		// fmt.Printf("sum: %v\n", s)
		if s > max {
			max = s
		}
	}

	for i := 0; i < len(grid); i++ {
		c := copyGrid(grid)
		s := sum(c, []int{i, len(grid[0]) - 1, west})
		// fmt.Printf("sum: %v\n", s)
		if s > max {
			max = s
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		c := copyGrid(grid)
		s := sum(c, []int{0, i, south})
		// fmt.Printf("sum: %v\n", s)
		if s > max {
			max = s
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		c := copyGrid(grid)
		s := sum(c, []int{len(grid) - 1, i, north})
		// fmt.Printf("sum: %v\n", s)
		if s > max {
			max = s
		}
	}

	fmt.Printf("count: %v\n", max)
}

func copyGrid(src [][]byte) [][]byte {
	grid := make([][]byte, len(src))
	for i, s := range src {
		n := make([]byte, len(s))
		copy(n, s)
		grid[i] = n
	}
	return grid
}

func sum(grid [][]byte, start []int) int {
	// fmt.Printf("start: %v\n", start)
	visited := map[string]map[int]bool{}
	queue := [][]int{start}
	var count int
	for len(queue) > 0 {
		var coords []int
		coords, queue = queue[0], queue[1:]

		y, x, d := coords[0], coords[1], coords[2]
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
			continue
		}

		c := grid[y][x]
		switch c {
		case '.', '#':
			if c == '.' {
				count++
			}
			grid[y][x] = '#'
			switch d {
			case north:
				queue = append(queue, []int{y - 1, x, north})
			case east:
				queue = append(queue, []int{y, x + 1, east})
			case south:
				queue = append(queue, []int{y + 1, x, south})
			case west:
				queue = append(queue, []int{y, x - 1, west})
			}

		case '|':
			k := fmtKey(y, x)
			m, ok := visited[k]
			if !ok {
				count++
				m = map[int]bool{}
			}

			if !m[d] {
				switch d {
				case north:
					queue = append(queue, []int{y - 1, x, north})
				case south:
					queue = append(queue, []int{y + 1, x, south})
				case east, west:
					queue = append(queue, []int{y - 1, x, north})
					queue = append(queue, []int{y + 1, x, south})
				}
			}

			m[d] = true
			visited[k] = m
		case '-':
			k := fmtKey(y, x)
			m, ok := visited[k]
			if !ok {
				count++
				m = map[int]bool{}
			}

			if !m[d] {
				switch d {
				case north, south:
					queue = append(queue, []int{y, x + 1, east})
					queue = append(queue, []int{y, x - 1, west})
				case east:
					queue = append(queue, []int{y, x + 1, east})
				case west:
					queue = append(queue, []int{y, x - 1, west})
				}
			}
			m[d] = true
			visited[k] = m
		case '/':
			k := fmtKey(y, x)
			m, ok := visited[k]
			if !ok {
				count++
				m = map[int]bool{}
			}

			if !m[d] {
				switch d {
				case north:
					queue = append(queue, []int{y, x + 1, east})
				case east:
					queue = append(queue, []int{y - 1, x, north})
				case south:
					queue = append(queue, []int{y, x - 1, west})
				case west:
					queue = append(queue, []int{y + 1, x, south})
				}
			}

			m[d] = true
			visited[k] = m
		case '\\':
			k := fmtKey(y, x)
			m, ok := visited[k]
			if !ok {
				count++
				m = map[int]bool{}
			}

			if !m[d] {
				switch d {
				case north:
					queue = append(queue, []int{y, x - 1, west})
				case east:
					queue = append(queue, []int{y + 1, x, south})
				case south:
					queue = append(queue, []int{y, x + 1, east})
				case west:
					queue = append(queue, []int{y - 1, x, north})
				}
			}

			m[d] = true
			visited[k] = m
		}
	}

	return count
}

func fmtKey(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}
