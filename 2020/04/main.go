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

var byrRE = regexp.MustCompile(`byr:(\d{4})\b`)
var iyrRE = regexp.MustCompile(`iyr:(\d{4})\b`)
var eyrRE = regexp.MustCompile(`eyr:(\d{4})\b`)
var hgtRE = regexp.MustCompile(`hgt:(\d+)(cm|in)\b`)
var systemRE = regexp.MustCompile(`hgt:\d+(cm|in)\b`)
var hclRE = regexp.MustCompile(`hcl:#([0-9a-f]{6})\b`)
var eclRE = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)\b`)
var pidRE = regexp.MustCompile(`pid:([0-9]{9})\b`)
var cidRE = regexp.MustCompile(`cid:(.+)`)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(data), "\n\n")
	var count int
	for _, str := range strs {
		p := parsePassport(str)
		if p.isValidV2() {
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
	system string
}

func (p passport) isValid() bool {
	return p.byr != 0 && p.iyr != 0 && p.eyr != 0 && p.hgt != 0 && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p passport) isValidV2() bool {
	return p.hasValidBYR() &&
		p.hasValidECL() &&
		p.hasValidEYR() &&
		p.hasValidHCL() &&
		p.hasValidHGT() &&
		p.hasValidIYR() &&
		p.hasValidPID()
}

func (p passport) hasValidBYR() bool {
	return p.byr >= 1920 && p.byr <= 2002
}

func (p passport) hasValidIYR() bool {
	return p.iyr >= 2010 && p.iyr <= 2020
}

func (p passport) hasValidEYR() bool {
	return p.eyr >= 2020 && p.eyr <= 2030
}

func (p passport) hasValidHGT() bool {
	if p.system == "cm" {
		return p.hgt >= 150 && p.hgt <= 193
	}
	if p.system == "in" {
		return p.hgt >= 59 && p.hgt <= 76
	}
	return false
}

func (p passport) hasValidHCL() bool {
	return p.hcl != ""
}

func (p passport) hasValidECL() bool {
	return p.ecl != ""
}

func (p passport) hasValidPID() bool {
	return p.pid != ""
}

func parsePassport(str string) passport {
	return passport{
		byr: parseInt(str, byrRE),
		iyr: parseInt(str, iyrRE),
		eyr: parseInt(str, eyrRE),
		hgt: parseInt(str, hgtRE),
		system: parseString(str, systemRE),
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
