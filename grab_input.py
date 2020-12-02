#! /usr/bin/python3

import requests
import os
import argparse
import datetime
from pathlib import Path

today = datetime.date.today() 

def grab_input(year, day):
    """
    grabs the input from adventofcode if the data does not already exists
    """
    print("checking input")
    input_path = f"{year}/inputs/day{day.zfill(2)}.txt"
    if not os.path.isfile(input_path):
        print("getting data...")
        Path(f"{year}/inputs").mkdir(parents=True, exist_ok=True)
        f = open(input_path, "wb")
        data = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies={'session': os.environ['AOC_SESS']}).content
        f.write(data)
    print("finished getting input")

def create_template(year, day):
    """
    creates the golang template
    """
    util_path = f"{year}/util.go"
    if not os.path.isfile(util_path):
        print("creating util file...")
        Path(f"{year}").mkdir(parents=True, exist_ok=True)
        f = open(util_path, 'w')
        f.write("""package main

import (
	"strconv"
	"strings"
)

func getInputAsIntSlice(d []byte) ([]int, error) {
	data := strings.Split(string(d), "\\n")
	numbers := make([]int, len(data)-1)
	var err error
	for i, s := range data {
		if len(s) == 0 {
			continue
		}
		numbers[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}

	return numbers, nil
}

func getInputAsStringSlice(d []byte) []string {
	return strings.Split(string(d), "\\n")
}
""")
    file_path = f"{year}/day{day.zfill(2)}.go"
    if not os.path.isfile(file_path):
        print("creating file of the day...")
        Path(f"{year}").mkdir(parents=True, exist_ok=True)
        f = open(file_path, 'w')
        f.write("""
package main

import (
	"fmt"
	"io/ioutil"
)

func partOne(someInput interface{}) interface{} {

}

// func partTwo(someInput interface{}) interface{} {
//
// }

func main() {

    totalTimeStart := time.Now()

	data, err := ioutil.ReadFile(\""""+year+"""/inputs/day"""+day.zfill(2)+""".txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

    // format your input here

    partOneStart := time.Now()
    partOneAnswer := partOne()
    fmt.Println("Part One:", partOneAnswer)
    partOneTime := time.Now().Sub(partOneStart)

    // partTwoStart := time.Now()
    // partTwoAnswer := partTwo()
    // fmt.Println("Part Two:", partTwoAnswer)
    // partTwoTime := time.Now().Sub(partOneStart)


    fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	// fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Now().Sub(totalTimeStart))

}
""")
    print("finished creating files")
    

parser = argparse.ArgumentParser()
parser.add_argument('-y', '--year', help="year", default=today.strftime("%Y"))
parser.add_argument('-d', '--day', help="day", default=today.strftime("%d"))

args = parser.parse_args()


def main():
    grab_input(args.year, args.day)
    create_template(args.year, args.day)

if __name__ == '__main__':
    main()