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
