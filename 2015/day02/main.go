package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var boxRegex = regexp.MustCompile(`(?m)(\d+)x(\d+)x(\d+)`)

func partOne(dimensions []string) int {
	var totalSquareFeet int
	for _, box := range dimensions {
		if len(box) == 0 {
			continue
		}

		d := boxRegex.FindStringSubmatch(box)
		l, _ := strconv.Atoi(d[1])
		w, _ := strconv.Atoi(d[2])
		h, _ := strconv.Atoi(d[3])

		arr := []int{l * w, w * h, h * l}
		sort.Ints(arr)
		smallest := arr[0]

		totalSquareFeet += (2 * l * w) + (2 * w * h) + (2 * h * l) + smallest

	}

	return totalSquareFeet
}

func partTwo(dimensions []string) int {
	var totalSquareFeet int
	for _, box := range dimensions {
		if len(box) == 0 {
			continue
		}

		d := boxRegex.FindStringSubmatch(box)
		l, _ := strconv.Atoi(d[1])
		w, _ := strconv.Atoi(d[2])
		h, _ := strconv.Atoi(d[3])

		arr := []int{l, w, h}
		sort.Ints(arr)

		totalSquareFeet += (l * w * h) + (arr[0] * 2) + (arr[1] * 2)

	}

	return totalSquareFeet
}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day02.txt")
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
