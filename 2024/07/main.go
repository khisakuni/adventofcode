package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
}

func part1() {
	var num int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		numStrs := strings.Split(strings.TrimSpace(parts[1]), " ")
		nums := make([]int, len(numStrs))
		for i, str := range numStrs {
			num, _ := strconv.Atoi(str)
			nums[i] = num
		}

		var dfs func(int, []int)

		var found bool
		dfs = func(current int, remaining []int) {
			if found {
				return
			}

			if len(remaining) == 0 {
				if current == target {
					found = true
				}

				return
			}

			if current > target {
				return
			}

			next := remaining[0]
			dfs(current+next, remaining[1:])
			dfs(current*next, remaining[1:])
		}

		dfs(0, nums)

		if found {
			num += target
		}

	}

	fmt.Printf("num: %d\n", num)
}
