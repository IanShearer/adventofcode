package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func partOne(instructions []string) int {

	var accumulator int
	seenInstruction := make(map[int]struct{})
	for i := 0; i < len(instructions); i++ {

		if _, ok := seenInstruction[i]; !ok {
			seenInstruction[i] = struct{}{}
		} else {
			return accumulator
		}

		code := instructions[i][:3]
		num, _ := strconv.Atoi(instructions[i][4:])

		switch code {
		case "nop":
		case "jmp":
			i += num - 1
		case "acc":
			accumulator += num
		}

	}

	return accumulator
}

func partTwo(instructions []string) int {

	var accumulator int
	for j := 0; j < len(instructions); j++ {
		accumulator = 0

		if instructions[j][:3] == "acc" {
			continue
		}
		copyInstructions := make([]string, len(instructions))
		copy(copyInstructions, instructions)
		inst := copyInstructions[j]
		instNum := inst[4:]
		instCode := inst[:3]
		if instCode == "jmp" {
			copyInstructions[j] = "nop " + instNum
		} else {
			copyInstructions[j] = "jmp " + instNum
		}

		seenInstruction := make(map[int]struct{})
		for i := 0; i < len(instructions); i++ {
			if _, ok := seenInstruction[i]; !ok {
				seenInstruction[i] = struct{}{}
			} else {
				break
			}

			code := copyInstructions[i][:3]
			num, _ := strconv.Atoi(copyInstructions[i][4:])

			switch code {
			case "nop":
			case "jmp":
				i += num - 1
			case "acc":
				accumulator += num
			}
			if i == len(instructions)-1 {
				return accumulator
			}

		}
	}

	return accumulator
}
func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day08.txt")
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
