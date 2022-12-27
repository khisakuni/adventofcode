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
	//R 4
	//U 4
	//L 3
	//D 1
	//R 4
	//D 1
	//L 5
	//R 2
	//	`)

	directions := map[string]image.Point{
		"U": {X: 0, Y: -1},
		"R": {X: 1, Y: 0},
		"D": {X: 0, Y: 1},
		"L": {X: -1, Y: 0},
	}

	rope := make([]image.Point, 2)
	visited := map[image.Point]bool{
		// Starting position.
		image.Point{}: true,
	}
	strs := strings.Split(string(input), "\n")
	for _, str := range strs {
		if str == "" {
			continue
		}

		// Parse the direction and count from input.
		var d string
		var count int
		_, _ = fmt.Sscanf(str, "%s %d", &d, &count)

		// Get the direction in image.Point form.
		direction := directions[d]

		fmt.Printf(">> %v, %v\n", d, count)

		for i := 0; i < count; i++ {
			// Update head.
			rope[0] = rope[0].Add(direction)

			// Iterate over rest of rope.
			for j := 1; j < len(rope); j++ {
				// Get the difference between one ahead and this one.
				diff := rope[j-1].Sub(rope[j])

				fmt.Printf("diff: %v\n", diff)

				// If vertical or horizontal difference is more than 1, need to move.
				if diff.X > 1 || diff.X < -1 || diff.Y > 1 || diff.Y < -1 {

					// Get the next X.
					var x int
					switch {
					case diff.X < 0:
						x = -1
					case diff.X > 0:
						x = 1
					default:
						x = 0
					}

					// Get the next Y.
					var y int
					switch {
					case diff.Y < 0:
						y = -1
					case diff.Y > 0:
						y = 1
					default:
						y = 0
					}

					rope[j] = rope[j].Add(image.Point{
						X: x,
						Y: y,
					})
				}

				if j == len(rope)-1 {
					visited[rope[j]] = true
				}
			}
		}
	}

	fmt.Printf("visited: %v\n", len(visited))
}
