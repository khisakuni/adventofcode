package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cell struct {
	active bool
}

type world [][][]cell

func (w world) printState() {
	for z, frame := range w {
		fmt.Printf("z=%d\n", z)
		for _, slice := range frame {
			for _, c := range slice {
				if c.active {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	w := world{}
	w = append(w, [][]cell{})
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		slice := []cell{}
		for _, char := range line {
			c := cell{}
			if char == '#' {
				c.active = true
			}
			slice = append(slice, c)
		}
		w[0] = append(w[0], slice)
	}
	w.printState()
}
