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

	var found bool

	// Look for S
	var i, j int
	for i < len(lines) {
		j = 0
		line := lines[i]
		for j < len(line) {
			if lines[i][j] == 'S' {
				found = true
				break
				// queue = append(queue, []int{0, i, j})
				// []byte(line)[j] = '0'
			}

			j++
		}

		if found {
			break
		}

		i++
	}

	startI := i
	startJ := j

	isStart := func(nextI, nextJ int) bool {
		return nextI == startI && nextJ == startJ
	}

	var direction string

	for !visited(lines[i][j]) {
		// fmt.Printf("CHAR: %v\n", string(lines[i][j]))
		// fmt.Println(strings.Join(lines, "\n"))
		char := lines[i][j]
		switch char {
		case '|':
			// Check up
			if direction == "up" {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' || visited(next) {
					if isStart(i-1, j) {
						l := []byte(lines[i-1])
						l[j] = '^'
						lines[i-1] = string(l)
					}
					direction = "up"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					i = i - 1
					continue
				}
			}

			// Check down
			if direction == "down" {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' || visited(next) {
					if isStart(i+1, j) {
						l := []byte(lines[i+1])
						l[j] = 'v'
						lines[i+1] = string(l)
					}
					direction = "down"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					i = i + 1
					continue
				}
			}
		case '-':
			// Check left
			if direction == "left" {
				next := lines[i][j-1]
				if next == 'L' || next == 'F' || next == '-' || visited(next) {
					direction = "left"
					l := []byte(lines[i])
					l[j] = '<'
					lines[i] = string(l)
					j = j - 1
					continue
				}
			}

			// Check right
			if direction == "right" {
				next := lines[i][j+1]
				if next == '7' || next == 'J' || next == '-' || visited(next) {
					direction = "right"
					l := []byte(lines[i])
					l[j] = '>'
					lines[i] = string(l)
					j = j + 1
					continue
				}
			}
		case '7':
			// Check down
			if direction == "right" {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' || visited(next) {
					// if isStart(i+1, j) {
					// 	l := []byte(lines[i+1])
					// 	l[j] = 'v'
					// 	lines[i+1] = string(l)
					// }
					direction = "down"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					i = i + 1
					continue
				}
			}

			// Check left
			if direction == "up" {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' || visited(next) {
					direction = "left"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					j = j - 1
					continue
				}
			}
		case 'F':
			// Check down
			if direction == "left" {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' || visited(next) {
					direction = "down"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					i = i + 1
					continue
				}
			}

			// Check right
			if direction == "up" {
				next := lines[i][j+1]
				if next == '-' || next == 'J' || next == '7' || visited(next) {
					direction = "right"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					j = j + 1
					continue
				}
			}
		case 'J':
			// Check up
			if direction == "right" {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' || visited(next) {
					direction = "up"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					i = i - 1
					continue
				}
			}

			// Check left
			if direction == "down" {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' || visited(next) {
					direction = "left"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					j = j - 1
					continue
				}
			}
		case 'L':
			// Check up

			if direction == "left" {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' || visited(next) {
					direction = "up"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					i = i - 1
					continue
				}
			}

			// Check right
			if direction == "down" {
				next := lines[i][j+1]
				if next == '-' || next == '7' || next == 'J' || visited(next) {
					direction = "right"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					j = j + 1
					continue
				}
			}
		case 'S':
			// Check up
			if i > 0 {
				next := lines[i-1][j]
				if next == '|' || next == '7' || next == 'F' || visited(next) {
					direction = "up"
					l := []byte(lines[i])
					l[j] = '^'
					lines[i] = string(l)
					i = i - 1
					continue
				}
			}

			// Check right
			if j < len(lines[0])-1 {
				next := lines[i][j+1]
				if next == '-' || next == '7' || next == 'J' || visited(next) {
					direction = "right"
					l := []byte(lines[i])
					l[j] = '>'
					lines[i] = string(l)
					j = j + 1
					continue
				}
			}

			// Check down
			if i < len(lines)-1 {
				next := lines[i+1][j]
				if next == '|' || next == 'L' || next == 'J' || visited(next) {
					direction = "down"
					l := []byte(lines[i])
					l[j] = 'v'
					lines[i] = string(l)
					i = i + 1
					continue
				}
			}

			// Check left
			if j > 0 {
				next := lines[i][j-1]
				if next == '-' || next == 'L' || next == 'F' || visited(next) {
					direction = "left"
					l := []byte(lines[i])
					l[j] = '<'
					lines[i] = string(l)
					j = j - 1
					continue
				}
			}
		}
	}

	fmt.Println(strings.Join(lines, "\n"))

	var total int
	for i, line := range lines {
		var count int
		var prev byte
		for j, c := range line {
			switch c {
			case '^':
				// fmt.Printf("(%d, %d) %v, prev: %v\n", i, j, string(c), string(prev))
				if prev == 'v' || prev == 0 {
					count++
				}
				prev = '^'
			case 'v':
				// fmt.Printf("(%d, %d) %v, prev: %v\n", i, j, string(c), string(prev))
				if prev == '^' || prev == 0 {
					count++
				}
				prev = 'v'
			case '<', '>':
			default:
				l := []byte(lines[i])
				if count%2 != 0 {
					total++
					l[j] = 'I'
				} else {
					l[j] = 'O'
				}
				lines[i] = string(l)
			}
		}
	}

	fmt.Println(strings.Join(lines, "\n"))

	fmt.Printf("total: %v\n", total)

}

func visited(char byte) bool {
	return char == '^' || char == '>' || char == 'v' || char == '<'
}
