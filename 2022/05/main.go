package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	//input := `
	//   [D]
	//[N] [C]
	//[Z] [M] [P]
	//1   2   3
	//
	//move 1 from 2 to 1
	//move 3 from 1 to 3
	//move 2 from 2 to 1
	//move 1 from 1 to 2
	//`

	strs := strings.Split(string(input), "\n")
	crateRe := regexp.MustCompile(`\[([A-Z])\]`)
	moveRe := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	crates := map[int]string{}
	for _, str := range strs {
		// Initialize crates
		cratePos := crateRe.FindAllStringSubmatchIndex(str, -1)
		for _, crate := range cratePos {
			start := crate[0]
			i := (start / 4) + 1
			val := str[crate[2]:crate[3]]
			crates[i] = val + crates[i]
		}

		moves := moveRe.FindAllStringSubmatch(str, -1)
		if len(moves) > 0 {
			count, _ := strconv.Atoi(moves[0][1])
			src, _ := strconv.Atoi(moves[0][2])
			dest, _ := strconv.Atoi(moves[0][3])
			//fmt.Printf(">> %v, %v, %v, %v\n", moves[0], count, src, dest)
			srcStr := crates[src]
			//fmt.Printf("%v -> %v\n", crates[dest], crates[dest]+string(srcStr[len(srcStr)-1]))
			crates[dest] = crates[dest] + srcStr[len(srcStr)-count:]
			crates[src] = srcStr[:len(srcStr)-count]
		}

		//fmt.Printf("> %v\n", crates)
	}

	result := make([]string, len(crates))
	for i, c := range crates {
		//fmt.Printf("%d: %v\n", i, c)
		result[i-1] = string(c[len(c)-1])
	}
	fmt.Printf("Result: %v\n", strings.Join(result, ""))
}
