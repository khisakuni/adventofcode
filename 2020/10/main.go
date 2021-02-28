package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	var nums []int
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	// Part 1
	sort.Ints(nums)
	oneDiffs := 1
	threeDiffs := 1
	for i := 0; i < len(nums) - 1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 1 {
			oneDiffs++
		}
		if diff == 3 {
			threeDiffs++
		}
	}

	fmt.Printf("one: %d, three: %d, %d\n", oneDiffs, threeDiffs, oneDiffs * threeDiffs)

	// Part 2
	// adapter -> num options
	m := map[int]int{}

	// Start with 0. Only 1 way to get there.
	m[0] = 1
	for _, num := range nums {
		// Iterate through possible diffs.
		for i := 1; i <= 3; i++ {

			// Sum option count for 1, 2, and 3 diffs away.
			m[num] += m[num - i]
		}
	}

	fmt.Printf("count: %d\n", m[nums[len(nums)-1]])
}
