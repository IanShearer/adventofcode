package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func parseLine(line string) (int, int, []rune, string) {

	d := strings.Split(line, " ")
	var start int
	var end int
	var err error

	startEnd := strings.Split(d[0], "-")
	start, err = strconv.Atoi(startEnd[0])
	if err != nil {
		panic("could not convert to number")
	}
	end, err = strconv.Atoi(startEnd[1])
	if err != nil {
		panic("could not convert to number")
	}

	character := d[1]
	password := d[2]

	return start, end, []rune(character), password

}

func partOne(someInput []string) int {

	var totalCorrectPasswords int
	for _, line := range someInput {
		s, e, c, p := parseLine(line)
		var characterCount int
		for _, character := range p {
			if character == c[0] {
				characterCount++
			}
		}
		if characterCount >= s && characterCount <= e {
			totalCorrectPasswords++
		}
	}

	return totalCorrectPasswords

}

func partTwo(someInput []string) int {

	var totalCorrectPasswords int
	for _, line := range someInput {
		s, e, c, p := parseLine(line)
		runes := []rune(p)
		if (runes[s-1] == c[0] || runes[e-1] == c[0]) && runes[e-1] != runes[s-1] {
			totalCorrectPasswords++
		}
	}

	return totalCorrectPasswords

}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day02.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	d := getInputAsStringSlice(data)

	partOneStart := time.Now()
	partOneAnswer := partOne(d)
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Now().Sub(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(d)
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Now().Sub(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Now().Sub(totalTimeStart))

}
