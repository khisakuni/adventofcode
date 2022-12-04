package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(input), "\n")
	var count int
	for _, str := range strs {
		if str == "" {
			continue
		}

		elves := strings.Split(str, ",")
		elf1 := strings.Split(elves[0], "-")
		elf1Min, _ := strconv.Atoi(elf1[0])
		elf1Max, _ := strconv.Atoi(elf1[1])

		elf2 := strings.Split(elves[1], "-")
		elf2Min, _ := strconv.Atoi(elf2[0])
		elf2Max, _ := strconv.Atoi(elf2[1])

		hasOverlap := (elf1Max <= elf2Max && elf1Max >= elf2Min) ||
			(elf1Min >= elf2Min && elf1Min <= elf2Max) ||
			(elf2Max <= elf1Max && elf2Max >= elf1Min) ||
			(elf2Min >= elf1Min && elf2Min <= elf1Max)
		if hasOverlap {
			count++
		}
	}

	fmt.Printf("Count: %v\n", count)
}
