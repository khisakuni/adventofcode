package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := bytes.Split(data, []byte{'\n'})
	rows := lines

	cache := map[string]int{}
	indexToRows := map[int][][]byte{}

	directions := 4
	cycles := 1000000000

	var prefixLen, cycleLen, i int
	for i < cycles {
		h := hash(rows)
		if dupe, ok := cache[h]; ok {
			// fmt.Printf("dupe of %v\n", dupe)
			prefixLen = dupe
			cycleLen = i - prefixLen
			break
		} else {
			cache[h] = i
			indexToRows[i] = rows
		}

		for j := 0; j < directions; j++ {
			rows = rotate(rows)
			roll(rows)
		}

		i++
	}

	n := ((cycles - prefixLen) % cycleLen) + prefixLen
	rows = indexToRows[n]

	var sum int
	rows = rotate(rows)
	for _, row := range rows {
		for j, c := range row {
			if c == 'O' {
				sum += j + 1
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func score(grid [][]byte) int {
	dupe := make([][]byte, len(grid))
	copy(dupe, grid)
	dupe = rotate(dupe)
	var sum int
	for _, row := range dupe {
		for j, c := range row {
			if c == 'O' {
				sum += j + 1
			}
		}
	}
	return sum
}

func rotate(grid [][]byte) [][]byte {
	rows := make([][]byte, len(grid[0]))
	for _, line := range grid {
		if len(line) == 0 {
			continue
		}

		for i, c := range line {
			rows[i] = append([]byte{byte(c)}, rows[i]...)
		}
	}

	return rows
}

func roll(grid [][]byte) {
	for _, row := range grid {
		j := len(row) - 1
		max := j
		for j >= 0 {
			if row[j] == 'O' {
				if j < max {
					row[max] = 'O'
					row[j] = '.'
				}

				max--
			}

			if row[j] == '#' {
				max = j - 1
			}

			j--
		}
	}
}

func hash(grid [][]byte) string {
	return string(bytes.Join(grid, []byte{'\n'}))
}
