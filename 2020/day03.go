package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func partOne(slope []string, movesX int, movesY int) int {

	var numOfTrees int
	currentXPos := 0
	width := len(slope[0])
	for i := 0; i < len(slope); i = i + movesY {
		if i == 0 {
			continue
		}
		currentXPos += movesX
		if currentXPos >= width {
			currentXPos = currentXPos - width
		}
		if slope[i][currentXPos] == '#' {
			numOfTrees++
		}
	}

	return numOfTrees
}

func partTwo(slope []string) int {

	a := partOne(slope, 1, 1)
	b := partOne(slope, 3, 1)
	c := partOne(slope, 5, 1)
	d := partOne(slope, 7, 1)
	e := partOne(slope, 1, 2)

	return a * b * c * d * e

}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("2020/inputs/day03.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	// format your input here
	info := getInputAsStringSlice(data)

	partOneStart := time.Now()
	partOneAnswer := partOne(info, 3, 1)
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Now().Sub(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(info)
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Now().Sub(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Now().Sub(totalTimeStart))

}
