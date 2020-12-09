package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

var preAmbleLen = 25

func partOne(input []int) int {

	for i := preAmbleLen; i < len(input); i++ {
		var doesEqual bool
		for j := i - preAmbleLen; j < i; j++ { // previous preamblelen
			for k := i - preAmbleLen; k < i; k++ {
				if input[i] == input[j]+input[k] {
					doesEqual = true
				}
			}
		}
		if !doesEqual {
			return input[i]
		}
	}

	return 0
}

func partTwo(input []int, encriptionAnswer int) int {

	var weakness int
	numbersUsed := make([]int, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-i; j++ {
			weakness += input[j+i]
			numbersUsed = append(numbersUsed, input[j+i])
			if weakness == encriptionAnswer {
				sort.Ints(numbersUsed)
				return numbersUsed[0] + numbersUsed[len(numbersUsed)-1]
			} else if weakness > encriptionAnswer {
				break
			}
		}
		weakness = 0
		numbersUsed = make([]int, 0)
	}

	return 0
}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day09.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}
	v, err := getInputAsIntSlice(data)
	if err != nil {
		panic(fmt.Sprintf("There was an issue with the numbers: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(v)
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Since(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(v, partOneAnswer)
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Since(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Since(totalTimeStart))

}
