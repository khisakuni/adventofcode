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

func (c cell) printColor() {
	if c.active {
		fmt.Print("\x1b[6;30;42m#\x1b[0m")
	} else {
		fmt.Print("\x1b[6;30;42m.\x1b[0m")
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

func (w world) printStateForCell(coord coords) {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	for z, f := range w {
		fmt.Printf("z=%d\n", z)
		for y, slice := range f {
			for x, c := range slice {
				if z==coord.z && y == coord.y && x ==coord.x {
					c.printColor()
					continue
				}
				c.print()
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}


func (w world) cycle() world {
	current := w.grow()
	next := current.copy()
	for z := 0; z < current.depth(); z++ {
		for y := 0; y < current.height(); y++ {
			for x := 0; x < current.width(); x++ {
				active := current.isActive(coords{z:z, y:y, x:x})
				next[z][y][x] = cell{active: active}
			}
		}
	}
	return next
}

func (w world) isActive(coord coords) bool {
	neighbors := w.neighbors(coord)
	count := activeCount(neighbors)
	c := w[coord.z][coord.y][coord.x]
	var active bool
	if c.active {
		active = count == 2 || count == 3
	} else {
		active = count == 3
	}
	return active
}

func (w world) copy() world {
	dupe := make(world, w.depth())
	for z, f := range w {
		dupe[z] = make(frame, w.height())
		for y, r := range f {
			dupe[z][y] = make(row, w.width())
			for x, c := range r {
				dupe[z][y][x] = cell{active: c.active}
			}
		}
	}
	return dupe
}

func (w world) activeCount() int {
	var count int
	for _, f := range w {
		for _, r := range f {
			for _, c := range r {
				if c.active {
					count++
				}
			}
		}
	}
	return count
}

func activeCount(cells []cell) int {
	var count int
	for _, c := range cells {
		if c.active {
			count++
		}
	}
	return count
}

func (w world) neighbors(coord coords) []cell {
	zs := []int{coord.z}
	if coord.z > 0 {
		zs = append(zs, coord.z-1)
	}
	if coord.z < w.depth()-1 {
		zs = append(zs, coord.z+1)
	}
	ys := []int{coord.y}
	if coord.y > 0 {
		ys = append(ys, coord.y-1)
	}
	if coord.y < w.height()-1 {
		ys = append(ys, coord.y+1)
	}
	xs := []int{coord.x}
	if coord.x > 0 {
		xs = append(xs, coord.x-1)
	}
	if coord.x < w.width()-1 {
		xs = append(xs, coord.x+1)
	}

	var neighbors []cell
	for _, z := range zs {
		for _, y := range ys {
			for _, x := range xs {
				if z == coord.z && y == coord.y && x == coord.x {
					continue
				}
				neighbors = append(neighbors, w[z][y][x])
			}
		}
	}
	return neighbors
}

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
	w = w.cycle().cycle().cycle().cycle().cycle().cycle()
	fmt.Printf("active: %d\n", w.activeCount())
}

