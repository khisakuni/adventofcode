package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	codeToShape = map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	shapeToPoints = map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	shapeToBeat = map[string]string{
		"Rock":     "Scissors",
		"Paper":    "Rock",
		"Scissors": "Paper",
	}
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(input), "\n")
	var points int
	for _, str := range strs {
		if str == "" {
			continue
		}

		parts := strings.Split(str, " ")
		opponent := codeToShape[parts[0]]
		self := codeToShape[parts[1]]

		switch {
		case opponent == self:
			points += 3
		case opponent == shapeToBeat[self]:
			points += 6
		}

		points += shapeToPoints[self]
	}

	fmt.Printf("Points: %d\n", points)
}
