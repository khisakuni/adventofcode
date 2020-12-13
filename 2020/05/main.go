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
	rowCount := 128
	colCount := 8
	m := make([][]int, rowCount)
	for i := range m {
		m[i] = make([]int, colCount)
	}

	strs := strings.Split(string(data), "\n")
	//var max float64
	for _, str := range strs {
		if str == "" {
			continue
		}
		row, col, id := parse(str)
		m[int(row)][int(col)] = int(id)
	}
	for j, r := range m {
		hasVal := false
		hasEmpty := false
		found := false
		for i, c := range r {
			if c > 0 && !hasVal {
				hasVal = true
				continue
			}
			if hasVal && c == 0 && !hasEmpty {
				hasEmpty = true
				continue
			}
			if hasEmpty && hasVal && c > 0 && !found {
				found = true
				fmt.Printf("found it: %v, %v, %v\n", j, i - 1, j * 8 + (i - 1))
				continue
			}
		}
	}
	//fmt.Printf("max: %v\n", m)
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
