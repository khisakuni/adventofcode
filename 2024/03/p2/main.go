package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	re := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	numRe := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(input, -1)
	var total int
	enabled := true
	for _, match := range matches {
		switch {
		case strings.HasPrefix(match, "mul("):
			if !enabled {
				continue
			}

			nums := numRe.FindAllString(match, 2)

			s := 1
			for _, n := range nums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}

				s *= num
			}

			total += s
		case strings.HasPrefix(match, "do("):
			enabled = true
		default:
			enabled = false

		}
	}

	fmt.Printf("TOTAL: %v\n", total)
}
