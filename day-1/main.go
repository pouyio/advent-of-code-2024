package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	var locationArr1 []int
	var locationArr2 []int
	distanceAcc := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		locationIDs := strings.Fields(scanner.Text())

		locationID1, err1 := strconv.Atoi(locationIDs[0])
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		locationArr1 = append(locationArr1, locationID1)

		locationID2, err2 := strconv.Atoi(locationIDs[1])
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		locationArr2 = append(locationArr2, locationID2)

		sort.Slice(locationArr1, func(i, j int) bool {
			return locationArr1[i] < locationArr1[j]
		})
		sort.Slice(locationArr2, func(i, j int) bool {
			return locationArr2[i] < locationArr2[j]
		})
	}
	for i, v := range locationArr1 {
		distance := v - locationArr2[i]
		if distance < 0 {
			distance = -distance
		}
		distanceAcc = distanceAcc + distance
	}

	fmt.Println("Total distance is: ", distanceAcc)
}

type state struct {
	ids map[string]int
	sum map[string]int
}

func (s *state) similarities() int {
	similaritiesAcc := 0
	for id, occurrences := range s.ids {
		similaritiesAcc = similaritiesAcc + s.sum[id]*occurrences
	}
	return similaritiesAcc
}

func calculateSum(locations *[]int, locationID string) int {
	locationIDNumber, err := strconv.Atoi(locationID)
	accumulator := 0
	if err != nil {
		fmt.Println(err)
		return 0
	}
	for _, value := range *locations {
		if value == locationIDNumber {
			accumulator++
		}
	}
	return accumulator * locationIDNumber
}

func part2() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	myState := state{ids: make(map[string]int), sum: make(map[string]int)}

	var locationArr1 []string
	var locationArr2 []int
	similarityMap := make(map[string]int)
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		locationIDs := strings.Fields(scanner.Text())

		locationArr1 = append(locationArr1, locationIDs[0])

		locationID2, err2 := strconv.Atoi(locationIDs[1])
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		locationArr2 = append(locationArr2, locationID2)
	}

	for _, locationID := range locationArr1 {
		if _, ok := similarityMap[locationID]; ok {
			myState.ids[locationID]++
		} else {
			myState.ids[locationID] = 1
			myState.sum[locationID] = calculateSum(&locationArr2, locationID)
		}
	}

	fmt.Println("Similarity score: ", myState.similarities())
}

func main() {
	part1()
	part2()
}
