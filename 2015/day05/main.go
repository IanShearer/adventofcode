package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var badPairs = []string{
	"ab",
	"cd",
	"pq",
	"xy",
}

func checkIfVowel(c byte) bool {
	return c == 'a' || c == 'e' ||
		c == 'i' || c == 'o' || c == 'u'
}

func checkIfBadPair(x, y byte) bool {
	var s strings.Builder
	s.WriteByte(x)
	s.WriteByte(y)

	str := s.String()
	for _, bp := range badPairs {
		if str == bp {
			return true
		}
	}
	return false
}

func partOne(str []string) int {

	var niceStrings int
	for _, s := range str {
		var vowelCount int
		var doubleCharacter bool
		var badPairPresent bool

		for i := range s {
			if checkIfVowel(s[i]) {
				vowelCount++
			}

			if i == 0 {
				continue
			}

			if s[i-1] == s[i] {
				doubleCharacter = true
			}

			if checkIfBadPair(s[i-1], s[i]) {
				badPairPresent = true
			}

		}

		if vowelCount >= 3 && doubleCharacter && !badPairPresent {
			niceStrings++
		}
	}
	return niceStrings
}

func checkLetterBetween(f, s, t byte) bool {
	return f == t && f != s
}

func partTwo(str []string) int {

	var niceStrings int
	for _, s := range str {
		var lettersBetween bool
		var hasTwoPairs bool
		var overlap bool

		l := len(s)
		pairs := make(map[string]int)
		for i := range s {
			// dont need to check last case
			if l-i == 1 {
				continue
			}

			if l-i > 2 {
				if checkLetterBetween(s[i], s[i+1], s[i+2]) {
					lettersBetween = true
				}

				if s[i] == s[i+1] && s[i] == s[i+2] {
					var b strings.Builder
					b.WriteByte(s[i])
					b.WriteByte(s[i+1])
					if p, ok := pairs[b.String()]; ok {
						pairs[b.String()] = p - 1
					} else {
						pairs[b.String()] = 0
					}
				}
			}

			var b strings.Builder
			b.WriteByte(s[i])
			b.WriteByte(s[i+1])
			if p, ok := pairs[b.String()]; ok {
				pairs[b.String()] = p + 1
			} else {
				pairs[b.String()] = 1
			}

		}

		if !overlap {
			for _, r := range pairs {
				if r == 2 {
					hasTwoPairs = true
				} else if r > 2 {
					hasTwoPairs = false
					break
				}
			}
		}

		if hasTwoPairs && lettersBetween {
			niceStrings++
		}
	}

	return niceStrings

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
