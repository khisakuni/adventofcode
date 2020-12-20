package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var colorStr = `(\w+ \w+) bag`
var colorRE = regexp.MustCompile(colorStr)
var colorNumberRE = regexp.MustCompile(`(\d+) ` + colorStr)

type node struct {
	value int
	Nodes map[string]*node
}

func (n *node) hasNode(target, current string) bool {
	currentHas := false
	if n.Nodes[current] == nil {
		return false
	}
	for k := range n.Nodes[current].Nodes {
		if k == target {
			currentHas = true
		}
		if n.hasNode(target, k) {
			currentHas = true
		}
	}
	return currentHas
}

func newNode(value int) *node {
	return &node{
		value: value,
		Nodes: map[string]*node{},
	}
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	root := newNode(0)
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		heads := colorRE.FindStringSubmatch(line)
		if len(heads) == 0 {
			continue
		}
		head := heads[1]
		nodes := colorNumberRE.FindAllStringSubmatch(line, -1)
		for _, n := range nodes {
			name := n[2]
			val, err := strconv.Atoi(n[1])
			if err != nil {
				panic(err)
			}
			headNode := root.Nodes[head]
			if headNode == nil {
				headNode = newNode(0)
			}
			headNode.Nodes[name] = newNode(val)
			root.Nodes[head] = headNode
		}
	}
	target := "shiny gold"
	count := 0
	for k := range root.Nodes {
		if root.hasNode(target, k) {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
