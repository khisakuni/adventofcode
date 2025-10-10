package main

import (
	_ "embed"
	"fmt"
	"strings"
	// "time"
)

// var input = `p=0,4 v=3,-3
// p=6,3 v=-1,-3
// p=10,3 v=-1,2
// p=2,0 v=2,-1
// p=0,0 v=1,3
// p=3,0 v=-2,-2
// p=7,6 v=-1,-3
// p=3,0 v=-1,-2
// p=9,3 v=2,3
// p=7,3 v=-1,2
// p=2,4 v=2,-3
// p=9,5 v=-3,-3`

//go:embed input.txt
var input string

const width = 101
const height = 103

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	var topLeft, topRight, bottomLeft, bottomRight int
	for _, line := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

		x = ((vx*100 + x) % width)
		if x < 0 {
			x = width + x
		}
		y = ((vy*100 + y) % height)
		if y < 0 {
			y = height + y
		}

		switch {
		case y < (height/2) && x < (width/2):
			topLeft++
		case y < (height/2) && x > (width/2):
			topRight++
		case y > (height/2) && x < (width/2):
			bottomLeft++
		case y > (height/2) && x > (width/2):
			bottomRight++
		}
	}

	fmt.Printf("part 1: %v\n", topLeft*topRight*bottomLeft*bottomRight)
}

func part2() {
	lines := strings.Split(input, "\n")
	var grid [][]bool
	for sec := 0; sec < 10000; sec++ {
		grid = [][]bool{}
		for i := 0; i < height; i++ {
			grid = append(grid, make([]bool, width))
		}
		for _, line := range lines {
			var x, y, vx, vy int
			fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

			x = ((vx*sec + x) % width)
			if x < 0 {
				x = width + x
			}
			y = ((vy*sec + y) % height)
			if y < 0 {
				y = height + y
			}
			grid[y][x] = true
		}

		area := 10
		for i := 0; i < height-area; i += area {
			for j := 0; j < width-area; j += area {

				var pop int
				for y := i; y < i+area; y++ {
					for x := j; x < j+area; x++ {
						if grid[y][x] {
							pop++
						}
					}
				}

				if float32(pop)/float32((area*area)) > 0.5 {
					fmt.Printf("part 2: %v\n", sec)
					print(grid)

					return
				}
			}

		}
	}

}

func print(grid [][]bool) {
	for _, line := range grid {
		for _, c := range line {
			if c {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
