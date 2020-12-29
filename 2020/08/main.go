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
}

type base struct {
	didRun bool
}

func (b base) ran() bool {
	return b.didRun
}

type jmp struct {
	base
	arg int
}

func (o *jmp) run(c ctx) ctx {
	o.didRun = true
	return ctx{accumulator: c.accumulator, pointer: c.pointer + o.arg}
}

type acc struct {
	base
	arg int
}

func (o *acc) run(c ctx) ctx {
	o.didRun = true
	return ctx{
		accumulator: c.accumulator + o.arg,
		pointer:     c.pointer + 1,
	}
}

type nop struct {
	base
}

func (o *nop) run(c ctx) ctx {
	o.didRun = true
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
		return &nop{}
	}
	return nil
}

func main() {
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
	c := ctx{
		accumulator: 0,
		pointer:     0,
	}
	current := instructions[0]
	for !current.ran() {
		c = current.run(c)
		current = instructions[c.pointer]
	}
	fmt.Printf("acc: %d\n", c.accumulator)
}
