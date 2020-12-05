package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Passport has all components of a passport
type Passport struct {
	byr bool
	iyr bool
	eyr bool
	hgt bool
	hcl bool
	ecl bool
	pid bool
	cid bool // optional
}

func checkIfCorrectPassport(p Passport) bool {
	return p.byr && p.iyr && p.eyr &&
		p.hgt && p.hcl && p.ecl &&
		p.pid
}

func partOne(passportData []string) int {
	var totalCorrectPassports int

	currentPassport := Passport{}
	for _, r := range passportData {
		if len(r) == 0 {
			// check if passport is valid
			if checkIfCorrectPassport(currentPassport) {
				totalCorrectPassports++
			}
			// reset the passport
			currentPassport = Passport{}
			continue
		}

		c := strings.Split(r, " ")
		for _, stuff := range c {
			p := strings.Split(stuff, ":")
			token := p[0]
			switch token {
			case "byr":
				currentPassport.byr = true
			case "iyr":
				currentPassport.iyr = true
			case "eyr":
				currentPassport.eyr = true
			case "hgt":
				currentPassport.hgt = true
			case "hcl":
				currentPassport.hcl = true
			case "ecl":
				currentPassport.ecl = true
			case "pid":
				currentPassport.pid = true
			case "cid":
				currentPassport.cid = true
			}
		}

	}

	if checkIfCorrectPassport(currentPassport) {
		totalCorrectPassports++
	}

	return totalCorrectPassports
}

var heightRegex = regexp.MustCompile(`(?m)([0-9]{2,3})+(cm|in)`)
var hairRegex = regexp.MustCompile(`(?m)#([a-f]|[0-9]){6}`)
var eyeRegex = regexp.MustCompile(`(?m)(amb|blu|brn|gry|grn|hzl|oth)`)
var pidRegex = regexp.MustCompile(`(?m)^(\d){9}$`)

func partTwo(passportData []string) int {
	var totalCorrectPassports int

	currentPassport := Passport{}
	for _, r := range passportData {
		if len(r) == 0 {
			// check if passport is valid
			if checkIfCorrectPassport(currentPassport) {
				totalCorrectPassports++
			}
			// reset the passport
			currentPassport = Passport{}
			continue
		}

		c := strings.Split(r, " ")
		for _, stuff := range c {
			p := strings.Split(stuff, ":")
			token := p[0]
			value := p[1]
			switch token {
			case "byr":
				y, _ := strconv.Atoi(value)
				if y >= 1920 && y <= 2002 {
					currentPassport.byr = true
				}
			case "iyr":
				y, _ := strconv.Atoi(value)
				if y >= 2010 && y <= 2020 {
					currentPassport.iyr = true
				}
			case "eyr":
				y, _ := strconv.Atoi(value)
				if y >= 2020 && y <= 2030 {
					currentPassport.eyr = true
				}
			case "hgt":
				for _, match := range heightRegex.FindAllStringSubmatch(value, -1) {
					h, _ := strconv.Atoi(match[1])
					um := match[2]
					if (um == "in" && h >= 59 && h <= 76) ||
						(um == "cm" && h >= 150 && h <= 193) {
						currentPassport.hgt = true
					}
				}
				// currentPassport.hgt = true
			case "hcl":
				if hairRegex.MatchString(value) {
					currentPassport.hcl = true
				}
			case "ecl":
				if eyeRegex.MatchString(value) {
					currentPassport.ecl = true
				}
			case "pid":
				if pidRegex.MatchString(value) {
					currentPassport.pid = true
				}
			case "cid":
				currentPassport.cid = true
			}
		}

	}

	if checkIfCorrectPassport(currentPassport) {
		totalCorrectPassports++
	}

	return totalCorrectPassports
}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day04.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(getInputAsStringSlice(data))
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Now().Sub(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(getInputAsStringSlice(data))
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Now().Sub(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Now().Sub(totalTimeStart))

}
