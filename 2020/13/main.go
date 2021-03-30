package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	var shortest int
	var shortestID int
	for _, raw := range strings.Split(lines[1], ",") {
		if raw == "x" {
			continue
		}
		id, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}
		times := int(timestamp / id)
		_, remainder := math.Modf(float64(timestamp) / float64(id))
		fmt.Printf("remainder >> %v ", remainder)
		if remainder == 0 {
			shortestID = id
			break
		} else if remainder > 0 {
			times++
		}
		current := id * times
		if shortest == 0 || current < shortest {
			shortest = current
			shortestID = id
		}
	}
	fmt.Printf("shortest: %d, %d\n", shortestID, shortestID*(shortest-timestamp))
}
