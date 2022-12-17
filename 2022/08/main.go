package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	//	input = []byte(`
	//30373
	//25512
	//65332
	//33549
	//35390
	//`)

	trees := map[image.Point]int{}
	for y, str := range strings.Fields(strings.TrimSpace(string(input))) {
		for x, part := range str {
			trees[image.Point{
				X: x,
				Y: y,
			}] = int(part - '0')
		}
	}

	var totalVisible int
	for point, tree := range trees {
		visible := 0
		for _, direction := range []image.Point{
			// Top
			{0, -1},
			// Right
			{1, 0},
			// Bottom
			{0, 1},
			// Left
			{-1, 0},
		} {
			i := 1
			for true {
				neighbor, ok := trees[point.Add(direction.Mul(i))]

				// Hit the border
				if !ok {
					visible = 1
					break
				}

				// Neighboring tree is taller than this tree.
				if neighbor >= tree {
					break
				}

				i++
			}

		}

		totalVisible += visible
	}

	fmt.Printf("visible: %d\n", totalVisible)
}
