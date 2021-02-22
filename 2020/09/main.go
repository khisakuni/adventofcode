package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const limit = 25

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// val -> index
	m := map[int]int{}

	// index -> val
	rm := map[int]int{}

	lines := strings.Split(string(data), "\n")
	badIndex := -1
	for i, line := range lines {
		if line == "" {
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// Preamble
		if len(m) < limit {
			m[val] = i
			rm[i] = val
			continue
		}

		var found bool
		for k := range m {
			diff := math.Abs(float64(k - val))
			//fmt.Printf("    diff: %v\n", diff)
			if _, ok := m[int(diff)]; ok {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("found bad number: %d\n", val)
			badIndex = i
			break
		}

		// Clear out old entries in cache.
		delete(m, rm[i-limit])
		delete(rm, i-limit)

		m[val] = i
		rm[i] = val
	}

	fmt.Printf("bad index: %v\n", badIndex)
	badNumber := mustParse(lines[badIndex])
	fmt.Printf("bad number: %d\n", badNumber)

	start := 0
	end := start
	for start <= badIndex-2 {
		sum := mustParse(lines[start]) //+ mustParse(lines[end])
		//fmt.Printf("sum: %d, %d\n", sum, start)
		for end <= badIndex && sum < badNumber {
			end++
			sum += mustParse(lines[end])
		}
		if sum == badNumber {
			fmt.Printf("found range: %v to %v\n", start, end)
			break
		}
		start++
		end = start
	}

	r := lines[start:end+1]
	nums := make([]int, len(r))
	for i, str := range r {
		nums[i] = mustParse(str)
	}
	sort.Ints(nums)
	fmt.Printf("smallest: %d, largest: %d, sum: %d\n", nums[0], nums[len(nums)-1], nums[0] + nums[len(nums)-1])

}

func mustParse(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
