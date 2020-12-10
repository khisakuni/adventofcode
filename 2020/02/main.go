package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)


func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n")
	var count int
	for _, str := range strs {
		if str == "" {
			continue
		}
		 pass := parsePassword(str)
		 if pass.isValid() {
		 		count++
		 }
	}
	fmt.Printf("Valid count: %d\n", count)
}

type password struct {
	letter rune
	min int
	max int
	value string
	frequency map[rune]int
}

var passwordRe = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

const (
	minIndex = 1
	maxIndex = 2
	letterIndex = 3
	valueIndex = 4
)

func parsePassword(raw string) password {
	matches := passwordRe.FindStringSubmatch(raw)
	min, err := strconv.Atoi(matches[minIndex])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(matches[maxIndex])
	if err != nil {
		panic(err)
	}
	letter := []rune(matches[letterIndex])[0]
	value := matches[valueIndex]
	freq := map[rune]int{}
	for _, v := range value {
		freq[v]++
	}
	return password{
		min: min,
		max: max,
		letter: letter,
		value: value,
		frequency: freq,
	}
}

func (p password) isValid() bool {
	count := p.frequency[p.letter]
	return count <= p.max && count >= p.min
}
