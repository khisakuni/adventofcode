package main

import (
	"io/ioutil"
	"math/bits"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	groups := strings.Split(string(data), "\n\n")
	var total int
	for _, group := range groups {
		groupFlags := ^((1 << ('z' - 'a')) & 0)
		for _, line := range strings.Split(group, "\n") {
			if line == "" {
				continue
			}
			flags := 0
			for _, c := range line {
				val := c - 'a'
				if val < 0 || val > 25 {
					continue
				}
				if (flags & (1 << val)) > 0 {
					continue
				}
				flags |= 1 << val
			}
			groupFlags &= flags
		}
		total += bits.OnesCount(uint(groupFlags))
	}
}
