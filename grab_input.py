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
    input_path = f"{year}/inputs/day{day.zfill(2)}.txt"
    if not os.path.isfile(input_path):
        Path(f"{year}/inputs").mkdir(parents=True, exist_ok=True)
        f = open(input_path, "wb")
        data = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies={'session': os.environ['AOC_SESS']}).content
        f.write(data)

def create_template(year, day):
    """
    creates the golang template
    """
    util_path = f"{year}/util.go"
    if not os.path.isfile(util_path):
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
        Path(f"{year}").mkdir(parents=True, exist_ok=True)
        f = open(file_path, 'w')
        f.write("""
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile(\""""+year+"""/inputs/day"""+day.zfill(2)+""".txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

}
""")
    


parser = argparse.ArgumentParser()
parser.add_argument('-y', '--year', help="year", default=today.strftime("%Y"))
parser.add_argument('-d', '--day', help="day", default=today.strftime("%d"))

args = parser.parse_args()


def main():
    grab_input(args.year, args.day)
    create_template(args.year, args.day)

if __name__ == '__main__':
    main()