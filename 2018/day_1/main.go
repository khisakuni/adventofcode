package main

import (
	"bufio"
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

	fmt.Printf("answer is: %d\n", reduce(numbers, 0))

	os.Exit(0)
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
