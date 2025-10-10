package main

import (
	_ "embed"
	"fmt"
	"strings"
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
