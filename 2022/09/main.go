package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
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
	//`)

	directions := map[string]image.Point{
		"U": {X: 0, Y: -1},
		"R": {X: 1, Y: 0},
		"D": {X: 0, Y: 1},
		"L": {X: -1, Y: 0},
	}

	head := image.Point{X: 0, Y: 0}
	tail := head
	visited := map[image.Point]bool{
		tail: true,
	}
	strs := strings.Split(string(input), "\n")
	for _, str := range strs {
		if str == "" {
			continue
		}

		// Get the direction name and counhead.
		parts := strings.Split(str, " ")
		d := parts[0]
		count, _ := strconv.Atoi(parts[1])

		// Get the direction in image.Point form.
		direction := directions[d]

		for i := 0; i < count; i++ {
			// Move the head.
			head = head.Add(direction)

			// Check if bottom needs to move.
			//r := image.Rectangle{
			//	Min: image.Point{
			//		X: head.X - 1,
			//		Y: head.Y - 1,
			//	},
			//	Max: image.Point{
			//		X: head.X + 1,
			//		Y: head.Y + 1,
			//	},
			//}

			dist := math.Sqrt(math.Pow(float64(head.X-tail.X), 2) + math.Pow(float64(head.Y-tail.Y), 2))
			//fmt.Printf("dist: %v\n", dist)
			if dist >= 2 {
				// Move the bottom.
				switch d {
				case "U":
					tail = image.Point{
						X: head.X,
						Y: head.Y + 1,
					}
				case "R":
					tail = image.Point{
						X: head.X - 1,
						Y: head.Y,
					}
				case "D":
					tail = image.Point{
						X: head.X,
						Y: head.Y - 1,
					}
				case "L":
					tail = image.Point{
						X: head.X + 1,
						Y: head.Y,
					}
				}

				// Add entry to map.
				visited[tail] = true
			}

			//fmt.Printf(">> direction: %v, count: %v, head: %v, tail: %v\n", d, count, head, tail)
		}

	}

	fmt.Printf("visited: %v\n", len(visited))
}
