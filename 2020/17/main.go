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
	w int
}

type universe []world
type world []frame
type frame []row
type row []cell

func (w world) printState() {
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
}

func (u universe) printState() {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	for i, w := range u {
		for z, f := range w {
			fmt.Printf("z=%d, w=%d\n", z, i)
			for _, s := range f {
				for _, c := range s {
					c.print()
				}
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}

func (u universe) cycle() universe {
	current := u.grow()
	next := current.copy()
	for i := 0; i < len(current); i++ {
		for z := 0; z < current[i].depth(); z++ {
			for y := 0; y < current[i].height(); y++ {
				for x := 0; x < current[i].width(); x++ {
					active := current.isActive(coords{z:z, y:y, x:x, w:i})
					next[i][z][y][x] = cell{active: active}
				}
			}
		}
	}
	return next
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

func (u universe) isActive(coord coords) bool {
	neighbors := u.neighbors(coord)
	count := activeCount(neighbors)
	c := u[coord.w][coord.z][coord.y][coord.x]
	var active bool
	if c.active {
		active = count == 2 || count == 3
	} else {
		active = count == 3
	}
	return active
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

func (u universe) copy() universe {
	dupe := make(universe, len(u))
	for i, w := range u {
		dupe[i] = w.copy()
	}
	return dupe
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

func (u universe) activeCount() int {
	var count int
	for _, w := range u {
		for _, f := range w {
			for _, r := range f {
				for _, c := range r {
					if c.active {
						count++
					}
				}
			}
		}
	}
	return count
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

func (u universe) neighbors(coord coords) []cell {
	ws := []int{coord.w}
	if coord.w > 0 {
		ws = append(ws, coord.w-1)
	}
	if coord.w < len(u)-1 {
		ws = append(ws, coord.w+1)
	}
	zs := []int{coord.z}
	if coord.z > 0 {
		zs = append(zs, coord.z-1)
	}
	if coord.z < u[0].depth()-1 {
		zs = append(zs, coord.z+1)
	}
	ys := []int{coord.y}
	if coord.y > 0 {
		ys = append(ys, coord.y-1)
	}
	if coord.y < u[0].height()-1 {
		ys = append(ys, coord.y+1)
	}
	xs := []int{coord.x}
	if coord.x > 0 {
		xs = append(xs, coord.x-1)
	}
	if coord.x < u[0].width()-1 {
		xs = append(xs, coord.x+1)
	}

	var neighbors []cell
	for _, w := range ws {
		for _, z := range zs {
			for _, y := range ys {
				for _, x := range xs {
					if z == coord.z && y == coord.y && x == coord.x && w == coord.w {
						continue
					}
					neighbors = append(neighbors, u[w][z][y][x])
				}
			}
		}
	}
	return neighbors
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

func (u universe) grow() universe {
	next := u
	for i := range next {
		next[i] = next[i].grow()
	}

	depth := next[0].depth()
	height := next[0].height()
	width := next[0].width()
	before := make(world, depth)
	for z := range before {
		frames := make(frame, height)
		for y := range frames {
			frames[y] = make(row, width)
		}
		before[z] = frames
	}

	after := make(world, depth)
	for z := range after {
		frames := make(frame, height)
		for y := range frames {
			frames[y] = make(row, width)
		}
		after[z] = frames
	}

	next = append(universe{before}, next...)
	next = append(next, after)

	return next
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
	u := universe{}
	u = append(u, world{frame{}})
	w := u[0]
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
	u = u.cycle().cycle().cycle().cycle().cycle().cycle()
	//u.printState()
	fmt.Printf("active: %d\n", u.activeCount())
}

