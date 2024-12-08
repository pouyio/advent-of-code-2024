package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	guard "github.com/pouyio/advent-of-code-2024/day-6/guard"
	guard1 "github.com/pouyio/advent-of-code-2024/day-6/guard1"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	originalMap := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		originalMap = append(originalMap, strings.Split(row, ""))
	}

	myGuard := guard.NewGuard(originalMap)

	for !myGuard.IsLastPosition() {
		myGuard.GoToNextPosition()
		myGuard.Print()
	}

	fmt.Println("Total steps:", myGuard.CountSteps())
}

func part2() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	originalMap := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		originalMap = append(originalMap, strings.Split(row, ""))
	}

	myGuard := guard1.NewGuard(originalMap)

	sum := 0
	for i, row := range myGuard.OriginalMap {
		for j := range row {
			cursorI, cursorJ := myGuard.GetCurrentPosition()
			if cursorI == i && cursorJ == j {
				continue
			}
			fmt.Println("Checking position: ", i, j)
			if myGuard.IsPositionFree(i, j) {
				myGuard.OriginalMap[i][j] = "#"
				if myGuard.CreatesLoop() {
					sum++
					myGuard.OriginalMap[i][j] = ""
					myGuard.ResetVisited()
					continue
				}
				myGuard.OriginalMap[i][j] = ""
				myGuard.ResetVisited()
			}
		}
	}

	fmt.Println("Different positions: ", sum)
}

func main() {
	part1()
	part2()
}
