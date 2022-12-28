package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	//	input = []byte(`addx 15
	//addx -11
	//addx 6
	//addx -3
	//addx 5
	//addx -1
	//addx -8
	//addx 13
	//addx 4
	//noop
	//addx -1
	//addx 5
	//addx -1
	//addx 5
	//addx -1
	//addx 5
	//addx -1
	//addx 5
	//addx -1
	//addx -35
	//addx 1
	//addx 24
	//addx -19
	//addx 1
	//addx 16
	//addx -11
	//noop
	//noop
	//addx 21
	//addx -15
	//noop
	//noop
	//addx -3
	//addx 9
	//addx 1
	//addx -3
	//addx 8
	//addx 1
	//addx 5
	//noop
	//noop
	//noop
	//noop
	//noop
	//addx -36
	//noop
	//addx 1
	//addx 7
	//noop
	//noop
	//noop
	//addx 2
	//addx 6
	//noop
	//noop
	//noop
	//noop
	//noop
	//addx 1
	//noop
	//noop
	//addx 7
	//addx 1
	//noop
	//addx -13
	//addx 13
	//addx 7
	//noop
	//addx 1
	//addx -33
	//noop
	//noop
	//noop
	//addx 2
	//noop
	//noop
	//noop
	//addx 8
	//noop
	//addx -1
	//addx 2
	//addx 1
	//noop
	//addx 17
	//addx -9
	//addx 1
	//addx 1
	//addx -3
	//addx 11
	//noop
	//noop
	//addx 1
	//noop
	//addx 1
	//noop
	//noop
	//addx -13
	//addx -19
	//addx 1
	//addx 3
	//addx 26
	//addx -30
	//addx 12
	//addx -1
	//addx 3
	//addx 1
	//noop
	//noop
	//noop
	//addx -9
	//addx 18
	//addx 1
	//addx 2
	//noop
	//noop
	//addx 9
	//noop
	//noop
	//noop
	//addx -1
	//addx 2
	//addx -37
	//addx 1
	//addx 3
	//noop
	//addx 15
	//addx -21
	//addx 22
	//addx -6
	//addx 1
	//noop
	//addx 2
	//addx 1
	//noop
	//addx -10
	//noop
	//noop
	//addx 20
	//addx 1
	//addx 2
	//addx 2
	//addx -6
	//addx -11
	//noop
	//noop
	//noop
	//`)

	inputs := strings.Split(string(input), "\n")
	x := 1
	var work int
	var add int
	var inputIndex int
	var strength int
	for cycle := 1; cycle <= 220; cycle++ {
		//fmt.Printf("cycle: %d, work: %v, x: %v\n", cycle, work, x)
		// Do the work
		skip := work > 0
		if work > 0 {
			work--
		} else {
			x += add
		}

		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			fmt.Printf("cycle: %v, x: %v\n", cycle, x)
			strength += cycle * x
		}

		if skip {
			continue
		}

		if inputIndex >= len(inputs) {
			break
		}

		in := inputs[inputIndex]
		inputIndex++
		var op string
		var num int
		_, _ = fmt.Sscanf(in, "%s %d", &op, &num)
		add = num
		switch op {
		case "noop":
			work = 0
		case "addx":
			work = 1
		}
	}

	fmt.Printf("strength: %v\n", strength)
}
