package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// var input = `
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400

// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176

// Button A: X+17, Y+86
// Button B: X+84, Y+37
// Prize: X=7870, Y=6450

// Button A: X+69, Y+23
// Button B: X+27, Y+71
// Prize: X=18641, Y=10279`

func main() {
	part1()
}

func part1() {
	games := strings.Split(input, "\n\n")
	var sum int
	for _, game := range games {
		game = strings.TrimSpace(game)
		lines := strings.Split(game, "\n")

		var aX, aY int
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &aX, &aY)

		var bX, bY int
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bX, &bY)

		var pX, pY int
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &pX, &pY)

		tokens := -1
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				if ((a*aX + b*bX) == pX) && ((a*aY + b*bY) == pY) {
					toks := a*3 + b
					if tokens == -1 || toks < tokens {
						tokens = toks
					}
				}
			}
		}

		if tokens > -1 {
			sum += tokens
		}
	}

	fmt.Printf("part 1: %v\n", sum)
}
