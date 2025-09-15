package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	// "os"
)

//go:embed input.txt
var input string

func main() {
	re := regexp.MustCompile(`(mul\(\d+,\d+\))`)
	numRe := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(input, -1)
	var total int
	for _, match := range matches {
		mul := match
		nums := numRe.FindAllString(mul, 2)

		s := 1
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}

			s *= num
		}

		total += s
	}

	fmt.Printf("TOTAL: %v\n", total)
}
