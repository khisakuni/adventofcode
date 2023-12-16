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

	visited := map[string]map[int]bool{}

	var count int
	queue := [][]int{{0, 0, east}}
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

	// for k := range visited {
	// 	fmt.Printf("visited: %v\n", k)
	// 	parts := strings.Split(k, "-")
	// 	y, _ := strconv.Atoi(parts[0])
	// 	x, _ := strconv.Atoi(parts[1])
	// 	grid[y][x] = '#'
	// }
	//
	// for _, row := range grid {
	// 	fmt.Printf("%v\n", string(row))
	// }

	fmt.Printf("count: %v\n", count)
}

func fmtKey(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}
