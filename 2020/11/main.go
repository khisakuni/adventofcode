package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type seatMap map[string]seat

type seat rune

const (
	seatStateFloor seat = '.'
	seatStateOccupied = '#'
	seatStateVacant = 'L'
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	m := readMap(data)

	// part 1
	part1Start := time.Now()
	changed := true
	for changed {
		m, changed = cycle(m)
	}
	fmt.Printf("part 1 took: %v", time.Since(part1Start))
	fmt.Printf("count: %d\n", countOccupied(m))

	// part 2
	m2 := readMap(data)
	part2Start := time.Now()
	changed = true
	for changed {
		m2, changed = cyclePart2(m2)
	}
	fmt.Printf("part 2 took %v", time.Since(part2Start))


	fmt.Printf("count: %d\n", countOccupied(m2))
}

func readMap(data []byte) seatMap {
	m := seatMap{}
	for i, row := range strings.Split(string(data), "\n")  {
		if row == "" {
			continue
		}
		for j , col := range row {
			m[seatKey(i, j)] = seat(col)
		}
	}
	return m
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

func cyclePart2(m seatMap) (seatMap, bool) {
	var changed bool
	next := seatMap{}
	for k, v := range m {
		nextState := v
		count := occupiedNeighborCount(m, k)
		if v == seatStateOccupied && count >= 5 {
			nextState = seatStateVacant
		}
		if v == seatStateVacant && count == 0 {
			nextState = seatStateOccupied
		}
		if !changed && v != nextState {
			changed = true
		}
		next[k] = nextState
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

func occupiedNeighborCount(m seatMap, key string) int {
	var count int
	neighbors := []bool{
		isNeighborOccupied(m, key, -1, -1), // top left
		isNeighborOccupied(m, key, -1, 0), // top
		isNeighborOccupied(m, key, -1, 1), // top right

		isNeighborOccupied(m, key, 0, -1), // left
		isNeighborOccupied(m, key, 0, 1), // right

		isNeighborOccupied(m, key, 1, -1), // bottom left
		isNeighborOccupied(m, key, 1, 0), // bottom
		isNeighborOccupied(m, key, 1, 1), // bottom right
	}
	for _, neighborIsOccupied := range neighbors {
		if neighborIsOccupied {
			count++
		}
	}
	return count
}

func isNeighborOccupied(m seatMap, current string, rowOffset, colOffset int) bool {
	row, col := parseKey(current)
	neighborRow := row+rowOffset
	neighborCol := col+colOffset
	neighborKey := seatKey(neighborRow, neighborCol)
	neighborSeat, ok := m[neighborKey]
	if !ok {
		return false
	}
	if neighborSeat == seatStateFloor {
		return isNeighborOccupied(m, neighborKey, rowOffset, colOffset)
	}
	return neighborSeat == seatStateOccupied
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

