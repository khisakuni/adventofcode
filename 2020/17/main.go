package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cell struct {
	active bool
}

type coords struct {
	x int
	y int
	z int // frame
}

type world []*frame
type frame []*row
type row []*cell

func (w world) printState() {
	for z, f := range w {
		fmt.Printf("z=%d\n", z)
		for _, slice := range *f {
			for _, c := range *slice {
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

func (w world) cycle() world {
	newWorld := w
	for z, f := range w {
		for y, r := range *f {
			for x := range *r {
				coord := coords{x: x, y: y, z: z}
				w.cellIsActive(coord)
			}
		}
	}
	/* Iterate through each cell and add new state to newWorld*/
	return newWorld
}

func (w world) cellIsActive(coord coords) bool {
	neighbors := w.neighbors(coord)
	fmt.Printf("NEIGHBOR COUNT: %v\n", len(neighbors))
	return false
}

func (w world) neighbors(coord coords) []*cell {
	zs := []int{coord.z-1, coord.z, coord.z+1}
	ys := []int{coord.y-1, coord.y, coord.y+1}
	xs := []int{coord.x-1, coord.x, coord.x+1}
	var cells []*cell
	for _, z := range zs {
		for _, y := range ys {
			for _, x := range xs {
				if x == coord.x && y == coord.y && z == coord.z {
					continue
				}
				f := w.getFrame(z)
				r := w.getRowForFrame(f, y)
				c := w.getCellForRow(r, x)
				cells = append(cells, c)
			}
		}
	}
	return cells
}

func (w world) depth() int {
	return len(w)
}

func (w world) height() int {
	return w[0].height()
}

func (w world) width() int {
	return (*(w)[0])[0].width()
}

func (w *world) getFrame(z int) *frame {
	if z == -1 {
		f := w.newFrame()
		w.pushFrame(f)
		return f
	}
	if z >= len(*w) {
		f := w.newFrame()
		w.appendFrame(f)
		return f
	}
	return (*w)[z]
}

func (w *world) pushFrame(f *frame) {
	*w = append(world{f}, *w...)
}

func (w *world) appendFrame(f *frame) {
	*w = append(*w, f)
}

func (w world) newFrame() *frame {
	f := make(frame, w.height())
	width := w.width()
	for i := range f {
		r := make(row, width)
		f[i] = &r
	}
	return &f
}

func (w *world) addCell(a cell, coord coords) {
	f := w.getFrame(coord.z)
	r := w.getRowForFrame(f, coord.y)
	c := w.getCellForRow(r, coord.x)
	c.active = a.active
}

func (w *world) getRowForFrame(f *frame, y int) *row {
	if y == -1 {
		for _, f := range *w {
			f.pushRow(f.newRow())
		}
		return (*f)[0]
	}
	if y >= f.height() {
		for _, f := range *w {
			f.appendRow(f.newRow())
		}
	}
	return (*f)[y]
}

func (w *world) getCellForRow(r *row, x int) *cell {
	if x == -1 {
		for _, f := range *w {
			for _, r := range *f {
				r.pushCell(&cell{})
			}
		}
		return (*r)[0]
	}
	if x >= r.width() {
		for _, f := range *w {
			for _, r := range *f {
				//fmt.Printf("before: %v ", len(r))
				r.pushCell(&cell{})
				//fmt.Printf("after: %v\n", len(r))
			}
		}
	}
	//fmt.Printf("x: %v, len: %v\n", x, r.width())
	return (*r)[x]
}

func (f *frame) pushRow(r *row) {
	*f = append(frame{r}, *f...)
}

func (f *frame) appendRow(r *row) {
	*f = append(*f, r)
}

func (f frame) newRow() *row {
	r := make(row, f.width())
	return &r
}

func (f frame) height() int {
	return len(f)
}

func (f frame) width() int {
	return f[0].width()
}

func (r *row) pushCell(c *cell) {
	*r = append(row{c}, *r...)
}

func (r *row) appendCell(c *cell) {
	*r = append(*r, c)
}

func (r row) width() int {
	return len(r)
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	w := world{}
	w = append(w, &frame{})
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
			slice = append(slice, &c)
		}
		f := w[0]
		f.appendRow(&slice)
		//w[0] = append(f, slice)
	}
	w.printState()
	w.cycle()
}

