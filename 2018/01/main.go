package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/khisakuni/adventofcode/2018/common"
)

func main() {
	fmt.Println("Running day 1")

	input, err := common.Input(filepath.Join("input.txt"))
	if err != nil {
		common.HandleError(err)
	}
	numbers := make([]int, len(input), len(input))
	for i, str := range input {
		num, err := strconv.Atoi(str)
		if err != nil {
			common.HandleError(err)
		}
		numbers[i] = num
	}

	fmt.Println("Part 1")
	fmt.Printf("answer is: %d\n", reduce(numbers, 0))

	fmt.Println("Part 2")
	fmt.Printf("answer is: %d\n", repeat(numbers))

	os.Exit(0)
}

func repeat(numbers []int) int {
	var r int
	var err error
	cache := map[int]bool{r: true}
	r, err = findRepeat(numbers, r, cache)
	for err != nil {
		r, err = findRepeat(numbers, r, cache)
	}
	return r
}

func findRepeat(numbers []int, agg int, cache map[int]bool) (int, error) {
	for _, n := range numbers {
		agg += n
		if v := cache[agg]; v {
			return agg, nil
		}
		cache[agg] = true
	}
	return agg, errors.New("no repeat found")
}

func reduce(ar []int, agg int) int {
	for _, n := range ar {
		agg += n
	}
	return agg
}
