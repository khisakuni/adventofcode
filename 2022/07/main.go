package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	type node struct {
		name     string
		size     int
		subNodes map[string]*node
	}

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	commandRegex := regexp.MustCompile(`^\$ (cd|ls) ?([a-z\/\.]*)`)
	fileRegex := regexp.MustCompile(`^([0-9]+) ([a-z.]+)$`)
	dirRegex := regexp.MustCompile(`^dir ([a-z]+)$`)

	stack := []*node{
		{
			name:     "ROOT",
			subNodes: map[string]*node{},
		},
	}

	var parts []string

	m := map[string]*node{}

	var head *node
	strs := strings.Split(string(input), "\n")
	for _, str := range strs {
		if str == "" {
			continue
		}

		head = stack[len(stack)-1]
		commandMatches := commandRegex.FindStringSubmatch(str)
		if len(commandMatches) > 0 {
			command := commandMatches[1]
			switch command {
			case "cd":
				arg := commandMatches[2]
				switch arg {
				case "..":
					stack = stack[:len(stack)-1]

					var size int
					for _, sub := range head.subNodes {
						size += sub.size
					}

					head.size = size
					p := path.Join(parts...)
					if head.size <= 100000 {
						m[p] = head
					} else {
						m[p] = nil
					}

					//fmt.Printf("Pop: %v (%v)\n", head.name, head.size)

					parts = parts[:len(parts)-1]
					head = stack[len(stack)-1]

					continue
				default:
					// Find or create child node and push onto stack
					sub, ok := head.subNodes[arg]
					if !ok {
						sub = &node{
							name:     arg,
							size:     0,
							subNodes: map[string]*node{},
						}
						head.subNodes[arg] = sub
					}

					stack = append(stack, sub)
					parts = append(parts, arg)

					continue
				}
			case "ls":
				// Ignore?
				//fmt.Printf("List for %v (%v)\n", head.name, head.size)
				continue
			}
		}

		fileMatches := fileRegex.FindStringSubmatch(str)
		var size int
		var name string
		if len(fileMatches) > 0 {
			size, _ = strconv.Atoi(fileMatches[1])
			name = fileMatches[2]
		}

		dirMatches := dirRegex.FindStringSubmatch(str)
		if len(dirMatches) > 0 {
			name = dirMatches[1]
		}

		//fmt.Printf("entry: %v (%v)\n", name, size)

		_, ok := head.subNodes[name]
		if !ok {
			head.subNodes[name] = &node{
				name:     name,
				size:     size,
				subNodes: map[string]*node{},
			}
		}
	}

	for len(stack) > 1 {
		stack = stack[:len(stack)-1]
		var size int
		for _, s := range head.subNodes {
			size += s.size
		}

		head.size = size
		p := path.Join(parts...)
		if head.size <= 100000 {
			m[p] = head
		} else {
			m[p] = nil
		}

		//fmt.Printf("Pop: %v (%v)\n", head.name, head.size)

		parts = parts[:len(parts)-1]
		nextHead := stack[len(stack)-1]
		head = nextHead
	}

	var size int
	for _, n := range m {
		if n != nil {
			size += n.size
			//fmt.Printf("name: %s, size: %d\n", name, n.size)
		}
	}

	fmt.Printf("Size: %v\n", size)
}
