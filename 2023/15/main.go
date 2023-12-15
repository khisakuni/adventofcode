package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("input.txt")
	inputs := bytes.Split(data, []byte{','})

	boxes := make([][]string, 256)
	for _, input := range inputs {
		if len(input) == 0 {
			continue
		}
		if input[len(input)-1] == '\n' {
			input = input[:len(input)-1]
		}

		lensNameIndex := bytes.IndexAny(input, "-=")
		lensName := input[:lensNameIndex]

		box := hash(lensName)
		if bytes.Contains(input, []byte{'='}) {
			i := slices.IndexFunc(boxes[box], func(n string) bool {
				if len(n) < len(lensName) {
					return false
				}

				if n[:lensNameIndex] == string(lensName) {
					return true
				}

				return false
			})
			if i >= 0 {
				boxes[box][i] = fmt.Sprintf("%s %s", lensName, input[lensNameIndex+1:])
			} else {
				boxes[box] = append(boxes[box], fmt.Sprintf("%s %s", lensName, input[lensNameIndex+1:]))
			}
		} else {
			i := slices.IndexFunc(boxes[box], func(n string) bool {
				if len(n) < len(lensName) {
					return false
				}

				if n[:lensNameIndex] == string(lensName) {
					return true
				}

				return false
			})

			if i >= 0 {
				boxes[box] = append(boxes[box][0:i], boxes[box][i+1:]...)
			}
		}

		// fmt.Printf("After \"%s\":\n", string(input))
		// for i := range boxes[:4] {
		// 	fmt.Printf("Box %d:", i)
		// 	for _, l := range boxes[i] {
		// 		fmt.Printf("[%s] ", l)
		// 	}
		// 	fmt.Println()
		// }
	}

	var sum int
	for i, slots := range boxes {
		boxNum := i + 1
		for j, slot := range slots {
			slotNum := j + 1
			parts := strings.Split(slot, " ")
			focalLen, _ := strconv.Atoi(parts[1])

			// fmt.Printf("%v: %v (box %v) * %v (slot) * %v (focal length) %v\n", parts[0], i+1, i, j+1, focalLen, boxNum*slotNum*focalLen)
			sum += boxNum * slotNum * focalLen
		}

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
