package main

import (
	// "bytes"
	// "crypto/md5"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	env := Env{
		Modules: map[string]Module{
			"output": &Output{},
		},
	}

	receiverToSenders := map[string][]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "->")
		sender := strings.TrimSpace(parts[0])
		receiverParts := strings.Split(parts[1], ",")
		var receivers []string
		for _, p := range receiverParts {
			receivers = append(receivers, strings.TrimSpace(p))
		}

		var senderName string
		switch sender[0] {
		case '%':
			env.Modules[sender[1:]] = &FlipFlop{
				On:        false,
				Receivers: receivers,
			}
			senderName = sender[1:]
		case '&':
			env.Modules[sender[1:]] = &Conjunction{
				Memory:    map[string]bool{},
				Receivers: receivers,
			}
			senderName = sender[1:]
		default:
			env.Modules[sender] = &Broadcaster{
				Receivers: receivers,
			}
			senderName = sender
		}

		for _, r := range receivers {
			receiverToSenders[r] = append(receiverToSenders[r], senderName)
		}
	}

	for k, v := range env.Modules {
		con, ok := v.(*Conjunction)
		if !ok {
			continue
		}

		// fmt.Printf(">>>>>>>>>>>>>>>>>>> %v, %v\n", k, receiverToSenders[k])

		for _, s := range receiverToSenders[k] {
			con.Memory[s] = false
		}
	}

	// con := env.Modules["con"].(*Conjunction)
	// con.Memory["b"] = false
	// con.Memory["a"] = false

	// initial := env.String() //md5.Sum([]byte(env.String()))
	// var state string
	// var state []byte
	//
	var counts []Count

	// fmt.Printf("%v\n", env.String())

	queue := []Pulse{{Sender: "button", High: false, Recipient: "broadcaster"}}
	for i := 0; i < 1000; i++ {
		var c Count
		for len(queue) > 0 {
			var pulse Pulse
			pulse, queue = queue[0], queue[1:]
			// fmt.Printf(">> %+v\n", pulse)

			if pulse.High {
				c.High++
			} else {
				c.Low++
			}

			// fmt.Printf("%v\n", pulse.Recipient)
			module, ok := env.Modules[pulse.Recipient]
			if !ok {
				continue
			}
			pulses := module.Receive(pulse)
			for i := range pulses {
				pulses[i].Sender = pulse.Recipient
			}

			// for _, p := range pulses {
			// 	fmt.Printf(">>>> %+v\n", p)
			// }

			if len(pulses) > 0 {
				queue = append(queue, pulses...)
			}

			// next := md5.Sum([]byte(env.String()))
			// state = next[:]

			// fmt.Printf("%v\n\n", env.String())
		}

		// state = env.String()
		counts = append(counts, c)
		queue = append(queue, Pulse{Sender: "button", High: false, Recipient: "broadcaster"})

		// next := md5.Sum([]byte(env.String()))
		// state = next[:]
	}

	// fmt.Printf("counts: %+v\n", counts)
	//
	// fmt.Printf("%v -> %v\n", env, initial)

	var highs int
	var lows int
	for i := 0; i < 1000; i++ {
		highs += counts[i%len(counts)].High
		lows += counts[i%len(counts)].Low
	}

	fmt.Printf("high: %v, low: %v, combined: %v\n", highs, lows, highs*lows)
}

type Count struct {
	High int
	Low  int
}

type Module interface {
	Receive(Pulse) []Pulse
}

type Pulse struct {
	Sender    string
	High      bool
	Recipient string
}

type Env struct {
	Modules map[string]Module
}

func (e Env) String() string {
	var keys []string
	for k := range e.Modules {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%v", e.Modules[k]))
	}

	return strings.Join(parts, " ")
}

type Broadcaster struct {
	Receivers []string
}

func (b *Broadcaster) String() string {
	return "b:"
}

func (b *Broadcaster) Receive(p Pulse) []Pulse {
	next := make([]Pulse, len(b.Receivers))
	for i, r := range b.Receivers {
		next[i] = Pulse{High: p.High, Recipient: r}
	}
	return next
}

type FlipFlop struct {
	On        bool
	Receivers []string
}

func (f *FlipFlop) String() string {
	return fmt.Sprintf("%%:%v", f.On)
}

func (f *FlipFlop) Receive(p Pulse) []Pulse {
	if p.High {
		return nil
	}

	next := make([]Pulse, len(f.Receivers))
	for i, r := range f.Receivers {
		next[i] = Pulse{Recipient: r}
		if f.On {
			next[i].High = false
		} else {
			next[i].High = true
		}
	}

	f.On = !f.On

	return next
}

type Conjunction struct {
	Memory    map[string]bool
	Receivers []string
}

func (c *Conjunction) String() string {
	var keys []string
	for k := range c.Memory {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	str := fmt.Sprintf("&:")
	for _, key := range keys {
		str += fmt.Sprintf("%s=%v,", key, c.Memory[key])
	}

	return str
}

func (c *Conjunction) Receive(p Pulse) []Pulse {
	c.Memory[p.Sender] = p.High

	high := false
	for _, v := range c.Memory {
		if !v {
			high = true
			break
		}
	}

	next := make([]Pulse, len(c.Receivers))
	for i, r := range c.Receivers {
		next[i] = Pulse{Recipient: r, High: high}
	}

	return next
}

type Output struct {
}

func (o Output) String() string {
	return "output"
}

func (o Output) Receive(p Pulse) []Pulse {
	return nil
}
