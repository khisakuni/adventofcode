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

var compass = []direction{directionEast, directionSouth, directionWest, directionNorth}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	directionIndex := 0
	p := pos{}
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		fmt.Println(line)
		d, units := parseLine(line)
		switch d {
		case directionNorth, directionEast, directionSouth, directionWest:
			advance(&p, d, units)
		case directionRight:
			turnUnits := units / 90
			directionIndex = (directionIndex + turnUnits) % len(compass)
			fmt.Printf("turning right: %v, %v\n", (units / 90), string(compass[directionIndex]))
		case directionLeft:
			turnUnits := units / -90
			directionIndex = directionIndex + (turnUnits % len(compass))
			if directionIndex < 0 {
				directionIndex = len(compass) + directionIndex
			}
			fmt.Printf("turning left: %v, %v\n", string(compass[directionIndex]), (units / -90))
		case directionForward:
			advance(&p, compass[directionIndex], units)
		}
	}

	fmt.Printf("pos >> %+v\n", p)
	fmt.Printf("manhattan distance: %v\n", manhattanValue(p))
}

func advance(p *pos, d direction, units int) {
	fmt.Printf("moving %s by %d\n", string(d), units)
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
	fmt.Printf("pos: %+v\n\n", p)
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
