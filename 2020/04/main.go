package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)

 */

var byrRE = regexp.MustCompile(`byr:(\d{4})`)
var iyrRE = regexp.MustCompile(`iyr:(\d{4})`)
var eyrRE = regexp.MustCompile(`eyr:(\d{4})`)
var hgtRE = regexp.MustCompile(`hgt:(\d+)(cm|in)?`)
var hclRE = regexp.MustCompile(`hcl:(.+)`)
var eclRE = regexp.MustCompile(`ecl:(.+)`)
var pidRE = regexp.MustCompile(`pid:(.+)`)
var cidRE = regexp.MustCompile(`cid:(\d+)`)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n\n")
	var count int
	for _, str := range strs {
		p := parsePassport(str)
		if p.isValid() {
			count++
		} else {
			fmt.Printf("invalid: %+v\n", p)
		}
	}
	fmt.Printf("Valid count: %d\n", count)
}

type passport struct {
	byr int
	iyr int
	eyr int
	hgt int
	hcl string
	ecl string
	pid string
	cid int
}

func (p passport) isValid() bool {
	return p.byr != 0 && p.iyr != 0 && p.eyr != 0 && p.hgt != 0 && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func parsePassport(str string) passport {
	return passport{
		byr: parseInt(str, byrRE),
		iyr:  parseInt(str, iyrRE),
		eyr: parseInt(str, eyrRE),
		hgt: parseInt(str, hgtRE),
		hcl: parseString(str, hclRE),
		ecl: parseString(str, eclRE),
		pid: parseString(str, pidRE),
		cid: parseInt(str, cidRE),
	}
}

func parseString(str string, re *regexp.Regexp) string {
	matches := re.FindStringSubmatch(str)
	if len(matches) == 0 {
		return ""
	}
	return matches[1]
}

func parseInt(str string, re *regexp.Regexp) int {
	matches := re.FindStringSubmatch(str)
	if len(matches) == 0 {
		return 0
	}
	i, _ := strconv.Atoi(matches[1])
	return i
}
