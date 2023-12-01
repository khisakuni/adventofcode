package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	regexes := []*regexp.Regexp{
		regexp.MustCompile(`(\d|one|three|four|five|six|seven)`),
		regexp.MustCompile(`(eight|nine)`),
		regexp.MustCompile(`(two)`),
	}

	m := map[string]int{
		"zero":  0,
		"0":     0,
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	var total int
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		firstStart := len(line)
		firstEnd := len(line)
		lastStart := -1
		lastEnd := -1
		for _, re := range regexes {
			indices := re.FindAllStringSubmatchIndex(line, -1)
			if len(indices) == 0 {
				continue
			}

			f := indices[0]
			if f[0] < firstStart {
				firstStart = f[0]
				firstEnd = f[1]
			}

			l := indices[len(indices)-1]
			if l[0] > lastStart {
				lastStart = l[0]
				lastEnd = l[1]
			}

		}

		first := m[line[firstStart:firstEnd]]
		last := m[line[lastStart:lastEnd]]

		sub := (first * 10) + last
		fmt.Printf("%v: %v\n", line, sub)

		total += sub
	}

	fmt.Printf("total: %v\n", total)
}
