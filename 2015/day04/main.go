package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func partOne(key []byte) int {

	var secretKey int
	for {
		s := []byte(strconv.Itoa(secretKey))
		n := append(key, s...)
		h := fmt.Sprintf("%x", md5.Sum(n))
		if h[0:5] == "00000" {
			return secretKey
		}
		secretKey++
	}

}

func partTwo(key []byte) int {

	var secretKey int
	for {
		s := []byte(strconv.Itoa(secretKey))
		n := append(key, s...)
		h := fmt.Sprintf("%x", md5.Sum(n))
		if h[0:6] == "000000" {
			return secretKey
		}
		secretKey++
	}

}

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day04.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(data)
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Since(partOneStart)

	partTwoStart := time.Now()
	partTwoAnswer := partTwo(data)
	fmt.Println("Part Two:", partTwoAnswer)
	partTwoTime := time.Since(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Since(totalTimeStart))

}
