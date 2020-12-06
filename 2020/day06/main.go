package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func partOne(ans []string) int {
	var totalUnqiueQuestionsAnsweredPerGroup int

	answers := make(map[rune]struct{})
	for _, s := range ans {
		if len(s) == 0 {
			totalUnqiueQuestionsAnsweredPerGroup += len(answers)
			answers = make(map[rune]struct{})
			continue
		}

		for _, a := range s {
			if _, ok := answers[a]; !ok {
				answers[a] = struct{}{}
			}
		}
	}

	return totalUnqiueQuestionsAnsweredPerGroup
}

func partTwo(ans []string) int {
	var totalAgreedUponAnswersPerGroup int

	answers := make(map[rune]int)
	var sizeOfGroup int
	for _, s := range ans {
		if len(s) == 0 {

			for _, v := range answers {
				if v == sizeOfGroup {
					totalAgreedUponAnswersPerGroup++
				}
			}

			answers = make(map[rune]int)
			sizeOfGroup = 0
			continue
		}

		for _, a := range s {
			p, ok := answers[a]
			if !ok {
				answers[a] = 1
			} else {
				answers[a] = p + 1
			}
		}

		sizeOfGroup++
	}

	return totalAgreedUponAnswersPerGroup
}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day06.txt")
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
