package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type bagInfo struct {
	name      string
	otherBags map[string]int
}

func parseBagInformation(b string) bagInfo {

	bag := bagInfo{}

	a := strings.Split(strings.ReplaceAll(b, ".", ""), " ")

	var bagString strings.Builder
	bagString.WriteString(a[0])
	bagString.WriteRune(' ')
	bagString.WriteString(a[1])
	bagString.WriteRune(' ')
	bagString.WriteString(a[2])

	bag.name = bagString.String()

	ob := a[4:]

	if ob[0] == "no" {
		return bag
	}

	childBags := strings.Split(strings.Join(ob, " "), ",")
	for _, bagOption := range childBags {
		s := strings.TrimSpace(bagOption)
		n, _ := strconv.Atoi(string(s[0])) // we dont care for this yet
		typeOfBag := s[2:]
		if bag.otherBags == nil {
			bag.otherBags = make(map[string]int, len(childBags))
		}
		bag.otherBags[typeOfBag] = n
	}

	return bag

}

func partOne(bagInformation []string) int {

	const bagToLookFor = "shiny gold bag"
	const bagToLookForPlural = "shiny gold bags"

	bags := make([]bagInfo, len(bagInformation))
	bagsThatContainCriticalBag := make(map[string]struct{})
	for i, bag := range bagInformation {
		bags[i] = parseBagInformation(bag)
	}

	for _, b := range bags {
		if b.name == bagToLookForPlural {
			continue
		}
		for bagName := range b.otherBags {
			if bagName == bagToLookFor || bagName == bagToLookForPlural {
				if _, ok := bagsThatContainCriticalBag[b.name]; !ok {
					bagsThatContainCriticalBag[b.name] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(bagsThatContainCriticalBag))
	uniqueBags := 1
	for uniqueBags != 0 {
		uniqueBags = 0
		for _, b := range bags {
			for bag := range bagsThatContainCriticalBag {
				for asdf := range b.otherBags {
					// fmt.Println(asdf[len(asdf)-1:], asdf)
					if bag == asdf || bag == asdf+"s" {
						if _, ok := bagsThatContainCriticalBag[b.name]; !ok {
							bagsThatContainCriticalBag[b.name] = struct{}{}
							uniqueBags++
						}
					}
				}
			}
		}
		fmt.Println(len(bagsThatContainCriticalBag))
	}

	return len(bagsThatContainCriticalBag)

}

// func partTwo(someInput interface{}) interface{} {
//
// }

func main() {

	totalTimeStart := time.Now()

	data, err := ioutil.ReadFile("day07.txt")
	if err != nil {
		panic(fmt.Sprintf("There was an issue reading the file: {%s}", err))
	}

	partOneStart := time.Now()
	partOneAnswer := partOne(getInputAsStringSlice(data))
	fmt.Println("Part One:", partOneAnswer)
	partOneTime := time.Since(partOneStart)

	// partTwoStart := time.Now()
	// partTwoAnswer := partTwo()
	// fmt.Println("Part Two:", partTwoAnswer)
	// partTwoTime := time.Since(partTwoStart)

	fmt.Println("------------------")
	fmt.Println("Time for Part One:  ", partOneTime)
	// fmt.Println("Time for Part Two:  ", partTwoTime)
	fmt.Println("Total Time:         ", time.Since(totalTimeStart))

}
