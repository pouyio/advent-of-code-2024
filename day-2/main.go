package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func IsValidReport(levels []string) bool {
	var levelsTrend float64
	for i, level := range levels {
		currentLevel, err := strconv.Atoi(level)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if i+1 < len(levels) {
			nextLevel, err := strconv.Atoi(levels[i+1])
			if err != nil {
				fmt.Println(err)
				return false
			}

			if i == 0 {
				levelsTrend = float64(currentLevel - nextLevel)
				reportsTrendAbs := math.Abs(levelsTrend)

				if reportsTrendAbs < 1 || reportsTrendAbs > 3 {
					return false
				}
			}

			currentTrend := float64(currentLevel - nextLevel)
			currentTrendAbs := math.Abs(currentTrend)

			if currentTrendAbs < 1 || currentTrendAbs > 3 {
				return false
			} else if levelsTrend/currentTrend < 0 {
				return false
			}
		}
	}
	return true
}

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

	totalUnsafeReports := 0
	totalReports := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		levels := strings.Fields(scanner.Text())
		totalReports++

		isValid := IsValidReport(levels)
		if !isValid {
			totalUnsafeReports++
		}
	}

	fmt.Println("Total safe reports: ", totalReports-totalUnsafeReports)
}

func ExtractElement(levels []string, index int) []string {
	result := make([]string, len(levels)-1)

	copy(result, levels[:index])
	copy(result[index:], levels[index+1:])

	return result
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

	totalUnsafeReports := 0
	totalReports := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		levels := strings.Fields(scanner.Text())
		totalReports++

		isValid := IsValidReport(levels)
		if isValid {
			continue
		} else {
			for i := range levels {
				newLevels := ExtractElement(levels, i)
				isValid := IsValidReport(newLevels)
				if isValid {
					break
				} else if i+1 == len(levels) {
					totalUnsafeReports++
				}
			}
		}

	}

	fmt.Println("Total safe reports: ", totalReports-totalUnsafeReports)
}

func main() {
	part1()
	part2()
}
