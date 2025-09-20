package main

import (
	_ "embed"
	"fmt"
	"strings"
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

	fmt.Printf("part: %v\n", len(m))
}
