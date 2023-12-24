package main

// Credit to https://github.com/dannyvankooten/advent-of-code/blob/main/2023/20-pulse-propagation/main.go

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	// lines := strings.Split(string(data), "\n")

	modules := map[string]Module{
		"button": {
			typ:     TYPE_BUTTON,
			targets: []string{"broadcaster"},
		},
	}

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		m := Module{}
		pos := strings.Index(line, " ")
		var name string
		if line[0] == TYPE_FLIPFLOP || line[0] == TYPE_CONJUNCTION {
			m.typ = rune(line[0])
			name = line[1:pos]
		} else {
			m.typ = TYPE_BROADCASTER
			name = line[:pos]
		}
		line = line[strings.Index(line, " -> ")+4:]
		m.targets = strings.Split(line, ", ")
		m.memory = make(map[string]bool)
		modules[name] = m
	}

	// init memory
	for k, input := range modules {
		for _, t := range input.targets {
			dest := modules[t]
			if dest.typ == TYPE_CONJUNCTION {
				dest.memory[k] = false
			}
		}
	}

	factors := make(map[string]int)
	pulses := make([]Pulse, 0)
	for i := 1; len(factors) != len(modules["qt"].memory); i++ {
		pulses = pulses[:0]

		// first pulse is a low from button to broadcaster
		pulses = append(pulses, Pulse{
			source: "button",
			dest:   "broadcaster",
			value:  false,
		})

		// keep processing pulses until nothing in queue
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			pulses = append(pulses, pulse.Handle(modules)...)

			// check memory values for each input
			for k, v := range modules["qt"].memory {
				_, ok := factors[k]
				if !ok && v {
					fmt.Printf(">>>>>> %v, %v\n", modules["qt"], i)
					factors[k] = i
				}
			}
		}
	}

	product := 1
	for _, v := range factors {
		product *= v
	}

	fmt.Printf("pushes: %v\n", product)
}

type Module struct {
	typ     rune
	targets []string
	status  bool
	memory  map[string]bool
}

type Pulse struct {
	source string
	value  bool
	dest   string
}

func (p *Pulse) Handle(modules map[string]Module) []Pulse {
	sout := false
	m := modules[p.dest]

	switch m.typ {
	case TYPE_FLIPFLOP:
		if p.value {
			return nil
		} else {
			m.status = !m.status
			sout = m.status
		}
	case TYPE_CONJUNCTION:
		m.memory[p.source] = p.value
		sout = false
		for _, v := range m.memory {
			if !v {
				sout = true
				break
			}
		}
	case TYPE_BROADCASTER:
		sout = p.value
	}

	out := make([]Pulse, len(m.targets))
	for i, r := range m.targets {
		out[i] = Pulse{
			source: p.dest,
			dest:   r,
			value:  sout,
		}
	}

	modules[p.dest] = m

	return out
}

const TYPE_BUTTON = 'b'
const TYPE_BROADCASTER = 'a'
const TYPE_CONJUNCTION = '&'
const TYPE_FLIPFLOP = '%'
