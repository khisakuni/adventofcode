package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	now := time.Now()
	part1()
	fmt.Printf("%v\n", time.Since(now))
	now = time.Now()
	part2()
	fmt.Printf("%v\n", time.Since(now))

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

func part2() {
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

			// concat operation
			position := 1
			n := next
			for n > 0 {
				n /= 10
				position *= 10
			}

			dfs((current*position)+next, remaining[1:])
		}

		dfs(0, nums)

		if found {
			num += target
		}

	}

	fmt.Printf("num: %d\n", num)
}
