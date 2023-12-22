package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	wfRe := regexp.MustCompile("([a-z]+){(.*)}")
	partRe := regexp.MustCompile("{(.*)}")

	ruleRe := regexp.MustCompile(`([a-z]+)(>|<)(\d+):([a-zA-Z]+)`)

	lines := strings.Split(string(data), "\n")

	workflows := map[string]Workflow{}

	var parts []Part
	var processParts bool
	for _, line := range lines {
		if line == "" {
			processParts = true
			continue
		}

		if processParts {
			valsStr := partRe.FindStringSubmatch(line)[1]
			vals := strings.Split(valsStr, ",")
			part := Part{
				X: parseInt(vals[0][2:]),
				M: parseInt(vals[1][2:]),
				A: parseInt(vals[2][2:]),
				S: parseInt(vals[3][2:]),
			}
			parts = append(parts, part)
		} else {
			parts := wfRe.FindStringSubmatch(line)[1:]
			wfName := parts[0]
			var wf Workflow

			var rules []Rule
			ruleParts := strings.Split(parts[1], ",")
			for i, part := range ruleParts {
				if i == len(ruleParts)-1 {
					rules = append(rules, Rule{Next: part})
					continue
				}
				p := ruleRe.FindStringSubmatch(part)[1:]
				v := p[0]
				op := p[1]
				val, _ := strconv.Atoi(p[2])
				next := p[3]

				rules = append(rules, Rule{
					Var:  v,
					Op:   op,
					Val:  val,
					Next: next,
				})
			}

			wf.Rules = rules
			workflows[wfName] = wf
		}
	}

	var sum int
	for _, p := range parts {
		if workflows["in"].Eval(workflows, p) {
			sum += p.X + p.M + p.A + p.S
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func parseInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

type Workflow struct {
	Rules []Rule
}

func (w Workflow) Eval(wfs map[string]Workflow, p Part) bool {
	for _, r := range w.Rules {
		out := r.Eval(p)
		if out.Accepted {
			return true
		}
		if out.Rejected {
			return false
		}
		if out.Next != "" {
			return wfs[out.Next].Eval(wfs, p)
		}
	}

	return false
}

type Rule struct {
	Var  string
	Val  int
	Op   string
	Next string
}

func (r Rule) Eval(p Part) Outcome {
	switch r.Op {
	case "":
		if r.Next == "A" {
			return Outcome{Accepted: true}
		}
		if r.Next == "R" {
			return Outcome{Rejected: true}
		}

		return Outcome{Next: r.Next}
	case ">":
		var val int
		switch r.Var {
		case "x":
			val = p.X
		case "m":
			val = p.M
		case "a":
			val = p.A
		case "s":
			val = p.S
		}

		if val > r.Val {
			if r.Next == "A" {
				return Outcome{Accepted: true}
			}
			if r.Next == "R" {
				return Outcome{Rejected: true}
			}

			return Outcome{Next: r.Next}

		}

		return Outcome{}
	case "<":
		var val int
		switch r.Var {
		case "x":
			val = p.X
		case "m":
			val = p.M
		case "a":
			val = p.A
		case "s":
			val = p.S
		}

		if val < r.Val {
			if r.Next == "A" {
				return Outcome{Accepted: true}
			}
			if r.Next == "R" {
				return Outcome{Rejected: true}
			}

			return Outcome{Next: r.Next}

		}

		return Outcome{}
	}

	return Outcome{}
}

type Outcome struct {
	Accepted bool
	Rejected bool
	Next     string
}

type Part struct {
	X int
	M int
	A int
	S int
}
