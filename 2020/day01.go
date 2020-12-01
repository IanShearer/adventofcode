package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("2020/inputs/day01.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	d, err := getInputAsIntSlice(data)
	if err != nil {
		panic(fmt.Sprintf("Could not get slice information: {%s}", err))
	}

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

	fmt.Println("Part one:", partOne)

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

	fmt.Println("Part two:", partTwo)

}
