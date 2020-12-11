package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n")
	slope := [][]rune{}
	for _, line := range strs {
		if line == "" {
			continue
		}
		slope = append(slope, []rune(line))
	}
	right := []int{1, 3, 5, 7, 1}
	down := []int{1, 1, 1, 1, 2}
	var trees []int
	total := 1
	for i, r := range right {
		count := doRun(slope, r, down[i])
		total = total * count
		trees = append(trees, count)
	}
	fmt.Printf("Trees: %v, total: %d\n", trees, total)
}

func doRun(slope [][]rune, right int, down int) int {
	height := len(slope)
	x := 0
	width := len(slope[0])
	var trees int
	for y := down; y < height; y += down {
		x = (x + right) % width
		pos := slope[y][x]
		if pos == '#' {
			trees++
		}
	}
	return trees
}

/*

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#

*/
