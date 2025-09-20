package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// var input = `............
// ........0...
// .....0......
// .......0....
// ....0.......
// ......A.....
// ............
// ............
// ........A...
// .........A..
// ............
// ............`

// var input = `..........
// ..........
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ..........
// ..........
// ..........`

func main() {
	part1()

	n := time.Now()
	part2()
	fmt.Printf("%s\n", time.Since(n))
}

func part2() {

	lines := strings.Split(input, "\n")
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	freqToAntennas := map[rune][][]int{}
	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}

			freqToAntennas[c] = append(freqToAntennas[c], []int{x, y})
		}
	}

	m := map[string]bool{}

	for _, antennas := range freqToAntennas {
		for _, coords := range antennas {
			for _, other := range antennas {
				x, y := coords[0], coords[1]
				xo, yo := other[0], other[1]

				if x == xo && y == yo {
					continue
				}

				xDist := x - xo
				if xDist < 0 {
					xDist *= -1
				}

				yDist := y - yo
				if yDist < 0 {
					yDist *= -1
				}

				var ax, axo int
				if x < xo {
					ax = -1 * xDist
					axo = xDist
				} else {
					ax = xDist
					axo = -1 * xDist
				}

				var ay, ayo int
				if y < yo {
					ay = -1 * yDist
					ayo = yDist
				} else {
					ay = yDist
					ayo = -1 * yDist
				}

				for x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines) {
					m[fmt.Sprintf("%d-%d", x, y)] = true
					grid[y][x] = '#'

					y = y + ay
					x = x + ax
				}

				for xo >= 0 && xo < len(lines[0]) && yo >= 0 && yo < len(lines) {
					m[fmt.Sprintf("%d-%d", xo, yo)] = true
					grid[yo][xo] = '#'

					xo = xo + axo
					yo = yo + ayo
				}

				// for _, l := range grid {
				// 	fmt.Printf("%s\n", string(l))
				// }

				// fmt.Println("")
			}

		}
	}

	// for _, l := range grid {
	// 	fmt.Printf("%s\n", string(l))
	// }
	// fmt.Println("")

	fmt.Printf("part 2: %v\n", len(m))
}

func part1() {
	lines := strings.Split(input, "\n")
	grid := [][]rune{}
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	freqToAntennas := map[rune][][]int{}
	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}

			freqToAntennas[c] = append(freqToAntennas[c], []int{x, y})
		}
	}

	m := map[string]bool{}

	for _, antennas := range freqToAntennas {
		for _, coords := range antennas {
			for _, other := range antennas {
				x, y := coords[0], coords[1]
				xo, yo := other[0], other[1]

				if x == xo && y == yo {
					continue
				}

				xDist := x - xo
				if xDist < 0 {
					xDist *= -1
				}

				yDist := y - yo
				if yDist < 0 {
					yDist *= -1
				}

				var ax, axo int
				if x < xo {
					ax = x - xDist
					axo = xo + xDist
				} else {
					ax = x + xDist
					axo = xo - xDist
				}

				var ay, ayo int
				if y < yo {
					ay = y - yDist
					ayo = yo + yDist
				} else {
					ay = y + yDist
					ayo = yo - yDist
				}

				if ax >= 0 && ax < len(lines[0]) && ay >= 0 && ay < len(lines) {
					m[fmt.Sprintf("%d-%d", ax, ay)] = true
					grid[ay][ax] = '#'
				}

				if axo >= 0 && axo < len(lines[0]) && ayo >= 0 && ayo < len(lines) {
					m[fmt.Sprintf("%d-%d", axo, ayo)] = true
					grid[ayo][axo] = '#'
				}

				// for _, l := range grid {
				// 	fmt.Printf("%s\n", string(l))
				// }

				// fmt.Println("")
			}

		}
	}

	// for _, l := range grid {
	// 	fmt.Printf("%s\n", string(l))
	// }
	// fmt.Println("")

	fmt.Printf("part 1: %v\n", len(m))
}
