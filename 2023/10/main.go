package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	var queue [][]int

	// Look for S
	var i, j int
	for i < len(lines) {
		j = 0
		line := lines[i]
		for j < len(line) {
			if lines[i][j] == 'S' {
				queue = append(queue, []int{0, i, j})
				[]byte(line)[j] = '0'
			}

			j++
		}

		i++
	}

	max := 0

	fmt.Printf("start: (%v, %v)\n", queue[0][0], queue[0][1])

	for len(queue) > 0 {
		var current []int
		current, queue = queue[0], queue[1:]

		count, i, j := current[0], current[1], current[2]

		if count > max {
			max = count
		}

		char := lines[i][j]

		line := []byte(lines[i])
		line[j] = 'x'
		lines[i] = string(line)

		switch char {
		case 'x':
			continue
		case '|':
			// Check up
			if i > 0 {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' {
					queue = append(queue, []int{count + 1, i - 1, j})
				}
			}

			// Check down
			if i < len(lines)-1 {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' {
					queue = append(queue, []int{count + 1, i + 1, j})
				}
			}
		case '-':
			// Check left
			if j > 0 {
				next := lines[i][j-1]
				if next == 'L' || next == 'F' || next == '-' {
					queue = append(queue, []int{count + 1, i, j - 1})
				}
			}

			// Check right
			if j < len(lines[0])-1 {
				next := lines[i][j+1]
				if next == '7' || next == 'J' || next == '-' {
					queue = append(queue, []int{count + 1, i, j + 1})
				}
			}
		case '7':
			// Check down
			if i < len(lines)-1 {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' {
					queue = append(queue, []int{count + 1, i + 1, j})
				}
			}

			// Check left
			if j > 0 {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' {
					queue = append(queue, []int{count + 1, i, j - 1})
				}
			}
		case 'F':
			// Check down
			if i < len(lines)-1 {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' {
					queue = append(queue, []int{count + 1, i + 1, j})
				}
			}

			// Check right
			if i < len(lines[0])-1 {
				next := lines[i][j+1]
				if next == '-' || next == 'J' || next == '7' {
					queue = append(queue, []int{count + 1, i, j + 1})
				}
			}
		case 'J':
			// Check up
			if i > 0 {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' {
					queue = append(queue, []int{count + 1, i - 1, j})
				}
			}

			// Check left
			if j > 0 {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' {
					queue = append(queue, []int{count + 1, i, j - 1})
				}
			}
		case 'L':
			// Check up
			if i > 0 {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' {
					queue = append(queue, []int{count + 1, i - 1, j})
				}
			}

			// Check right
			if j < len(lines[0])-1 {
				next := lines[i][j+1]
				if next == '-' || next == '7' || next == 'J' {
					queue = append(queue, []int{count + 1, i, j + 1})
				}
			}
		case 'S':
			// Check up
			if i > 0 {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' {
					queue = append(queue, []int{count + 1, i - 1, j})
				}
			}

			// Check right
			if j < len(lines[0])-1 {
				next := lines[i][j+1]
				if next == '-' || next == '7' || next == 'J' {
					queue = append(queue, []int{count + 1, i, j + 1})
				}
			}

			// Check down
			if i < len(lines)-1 {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' {
					queue = append(queue, []int{count + 1, i + 1, j})
				}
			}

			// Check left
			if j > 0 {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' {
					queue = append(queue, []int{count + 1, i, j - 1})
				}
			}
		}
	}

	fmt.Printf("max: %v\n", max)
}
