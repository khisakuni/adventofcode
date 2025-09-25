package main

import (
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input.txt
var input string

// var input = `2333133121414131402`

func main() {
	part1()
	part2()
}

func part1() {
	seq := []int{}
	var fileID int
	for i, c := range input {
		num := int(c - '0')
		isFile := (i % 2) == 0

		if isFile {
			for i := 0; i < num; i++ {
				seq = append(seq, fileID)
			}

			fileID++
		} else {
			for i := 0; i < num; i++ {
				seq = append(seq, -1)
			}
		}
	}

	for i := len(seq) - 1; i >= 0; i-- {
		empty := slices.Index(seq, -1)
		if empty >= i {
			break
		}

		seq[empty] = seq[i]
		seq[i] = -1
	}

	var checksum int
	for i, c := range seq {
		if c == -1 {
			break
		}

		checksum += i * c
	}

	fmt.Printf("part1: %v\n", checksum)
}

func part2() {
	seq := []int{}
	var fileID int
	for i, c := range input {
		num := int(c - '0')
		isFile := (i % 2) == 0

		if isFile {
			for i := 0; i < num; i++ {
				seq = append(seq, fileID)
			}

			fileID++
		} else {
			for i := 0; i < num; i++ {
				seq = append(seq, -1)
			}
		}
	}

	right := len(seq) - 1
	for right > 0 {
		for seq[right] == -1 {
			right--
		}

		end := right
		start := end
		current := seq[start]
		for start >= 0 && seq[start] == current {
			start--
		}

		l := end - start

		emptyStart := 0
		emptyEnd := 0
		for emptyEnd < end {
			for emptyStart < end && seq[emptyStart] > -1 {
				emptyStart++
			}

			emptyEnd = emptyStart
			for emptyEnd < end && seq[emptyEnd] == -1 {
				emptyEnd++
			}

			if emptyEnd-emptyStart >= l {
				for i := emptyStart; i < emptyStart+l; i++ {
					seq[i] = current
				}

				for i := start + 1; i <= end; i++ {
					seq[i] = -1
				}
				break
			}

			emptyStart = emptyEnd + 1
		}

		right = start
	}

	var checksum int
	for i, c := range seq {
		if c == -1 {
			continue
		}

		checksum += i * c
	}

	fmt.Printf("part 2: %v\n", checksum)

}
