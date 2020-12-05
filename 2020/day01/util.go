package main

import (
	"strconv"
	"strings"
)

func getInputAsIntSlice(d []byte) ([]int, error) {
	data := strings.Split(string(d), "\n")
	numbers := make([]int, len(data))
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
	return strings.Split(string(d), "\n")
}
