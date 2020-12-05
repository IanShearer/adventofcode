package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func partOne(floors string) int {
	var numberOfFloors int
	for _, floor := range floors {
		if floor == '(' {
			numberOfFloors++
		} else {
			numberOfFloors--
		}
	}

	return numberOfFloors
}

func partTwo(floors string) int {
	currentFloor := 1
	var numberOfFloors int
	for _, floor := range floors {
		if floor == '(' {
			currentFloor++
		} else {
			currentFloor--
		}
		numberOfFloors++
		if currentFloor == 0 {
			break
		}
	}

	return numberOfFloors
}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day01.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(string(data))
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Since(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(string(data))
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Since(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Since(totalTimeStart))

}
