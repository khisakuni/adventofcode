package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// var input = `125 17`

func main() {
	part1()
	part2()
}

func part1() {
	t := time.Now()
	defer func() {
		fmt.Printf("%v\n", time.Since(t))
	}()
	stones := strings.Split(input, " ")

	for i := 0; i < 25; i++ {
		var nextStones []string
		for _, stone := range stones {
			switch {
			case stone == "0":
				nextStones = append(nextStones, "1")
			case len(stone)%2 == 0:
				left := strings.TrimLeft(stone[:len(stone)/2], "0")
				if left == "" {
					left = "0"
				}
				right := strings.TrimLeft(stone[len(stone)/2:], "0")
				if right == "" {
					right = "0"
				}
				nextStones = append(nextStones, left, right)
			default:
				num, _ := strconv.Atoi(stone)
				num *= 2024
				nextStones = append(nextStones, strconv.Itoa(num))
			}

		}

		stones = nextStones
	}

	fmt.Printf("part 1: %v\n", len(stones))
}

func part2() {
	t := time.Now()
	defer func() {
		fmt.Printf("%v\n", time.Since(t))
	}()
	stonesStrs := strings.Split(input, " ")

	m := map[int]int{}
	for _, s := range stonesStrs {
		num, _ := strconv.Atoi(s)
		m[num]++
	}

	for i := 0; i < 75; i++ {
		next := map[int]int{}
		for stone, count := range m {
			if stone == 0 {
				next[1] += count
				continue
			}

			var n int
			s := stone
			for s > 0 {
				s /= 10
				n++
			}

			if n%2 == 0 {
				pow := int(math.Pow10(n / 2))
				left := stone / pow
				right := stone % pow
				next[left] += count
				next[right] += count
			} else {
				next[stone*2024] += count

			}
		}

		m = next
	}

	var total int
	for _, v := range m {
		total += v
	}

	fmt.Printf("part 2: %v\n", total)
}
