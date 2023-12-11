package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	rows := make([]int, len(lines))
	cols := make([]int, len(lines[0]))
	var coords [][]int

	var count int
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				rows[i]++
				cols[j]++
				count++
				coords = append(coords, []int{i, j, count})
			}
		}
	}

	var pairs [][][]int

	for i, coord := range coords {
		for _, other := range coords[i+1:] {
			pairs = append(pairs, [][]int{coord, other})
		}
	}

	// fmt.Printf("num pairs: %v\n", len(pairs))
	var sum int
	for _, pair := range pairs {
		coordA := pair[0]
		coordB := pair[1]

		ay := coordA[0]
		ax := coordA[1]
		// a := coordA[2]

		by := coordB[0]
		bx := coordB[1]
		// b := coordB[2]

		width := int(math.Abs(float64(ax - bx)))
		height := int(math.Abs(float64(ay - by)))

		if ay > by {
			for _, r := range rows[by:ay] {
				if r == 0 {
					height += 999999
				}
			}
		} else {
			for _, r := range rows[ay:by] {
				if r == 0 {
					height += 999999
				}
			}
		}

		if ax > bx {
			for _, c := range cols[bx:ax] {
				if c == 0 {
					width += 999999
				}
			}

		} else {
			for _, c := range cols[ax:bx] {
				if c == 0 {
					width += 999999
				}
			}
		}

		diag := width + height

		// fmt.Printf("[%d, %d] width: %d, height: %d >> %v\n", a, b, width, height, int(diag))

		sum += diag
	}

	fmt.Printf("sum: %v\n", sum)

}

func gcd(a, b int) int {
	if a < b {
		tmp := b
		b = a
		a = tmp
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}

	return b
}
