package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type coordinates struct {
	x int
	y int
}

func partOne(directions string) int {

	m := make(map[coordinates]struct{}, 0)
	previousCoordinate := coordinates{}
	m[previousCoordinate] = struct{}{}
	for _, d := range directions {
		switch d {
		case '^':
			previousCoordinate.y++
		case '>':
			previousCoordinate.x++
		case 'v':
			previousCoordinate.y--
		case '<':
			previousCoordinate.x--
		}
		m[previousCoordinate] = struct{}{}
	}

	return len(m)

}

func partTwo(directions string) int {

	m := make(map[coordinates]struct{}, 0)
	santaCoordinates := coordinates{}
	roboSantaCoordinates := coordinates{}
	m[santaCoordinates] = struct{}{}
	for i, d := range directions {
		switch d {
		case '^':
			if i%2 == 0 {
				santaCoordinates.y++
			} else {
				roboSantaCoordinates.y++
			}
		case '>':
			if i%2 == 0 {
				santaCoordinates.x++
			} else {
				roboSantaCoordinates.x++
			}
		case 'v':
			if i%2 == 0 {
				santaCoordinates.y--
			} else {
				roboSantaCoordinates.y--
			}
		case '<':
			if i%2 == 0 {
				santaCoordinates.x--
			} else {
				roboSantaCoordinates.x--
			}
		}
		if i%2 == 0 {
			m[santaCoordinates] = struct{}{}
		} else {
			m[roboSantaCoordinates] = struct{}{}
		}
	}

	return len(m)

}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day03.txt")
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
