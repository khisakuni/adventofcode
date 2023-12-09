package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var sum int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		nums := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			nums[i] = num
		}

		sum += calc(nums)
	}

	fmt.Printf("sum: %v\n", sum)
}

func calc(nums []int) int {
	var diffs []int

	allZeros := true
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		// s += diff
		if diff != 0 {
			allZeros = false
		}
		diffs = append(diffs, diff)
	}

	if allZeros {
		return nums[0]
	}

	res := calc(diffs)

	return nums[len(nums)-1] + res
}
