package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func reverseRow(row []string) []string {
	reversedRow := make([]string, len(row))
	for j := range row {
		reversedRow[j] = row[len(row)-j-1]
	}
	return reversedRow
}

func reverseRows(matrix [][]string) [][]string {
	mirroredMatrix := make([][]string, len(matrix))

	for i, row := range matrix {
		mirroredMatrix[i] = reverseRow(row)
	}

	return mirroredMatrix
}

func findSomeRule(value1, value2 string, rules [][]string) bool {
	for _, rule := range rules {
		if rule[0] == value1 && rule[1] == value2 {
			return true
		}
	}
	return false
}

func checkAllNextValues(pageIndex int, update []string, rules [][]string) bool {
	currentValue := update[pageIndex]
	for i := pageIndex + 1; i < len(update)-1; i++ {
		if !findSomeRule(currentValue, update[i], rules) {
			return false
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
	rules := [][]string{}
	inRules := true
	updates := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			inRules = false
			continue
		}
		if inRules {
			rules = append(rules, strings.Split(row, "|"))
		} else {
			updates = append(updates, strings.Split(row, ","))
		}
	}

	validUpdates := [][]string{}
	for _, update := range updates {
		validUpdate := true
		for pageIndex := range update {
			if !(checkAllNextValues(pageIndex, update, rules) && checkAllNextValues(pageIndex, reverseRow(update), reverseRows(rules))) {
				validUpdate = false
				break
			}
		}

		if validUpdate {
			validUpdates = append(validUpdates, update)
		}
	}

	sum := 0
	for _, validUpdate := range validUpdates {
		value, _ := strconv.Atoi(validUpdate[len(validUpdate)/2])
		sum += value
	}

	fmt.Println("Middle page sum: ", sum)
}

func fixUpdate(update []string, rules [][]string) []string {
	sort.Slice(update, func(i, j int) bool {
		return findSomeRule(update[i], update[j], rules)
	})
	return update
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
	rules := [][]string{}
	inRules := true
	updates := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			inRules = false
			continue
		}
		if inRules {
			rules = append(rules, strings.Split(row, "|"))
		} else {
			updates = append(updates, strings.Split(row, ","))
		}
	}

	fixedUpdates := [][]string{}
	for _, update := range updates {
		for pageIndex := range update {
			if !checkAllNextValues(pageIndex, update, rules) || !checkAllNextValues(pageIndex, reverseRow(update), reverseRows(rules)) {
				fixedUpdates = append(fixedUpdates, fixUpdate(update, rules))
				break
			}
		}

	}

	sum := 0
	for _, validUpdate := range fixedUpdates {
		value, _ := strconv.Atoi(validUpdate[len(validUpdate)/2])
		sum += value
	}

	fmt.Println("Middle page sum: ", sum)
}

func main() {
	// part1()
	part2()
}
