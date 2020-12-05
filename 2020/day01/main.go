package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	totalTime := time.Now()

	data, err := ioutil.ReadFile("day01.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	d, err := getInputAsIntSlice(data)
	if err != nil {
		panic(fmt.Sprintf("Could not get slice information: {%s}", err))
	}

	partOneTimeBegin := time.Now()

	var partOne int
	var found bool
	for i := 0; i < len(d); i++ {
		for _, expense := range d {
			if d[i]+expense == 2020 {
				partOne = d[i] * expense
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	partOneTime := time.Since(partOneTimeBegin)

	fmt.Println("Part one:", partOne)

	partTwoTimeBegin := time.Now()

	var partTwo int
	found = false
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d); j++ {
			for _, expense := range d {
				if d[i]+d[j]+expense == 2020 {
					partTwo = d[i] * d[j] * expense
					found = true
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	partTwoTime := time.Since(partTwoTimeBegin)
	endTime := time.Since(totalTime)

	fmt.Println("Part two:", partTwo)
	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", endTime)

}
