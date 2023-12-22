package main

import (
	// "errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// func parseWorkFlows(lines []string) (map[string]string, error) {
// 	var key string
// 	wfs := make(map[string]string)
// 	for _, line := range lines {
// 		if len(line) > 2 {
// 			ws := strings.Split(line[:len(line)-1], "{")
// 			if len(ws) != 2 {
// 				return wfs, errors.New("invalid input")
// 			}
// 			ops := strings.Split(ws[1], ",")
// 			if len(ops) == 2 {
// 				key := ws[0]
// 				wfs[key] = ws[1]
// 			} else {
// 				for i := 0; i < len(ops)-1; i++ {
// 					if i == 0 {
// 						key = ws[0]
// 					} else {
// 						key = fmt.Sprintf("%s%d", ws[0], i)
// 					}
// 					if i == len(ops)-2 {
// 						wfs[key] = fmt.Sprintf("%s,%s", ops[i], ops[i+1])
// 					} else {
// 						wfs[key] = fmt.Sprintf("%s,%s%d", ops[i], ws[0], i+1)
// 					}
// 				}
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	return wfs, nil
// }
//
// func getAccepted(wfKey string, xMin, xMax, mMin, mMax, aMin, aMax, sMin, sMax int, wfs map[string]string) int {
// 	if wfKey == "A" {
// 		return (xMax - xMin + 1) * (mMax - mMin + 1) * (aMax - aMin + 1) * (sMax - sMin + 1)
// 	} else if wfKey == "R" {
// 		return 0
// 	} else {
// 		acc := 0
// 		condition := wfs[wfKey]
// 		a := strings.Split(condition, ":")
// 		options := strings.Split(a[1], ",")
// 		trueOption := options[0]
// 		falseOption := options[1]
// 		param := a[0][0]
// 		symbol := a[0][1]
// 		value, _ := strconv.Atoi(a[0][2:])
// 		switch param {
// 		case 'x':
// 			if symbol == '>' {
// 				acc += getAccepted(trueOption, value+1, xMax, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, value, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
// 			} else {
// 				acc += getAccepted(trueOption, xMin, value-1, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, value, xMax, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
// 			}
// 		case 'm':
// 			if symbol == '>' {
// 				acc += getAccepted(trueOption, xMin, xMax, value+1, mMax, aMin, aMax, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, mMin, value, aMin, aMax, sMin, sMax, wfs)
// 			} else {
// 				acc += getAccepted(trueOption, xMin, xMax, mMin, value-1, aMin, aMax, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, value, mMax, aMin, aMax, sMin, sMax, wfs)
// 			}
// 		case 'a':
// 			if symbol == '>' {
// 				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, value+1, aMax, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, value, sMin, sMax, wfs)
// 			} else {
// 				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, value-1, sMin, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, value, aMax, sMin, sMax, wfs)
// 			}
// 		case 's':
// 			if symbol == '>' {
// 				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, aMax, value+1, sMax, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value, wfs)
// 			} else {
// 				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value-1, wfs)
// 				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, aMax, value, sMax, wfs)
// 			}
// 		}
// 		return acc
// 	}
// }
//
// func getAllCombinations(input []string) (int, error) {
// 	if wfs, err := parseWorkFlows(input); err != nil {
// 		return 0, err
// 	} else {
// 		for k, wf := range wfs {
// 			fmt.Printf(">>> %v: %v\n", k, wf)
// 		}
// 		return getAccepted("in", 1, 4000, 1, 4000, 1, 4000, 1, 4000, wfs), nil
// 	}
// }

func main() {
	data, _ := os.ReadFile("input.txt")

	wfRe := regexp.MustCompile("([a-z]+){(.*)}")

	ruleRe := regexp.MustCompile(`([a-z]+)(>|<)(\d+):([a-zA-Z]+)`)

	lines := strings.Split(string(data), "\n")

	workflows := map[string][]Rule{}

	for _, line := range lines {
		if line == "" {
			break
		}

		parts := wfRe.FindStringSubmatch(line)[1:]
		wfName := parts[0]
		ruleParts := strings.Split(parts[1], ",")
		for i, part := range ruleParts {
			if i == len(ruleParts)-1 {
				workflows[wfName] = append(workflows[wfName], Rule{Next: part})
				continue
			}
			p := ruleRe.FindStringSubmatch(part)[1:]
			v := p[0]
			op := p[1]
			val, _ := strconv.Atoi(p[2])
			next := p[3]

			r := Rule{
				Workflow: wfName,
				Var:      v,
				Op:       op,
				Val:      val,
				Next:     next,
			}

			workflows[wfName] = append(workflows[wfName], r)
		}

	}

	// for k, v := range workflows {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }

	combos := dfs(workflows, "in", Spec{
		Max: map[string]int{
			"x": 4000,
			"m": 4000,
			"a": 4000,
			"s": 4000,
		},
		Min: map[string]int{
			"x": 1,
			"m": 1,
			"a": 1,
			"s": 1,
		},
	})

	// expected := 167409079868000

	// fmt.Printf("specs: %v, diff: (%v)\n", combos, expected-combos)
	// accepted := graph["A"]
	// for _, node := range accepted {
	//   if node.Var == "" {
	//     fmt.Printf("found root! %v\n", node.Workflow)
	//   }
	//
	// }
	// val, err := getAllCombinations(lines)
	fmt.Printf("answer: %v\n", combos)
}

func dfs(graph map[string][]Rule, name string, spec Spec) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		// fmt.Printf("found! max: %v, min: %v\n", spec.Max, spec.Min)
		total := (spec.Max["x"] - spec.Min["x"] + 1) * (spec.Max["m"] - spec.Min["m"] + 1) * (spec.Max["a"] - spec.Min["a"] + 1) * (spec.Max["s"] - spec.Min["s"] + 1)
		return total
	}

	rules := graph[name]
	neg := Spec{
		Max: copyMap(spec.Max),
		Min: copyMap(spec.Min),
	}
	var total int
	// fmt.Printf("wf: %v: rules: %v\n", name, rules)
	for _, r := range rules {
		if r.Op == ">" {
			s := Spec{
				Max: copyMap(spec.Max),
				Min: copyMap(spec.Min),
			}
			for k, v := range neg.Max {
				s.Max[k] = v
			}
			for k, v := range neg.Min {
				s.Min[k] = v
			}
			s.Min[r.Var] = r.Val + 1
			neg.Max[r.Var] = r.Val
			total += dfs(graph, r.Next, s)
		} else if r.Op == "<" {
			s := Spec{
				Max: copyMap(spec.Max),
				Min: copyMap(spec.Min),
			}
			for k, v := range neg.Max {
				s.Max[k] = v
			}
			for k, v := range neg.Min {
				s.Min[k] = v
			}
			s.Max[r.Var] = r.Val - 1
			neg.Min[r.Var] = r.Val
			total += dfs(graph, r.Next, s)
		} else {
			total += dfs(graph, r.Next, neg)
		}
	}

	// fmt.Printf("dead end\n")
	return total
}

func copyMap(in map[string]int) map[string]int {
	out := map[string]int{}
	for k, v := range in {
		out[k] = v
	}

	return out
}

func parseInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

type Spec struct {
	Max map[string]int
	Min map[string]int
}

type Rule struct {
	Workflow string
	Var      string
	Val      int
	Op       string
	Next     string
}

func (r Rule) String() string {
	if r.Var == "" {
		return fmt.Sprintf("(-> %v)", r.Next)
	}
	return fmt.Sprintf("(%s %s %d -> %v)", r.Var, r.Op, r.Val, r.Next)
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
