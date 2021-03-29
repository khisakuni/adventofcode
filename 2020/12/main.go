package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type direction rune

const (
	directionNorth = 'N'
	directionEast  = 'E'
	directionSouth = 'S'
	directionWest  = 'W'

	directionRight   = 'R'
	directionLeft    = 'L'
	directionForward = 'F'
)

type pos struct {
	// N/S
	lat int

	// E/W
	long int
}

func (p pos) currentLat() direction {
	if p.lat > 0 {
		return directionNorth
	}
	return directionSouth
}

func (p pos) currentLong() direction {
	if p.long > 0 {
		return directionEast
	}
	return directionWest
}

var compass = []direction{directionEast, directionSouth, directionWest, directionNorth}

func getDirectionIndex(d direction) int {
	for i, dir := range compass {
		if dir == d {
			return i
		}
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	directionIndex := 0
	p := pos{}
	//waypoint := pos{lat: 1, long: 10}
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		fmt.Println(line)
		d, units := parseLine(line)
		directionIndex = part1(&p, directionIndex, d, units)
	}

	fmt.Printf("pos >> %+v\n", p)
	fmt.Printf("manhattan distance: %v\n", manhattanValue(p))

	// -----------
	p2 := pos{}
	waypoint := pos{lat: 1, long: 10}
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		d, units := parseLine(line)
		part2(&waypoint, &p2, d, units)
	}

	fmt.Printf("manhattan distance p2: %v\n", manhattanValue(p2))
}

func part2(waypoint *pos, p *pos, d direction, units int) {
	switch d {
	case directionNorth, directionEast, directionSouth, directionWest:
		advance(waypoint, d, units)
	case directionForward:
		p.lat = p.lat + (waypoint.lat * units)
		p.long = p.long + (waypoint.long * units)
	case directionRight, directionLeft:
		var turnUnits int
		var latDir direction
		var longDir direction
		if d == directionRight {
			turnUnits = units / 90
		} else {
			turnUnits = units / -90
		}
		nextLat := (getDirectionIndex(waypoint.currentLat()) + turnUnits) % len(compass)
		nextLong := (getDirectionIndex(waypoint.currentLong()) + turnUnits) % len(compass)
		if nextLat < 0 {
			nextLat = len(compass) + nextLat
		}
		if nextLong < 0 {
			nextLong = len(compass) + nextLong
		}

		latDir = compass[nextLat]
		longDir = compass[nextLong]

		currentLat := int(math.Abs(float64(waypoint.lat)))
		currentLong := int(math.Abs(float64(waypoint.long)))
		switch latDir {
		case directionEast:
			waypoint.long = currentLat
		case directionSouth:
			waypoint.lat = -1 * currentLat
		case directionWest:
			waypoint.long = -1 * currentLat
		case directionNorth:
			waypoint.lat = currentLat
		}

		switch longDir {
		case directionEast:
			waypoint.long = currentLong
		case directionSouth:
			waypoint.lat = -1 * currentLong
		case directionWest:
			waypoint.long = -1 * currentLong
		case directionNorth:
			waypoint.lat = currentLong
		}
	}
}

func part1(p *pos, directionIndex int, d direction, units int) int {
	switch d {
	case directionNorth, directionEast, directionSouth, directionWest:
		advance(p, d, units)
	case directionRight:
		turnUnits := units / 90
		directionIndex = (directionIndex + turnUnits) % len(compass)
		fmt.Printf("turning right: %v, %v\n", units/90, string(compass[directionIndex]))
	case directionLeft:
		turnUnits := units / 90
		directionIndex = (directionIndex + (-1+len(compass))*turnUnits) % len(compass)
		fmt.Printf("turning left: %v, %v\n", string(compass[directionIndex]), units/-90)
	case directionForward:
		advance(p, compass[directionIndex], units)
	}
	return directionIndex
}

func advance(p *pos, d direction, units int) {
	switch d {
	case directionNorth:
		p.lat += units
	case directionEast:
		p.long += units
	case directionSouth:
		p.lat -= units
	case directionWest:
		p.long -= units
	}
}

func manhattanValue(p pos) float64 {
	return math.Abs(float64(p.lat)) + math.Abs(float64(p.long))
}

func parseLine(raw string) (direction, int) {
	num, err := strconv.Atoi(raw[1:])
	if err != nil {
		panic(err)
	}
	return direction(raw[0]), num
}
