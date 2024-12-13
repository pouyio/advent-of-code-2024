package main

import (
	"bufio"
	"fmt"
	Map "github.com/pouyio/advent-of-code-2024/day-10/models"
	"os"
	"strconv"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	topoMap := Map.NewMap()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		topoMap.OriginalMap = append(topoMap.OriginalMap, []int{})
		for x, letter := range scanner.Text() {
			value, _ := strconv.Atoi(string(letter))
			if value == 0 {
				topoMap.Trailheads = append(topoMap.Trailheads, [2]int{x, y})
			}
			topoMap.OriginalMap[y] = append(topoMap.OriginalMap[y], value)
		}
		y++
	}

	acc := 0
	for _, trailhead := range topoMap.Trailheads {
		topoMap.CalculateTrailheadScore(trailhead)
		acc += len(topoMap.TrailHeadsScores)
		topoMap.TrailHeadsScores = map[[2]int]bool{}
	}

	fmt.Println("Total:", acc)
}

func part2() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	topoMap := Map.NewMap()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		topoMap.OriginalMap = append(topoMap.OriginalMap, []int{})
		for x, letter := range scanner.Text() {
			value, _ := strconv.Atoi(string(letter))
			if value == 0 {
				topoMap.Trailheads = append(topoMap.Trailheads, [2]int{x, y})
			}
			topoMap.OriginalMap[y] = append(topoMap.OriginalMap[y], value)
		}
		y++
	}

	for _, trailhead := range topoMap.Trailheads {
		topoMap.CalculateTrailheadScore1(trailhead)
	}

	fmt.Println("Total:", topoMap.TrailHeadsScoresRating)
}

func main() {
	part1()
	part2()
}
