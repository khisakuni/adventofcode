package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n")
	numbers := []int{}
	for _, n := range strs {
		if n == "" {
			continue
		}
		n, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	m := map[int]bool{}
	for _, n := range numbers {
		m[n] = true
	}


	var n1, n2, n3 int
	for k := range m {
		for l := range m {
			if k == l {
				continue
			}
			if _, ok := m[2020 - k - l]; ok {
				n1 = k
				n2 = l
				n3 = 2020 - k - l
				break
			}
		}
	}


	fmt.Printf("Result: %d * %d * %d = %d\n", n1, n2, n3, n1 * n2 * n3)
}
