package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cell struct {
	active bool
}

func (c cell) print() {
	if c.active {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

type coords struct {
	x int
	y int
	z int // frame
}

type world []frame
type frame []row
type row []cell

func (w world) printState() {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	for z, f := range w {
		fmt.Printf("z=%d\n", z)
		for _, slice := range f {
			for _, c := range slice {
				c.print()
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}

func (w world) activeCount() int {
	return 0
}

func (w world) cycle() world {
	next := w.grow()
	for z := 1; z < w.depth(); z++ {
		for y := 1; y < w.height(); y++ {
			for x := 1; x < w.width(); x++ {
				//next[z][y][x]
				//c.print()
			}
			//fmt.Println()
		}
		//fmt.Println()
	}
	return next
}

/*
         y      z
         |     /
         |    /
         |   /
         |  /
         | /
_________|/__________ x
				 /
				/|
       / |
      /  |
     /   |

*/


func (w world) grow() world {
	next := w
	for z := range next {
		for y := range next[z] {
			r := next[z][y]
			r = append(row{cell{}}, r...)
			r = append(r, cell{})
			next[z][y] = r
		}
		f := next[z]
		f = append(
			frame{make(row, next.width())},
			f...
		)
		f = append(
			f,
			make(row, w.width()),
		)
		next[z] = f
	}

	// Make -1 and +1 frames
	f := make(frame, next.height())
	for y := range f {
		f[y] = make(row, next.width())
	}

	after := make(frame, next.height())
	for y := range after {
		after[y] = make(row, next.width())
	}

	// Add -1 and +1 rows to each existing frame
	next = append([]frame{f}, next...)
	next = append(next, after)
	return next
}

func (w world) depth() int {
	return len(w)
}

func (w world) height() int {
	return len(w[0])
}

func (w world) width() int {
	return len(w[0][0])
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	w := world{}
	w = append(w, frame{})
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		slice := row{}
		for _, char := range line {
			c := cell{}
			if char == '#' {
				c.active = true
			}
			slice = append(slice, c)
		}
		f := w[0]
		w[0] = append(f, slice)
	}
	w.printState()

	fmt.Println(">>>>>>>>>>>>>>>>>>")
	w = w.cycle().cycle().cycle()

	w.printState()
}

