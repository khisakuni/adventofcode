package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// var input = `125 17`

func main() {
	part1()
}

func part1() {
	stones := strings.Split(input, " ")

	for i := 0; i < 25; i++ {
		var nextStones []string
		for _, stone := range stones {
			switch {
			case stone == "0":
				nextStones = append(nextStones, "1")
				// fmt.Printf("stone: %v -> %v\n", stone, nextStones[len(nextStones)-1])
			case len(stone)%2 == 0:
				left := strings.TrimLeft(stone[:len(stone)/2], "0")
				if left == "" {
					left = "0"
				}
				right := strings.TrimLeft(stone[len(stone)/2:], "0")
				if right == "" {
					right = "0"
				}
				nextStones = append(nextStones, left, right)
				// fmt.Printf("stone: %v -> %v, %v\n", stone, nextStones[len(nextStones)-1], nextStones[len(nextStones)-2])
			default:
				num, _ := strconv.Atoi(stone)
				num *= 2024
				nextStones = append(nextStones, strconv.Itoa(num))
				// fmt.Printf("stone: %v -> %v\n", stone, nextStones[len(nextStones)-1])
			}

		}

		stones = nextStones
		// fmt.Printf("%v\n", nextStones)
	}

	fmt.Printf("part 1: %v\n", len(stones))
}
