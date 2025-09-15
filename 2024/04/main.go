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
	part2()

}

func part2() {
	lines := strings.Split(input, "\n")
	var count int

	for y, line := range lines {
		for x, c := range line {
			if c == 'A' && y >= 1 && x >= 1 && x < len(line)-1 && y < len(lines)-1 {
				ltr := string([]byte{lines[y-1][x-1], lines[y+1][x+1]})
				rtl := string([]byte{lines[y-1][x+1], lines[y+1][x-1]})

				if (ltr == "MS" || ltr == "SM") && (rtl == "MS" || rtl == "SM") {
					count++
				}
			}

		}
	}

	fmt.Printf("part2: %d\n", count)
}

func part1() {
	lines := strings.Split(input, "\n")
	var count int
	var dfs func(x, y, xdir, ydir int, remaining []rune)

	dfs = func(x, y, xdir, ydir int, remaining []rune) {
		if x < 0 || y < 0 {
			return
		}
		if x >= len(lines[0]) || y >= len(lines) {
			return
		}

		if lines[y][x] != byte(remaining[0]) {
			return
		}

		remaining = remaining[1:]
		if len(remaining) == 0 {
			count++
			return
		}

		// Up
		dfs(x+xdir, y+ydir, xdir, ydir, remaining)
	}

	for y, line := range lines {
		for x, c := range line {
			if c == 'X' {
				// Up
				dfs(x, y, 0, -1, []rune{'X', 'M', 'A', 'S'})

				// Down
				dfs(x, y, 0, 1, []rune{'X', 'M', 'A', 'S'})

				// Left
				dfs(x, y, -1, 0, []rune{'X', 'M', 'A', 'S'})

				// Right
				dfs(x, y, 1, 0, []rune{'X', 'M', 'A', 'S'})

				// Up + right
				dfs(x, y, 1, -1, []rune{'X', 'M', 'A', 'S'})

				// Up + Left
				dfs(x, y, -1, -1, []rune{'X', 'M', 'A', 'S'})

				// Down + Right
				dfs(x, y, 1, 1, []rune{'X', 'M', 'A', 'S'})

				// Down + Left
				dfs(x, y, -1, 1, []rune{'X', 'M', 'A', 'S'})
			}

		}
	}

	fmt.Printf("found: %d\n", count)
}
