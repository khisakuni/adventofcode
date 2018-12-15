package main

import (
	"fmt"
	"path/filepath"

	"github.com/khisakuni/adventofcode/2018/common"
)

func main() {
	fmt.Println("Running day 2")

	input, err := common.Input(filepath.Join("input.txt"))
	if err != nil {
		common.HandleError(err)
	}
	// input := []string{
	// 	"abcdef",
	// 	"bababc",
	// 	"abbcde",
	// 	"abcccd",
	// 	"aabcdd",
	// 	"abcdee",
	// 	"ababab",
	// }

	two := 0
	three := 0
	for _, str := range input {
		if hasRep(str, 3) {
			three++
		}
		if hasRep(str, 2) {
			two++
		}
	}

	fmt.Printf("Answer is %d\n", two*three)
}

func hasRep(str string, num int) bool {
	if len(str) == 0 {
		return false
	}
	m := map[rune]int{}
	for _, c := range str {
		m[c]++
	}
	for _, v := range m {
		if v == num {
			return true
		}
	}
	return false
}
