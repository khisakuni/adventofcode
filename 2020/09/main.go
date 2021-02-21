package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	for i, line := range strings.Split(string(data), "\n") {
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
			fmt.Printf("    diff: %v\n", diff)
			if _, ok := m[int(diff)]; ok {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("found bad number: %d\n", val)
			return
		}

		// Clear out old entries in cache.
		delete(m, rm[i-limit])
		delete(rm, i-limit)

		m[val] = i
		rm[i] = val
	}
}
