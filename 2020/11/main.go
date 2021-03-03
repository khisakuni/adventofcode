package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type seatMap map[string]seat

type seat rune

const (
	seatStateFloor rune = '.'
	seatStateOccupied = '#'
	seatStateVacant = 'L'
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	m := seatMap{}
	for i, row := range strings.Split(string(data), "\n")  {
		if row == "" {
			continue
		}
		for j , col := range row {
			m[seatKey(i, j)] = seat(col)
		}
	}

	changed := true
	for changed {
		m, changed = cycle(m)
	}
	fmt.Printf("count: %d\n", countOccupied(m))
}

func print(m seatMap) {
	width, height := 10, 10
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			key := seatKey(i, j)
			fmt.Print(string(m[key]))
		}
		fmt.Println()
	}
}

func cycle(m seatMap) (seatMap, bool) {
	var changed bool
	next := seatMap{}
	for k, v := range m {
		var occupied int
		for _, neighbor := range neighborKeys(k) {
			if n, ok := m[neighbor]; ok {
				if n == seatStateOccupied {
					occupied++
				}
			}
		}
		nextState := v
		if v == seatStateVacant && occupied == 0 {
			nextState = seatStateOccupied
		}
		if v == seatStateOccupied && occupied >= 4 {
			nextState = seatStateVacant
		}
		next[k] = nextState
		if nextState != v && !changed {
			changed = true
		}
	}
	return next, changed
}

func countOccupied(m seatMap) int {
	var count int
	for _, v := range m {
		if v == seatStateOccupied {
			count++
		}
	}
	return count
}

func neighborKeys(key string) []string {
	row, col := parseKey(key)
	return []string{
		// Top
		seatKey(row-1, col-1),
		seatKey(row-1, col),
		seatKey(row-1, col+1),

		// Mid
		seatKey(row, col-1),
		seatKey(row, col+1),

		// Bottom
		seatKey(row+1, col-1),
		seatKey(row+1, col),
		seatKey(row+1, col+1),
	}
}

func seatKey(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
}

func parseKey(key string) (int, int) {
	vals := strings.Split(key, ":")
	if len(vals) != 2 {
		panic("malformed key: " + key)
	}
	row, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	col, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	return row, col
}

