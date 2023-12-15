package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	inputs := bytes.Split(data, []byte{','})
	var sum int
	for _, input := range inputs {
		if len(input) == 0 {
			continue
		}

		sum += hash(input)
	}

	fmt.Printf("sum: %v\n", sum)
}

func hash(input []byte) int {
	var sum int
	for _, c := range input {
		sum += int(c)
		sum *= 17
		sum = sum % 256
	}

	return sum
}
