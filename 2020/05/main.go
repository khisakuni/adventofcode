package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n")
	var max float64
	for _, str := range strs {
		if str == "" {
			continue
		}
		row, col, id := parse(str)
		fmt.Printf("%s\n", str)
		fmt.Printf("  row %f, column %f, seat ID %f.\n", row, col, id)
		if id > max {
			max = id
		}
	}
	fmt.Printf("max: %v\n", max)
}

func parse(str string) (row, col, id float64) {
	var minRow float64 = 0
	var maxRow float64 = 127
	var minCol float64 = 0
	var maxCol float64 = 7
	for i, c := range str {
		if c == 'F' {
			if i == 6 {
				row = minRow
			}
			maxRow = maxRow - math.Round((maxRow - minRow) / 2)
		}
		if c == 'B' {
			if i == 6 {
				row = maxRow
			}
			minRow = minRow + math.Round((maxRow - minRow) / 2)
		}
		if c == 'L' {
			if i == 9 {
				col = minCol
			}
			maxCol = maxCol - math.Round((maxCol - minCol) / 2)
		}
		if c == 'R' {
			if i == 9 {
				col = maxCol
			}
			minCol = minCol + math.Round((maxCol - minCol) / 2)
		}
	}
	id = row * 8 + col
	return row, col, id
}
