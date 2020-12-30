package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var instructionRE = regexp.MustCompile(`(jmp|acc|nop) \+?(-?\d+)`)

type ctx struct {
	accumulator int
	pointer     int
}

type instruction interface {
	run(ctx) ctx
	ran() bool
	reset()
}

type base struct {
	count int
}

func (b base) ran() bool {
	return b.count > 0
}

func (b *base) reset() {
	b.count = 0
}

type jmp struct {
	base
	arg int
}

func (o *jmp) run(c ctx) ctx {
	o.count += 1
	return ctx{accumulator: c.accumulator, pointer: c.pointer + o.arg}
}

type acc struct {
	base
	arg int
}

func (o *acc) run(c ctx) ctx {
	return ctx{
		accumulator: c.accumulator + o.arg,
		pointer:     c.pointer + 1,
	}
}

type nop struct {
	base
	arg int
}

func (o *nop) run(c ctx) ctx {
	o.count += 1
	return ctx{
		accumulator: c.accumulator,
		pointer:     c.pointer + 1,
	}
}

func parseInstruction(line string) instruction {
	matches := instructionRE.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}
	op := matches[1]
	arg, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	switch op {
	case "acc":
		return &acc{arg: arg}
	case "jmp":
		return &jmp{arg: arg}
	case "nop":
		return &nop{arg: arg}
	}
	return nil
}

func main() {
	fmt.Printf(">> %v\n", fixLoop())
}

func accumulate(instructions []instruction) (int, bool) {
	c := ctx{
		accumulator: 0,
		pointer:     0,
	}
	current := instructions[0]
	var terminated bool
	for !current.ran() {
		c = current.run(c)
		if c.pointer > len(instructions)-1 {
			terminated = true
			break
		}
		current = instructions[c.pointer]
	}
	return c.accumulator, terminated
}

func fixLoop() int {
	instructions := getInstructions()
	for i := 0; i < len(instructions); i++ {
		if inst := instructions[i]; flipable(inst) {
			reset(instructions)
			cp := make([]instruction, len(instructions))
			copy(cp, instructions)
			cp[i] = flip(inst)
			a, terminated := accumulate(cp)
			if terminated {
				fmt.Printf("FLIP INDEX >> %v\n", i)
				return a
			}
		}
	}
	return 0
}

func reset(instructions []instruction) {
	for _, inst := range instructions {
		inst.reset()
	}
}

func getInstructions() []instruction {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	instructions := []instruction{}
	for _, line := range strings.Split(string(data), "\n") {
		inst := parseInstruction(line)
		if inst == nil {
			continue
		}
		instructions = append(instructions, inst)
	}
	return instructions
}

func flip(inst instruction) instruction {
	switch v := inst.(type) {
	case *jmp:
		return &nop{}
	case *nop:
		return &jmp{arg: v.arg}
	default:
		panic("not a jmp or a nop")
	}
}

func flipable(inst instruction) bool {
	switch inst.(type) {
	case *jmp:
		return true
	case *nop:
		return true
	default:
		return false
	}
}
