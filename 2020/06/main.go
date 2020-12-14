package main

import (
	"fmt"
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
		var flags uint = 0
		for _, c := range group {
			val := c - 'a'
			if val < 0 || val > 25 {
				continue
			}
			if (flags & (1 << val)) > 0 {
				continue
			}
			flags |= 1 << val
		}
		total += bits.OnesCount(flags)
	}
	fmt.Printf("total: %d\n", total)

}
