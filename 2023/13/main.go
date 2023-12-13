package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	blocks := strings.Split(string(data), "\n\n")

	var sum int
	for blockNum, block := range blocks {

		// Consider rows
		var found bool
		var fixed bool

		lines := strings.Split(block, "\n")

		// Convert to numbers
		nums := []uint32{}
		for _, line := range lines {
			if line == "" {
				continue
			}
			line = strings.ReplaceAll(line, "#", "1")
			line = strings.ReplaceAll(line, ".", "0")

			if len(line) < 32 {
				line = strings.Repeat("0", 32-len(line)) + line
			}

			num, err := strconv.ParseUint(line, 2, 32)
			if err != nil {
				panic(err)
			}

			nums = append(nums, uint32(num))
		}

		for i, num := range nums[:len(nums)-1] {
			isOneOff := bits.OnesCount32(num^nums[i+1]) == 1
			if num == nums[i+1] || isOneOff {
				j := i - 1
				k := i + 2
				for k < len(nums) && j >= 0 {
					oneOff := bits.OnesCount32(nums[j]^nums[k]) == 1
					if oneOff && !isOneOff {
						isOneOff = true
						k++
						j--
						continue
					}

					if nums[j] != nums[k] {
						break
					}

					k++
					j--
				}

				if (k == len(nums) || j == -1) && isOneOff {
					sum += (i + 1) * 100
					found = true
					if isOneOff {
						fixed = true
					}
					break
				}

			}
		}

		if found && fixed {
			continue
		}

		// Consider columns
		// Rotate image 90 deg.
		rowBytes := make([][]byte, len(lines[0]))
		for _, line := range lines {
			if line == "" {
				continue
			}

			for i, c := range line {
				b := '0'
				if c == '#' {
					b = '1'
				}
				rowBytes[i] = append([]byte{byte(b)}, rowBytes[i]...)
			}
		}

		var rows []uint32
		for _, line := range rowBytes {

			if len(line) < 32 {
				line = append([]byte(strings.Repeat("0", 32-len(line))), line...)
			}

			num, err := strconv.ParseUint(string(line), 2, 32)
			if err != nil {
				panic(err)
			}

			rows = append(rows, uint32(num))
		}

		// fmt.Printf("\n%v\n", strings.Join(rows, "\n"))

		for i, num := range rows[:len(rows)-1] {
			isOneOff := bits.OnesCount32(num^rows[i+1]) == 1
			// fmt.Printf("col: %v, %v:%v\n", i, num, rows[i+1])
			if num == rows[i+1] || isOneOff {
				j := i - 1
				k := i + 2
				for k < len(rows) && j >= 0 {
					oneOff := bits.OnesCount32(rows[j]^rows[k]) == 1
					if oneOff && !isOneOff {
						isOneOff = true
						k++
						j--
						continue
					}

					if rows[j] != rows[k] {
						break
					}

					k++
					j--
				}

				if (k == len(rows) || j == -1) && isOneOff {
					sum += i + 1
					found = true
					fixed = true
					break
				}

			}
		}

		if !found {
			fmt.Printf("did not find: %v\n", blockNum)
		}
		if !fixed {
			fmt.Printf("did not fix: %v\n", blockNum)
		}
	}

	fmt.Printf("sum: %v\n", sum)

}
