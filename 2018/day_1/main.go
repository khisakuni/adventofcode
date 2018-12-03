package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	fmt.Println("Running day 1")

	f, err := os.Open(filepath.Join("input.txt"))
	if err != nil {
		handleError(err)
		return
	}
	defer f.Close()

	numbers := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			handleError(err)
		}
		numbers = append(numbers, i)
	}
	if err := scanner.Err(); err != nil {
		handleError(err)
		return
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

func handleError(err error) {
	fmt.Printf("encountered error: %v\n", err)
	os.Exit(1)
}
