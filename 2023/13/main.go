package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	blocks := strings.Split(string(data), "\n\n")

	var sum int
	for blockNum, block := range blocks {
		// fmt.Printf("BLOCK %v\n", blockNum)

		// Consider rows
		var found bool
		lines := strings.Split(block, "\n")
		for i, line := range lines[:len(lines)-1] {
			// Found potential match
			// fmt.Printf("comparing:\n%v\n%v\n\n", line, lines[i+1])
			if line == lines[i+1] {
				// fmt.Printf("found match!\n")
				j := i
				k := i + 1
				for k < len(lines) && j >= 0 {
					// fmt.Printf("comparing\n%v\n%v\n\n", lines[j], lines[k])
					if lines[j] != lines[k] {
						// fmt.Printf("not the same\n%v\n%v\n\n", lines[j], lines[k])
						break
					}
					k++
					j--
				}
				// fmt.Printf("k >>>>>>> %v, %v\n", k, len(lines)-1)
				if k == len(lines) || j == -1 {
					// fmt.Printf("horizontal sum: %v\n", i+1)
					sum += (i + 1) * 100
					found = true
					break
				}
			}
		}

		if found {
			continue
		}

		// fmt.Printf("VERTICAL\n")
		// Consider columns
		// Rotate image 90 deg.
		rows := make([][]byte, len(lines[0]))
		for _, line := range lines {
			for i, c := range line {
				rows[i] = append([]byte{byte(c)}, rows[i]...)
			}
		}

		for i, line := range rows[:len(rows)-1] {
			// Found potential match
			// fmt.Printf("comparing:\n%v\n%v\n\n", string(line), string(rows[i+1]))
			if string(line) == string(rows[i+1]) {
				// fmt.Printf("found match!\n")
				j := i
				k := i + 1
				for k < len(rows) && j >= 0 {
					if string(rows[j]) != string(rows[k]) {
						// fmt.Printf("not the same\n%v\n%v\n\n", string(rows[j]), string(rows[k]))
						break
					}
					k++
					j--
				}
				if k == len(rows) || j == -1 {
					sum += i + 1
					// fmt.Printf("vertical sum: %v\n", i+1)
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Printf("did not find: %v\n", blockNum)
		}
	}

	fmt.Printf("sum: %v\n", sum)

}
