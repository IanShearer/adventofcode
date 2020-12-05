package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

func getSeatID(row int, column int) int {
	return (row * 8) + column
}

func trverseRowCol(pass string) int {

	var row int
	var column int
	rowLimit := 128
	upperrow := 127
	uppercol := 7
	lowerrow := 0
	lowercol := 0
	for i, c := range pass {
		if i == 7 {
			rowLimit = 8
		}
		rowLimit = rowLimit >> 1
		switch c {
		case 'B':
			if i == 6 {
				row = upperrow
				continue
			}
			lowerrow = lowerrow + rowLimit
		case 'F':
			if i == 6 {
				row = lowerrow
				continue
			}
			upperrow = upperrow - rowLimit
		case 'L':
			uppercol = uppercol - rowLimit
			if i == 9 {
				column = uppercol
				continue
			}
		case 'R':
			lowercol = lowercol + rowLimit
			if i == 9 {
				column = lowercol
				continue
			}
		}
	}
	return getSeatID(row, column)

}

func partOne(boardingPasses []string) int {
	var highestID int

	for _, pass := range boardingPasses {
		id := trverseRowCol(pass)
		if id > highestID {
			highestID = id
		}
	}

	return highestID
}

func partTwo(boardingPasses []string) int {

	ids := make([]int, len(boardingPasses))

	for i, pass := range boardingPasses {
		ids[i] = trverseRowCol(pass)
	}

	sort.Ints(ids[:])

	myID := 0

	previousID := 8
	for _, id := range ids {
		if previousID != id {
			myID = previousID
		}
		previousID = id + 1
	}

	return myID

}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day05.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(getInputAsStringSlice(data))
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Since(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(getInputAsStringSlice(data))
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Since(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Since(totalTimeStart))

}
