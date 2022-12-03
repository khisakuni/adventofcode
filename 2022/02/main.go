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
	}

	shapeToPoints = map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	shapeToLose = map[string]string{
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
		outcome := parts[1]

		switch outcome {
		case "X":
			self := shapeToLose[opponent]
			points += shapeToPoints[self]
			// Lose
		case "Y":
			// Tie
			points += shapeToPoints[opponent] + 3
		case "Z":
			// Win
			self := shapeToLose[shapeToLose[opponent]]
			points += shapeToPoints[self] + 6
		}
	}

	fmt.Printf("Points: %d\n", points)
}
