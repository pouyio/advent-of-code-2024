package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateMul(digits [][]string) int {
	digitsSingle := digits[0][1:len(digits[0])]
	digit1, _ := strconv.Atoi(digitsSingle[0])
	digit2, _ := strconv.Atoi(digitsSingle[1])
	return digit1 * digit2
}

func part1() {
	const Pattern = "XMAS"
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	matrix := [][]string{}
	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		matrix = append(matrix, strings.Split(row, ""))
	}

	occurrences := 0
	for i, row := range matrix {
		for j, word := range row {
			if word == string(Pattern[0]) {
				if matchToRight(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToLeft(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToUp(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToDown(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToDownRight(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToDownLeft(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToUpLeft(Pattern, matrix, i, j) {
					occurrences++
				}
				if matchToUpRight(Pattern, matrix, i, j) {
					occurrences++
				}
			}
		}
	}
	fmt.Println("XMAS found: ", occurrences)
}

func safeAccess(matrix [][]string, i, j int) string {
	if len(matrix) > i && i >= 0 {
		if len(matrix[i]) > j && j >= 0 {
			return matrix[i][j]
		}
	}
	return ""
}

func matchToRight(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i, j+index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}

func matchToLeft(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i, j-index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToUp(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i-index, j)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToDown(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i+index, j)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToDownRight(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i+index, j+index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToDownLeft(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i+index, j-index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToUpLeft(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i-index, j-index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}
func matchToUpRight(pattern string, matrix [][]string, i, j int) bool {
	match := true
	for index := range pattern {
		safeMatrixValue := safeAccess(matrix, i-index, j+index)
		if string(pattern[index]) != safeMatrixValue {
			match = false
			break
		}
	}
	return match
}

func reverse(pattern string) string {
	runes := []rune(pattern)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func matchFromCenter(pattern string, matrix [][]string, i, j int) bool {
	return (matchToDownRight(pattern, matrix, i-1, j-1) || matchToDownRight(reverse(pattern), matrix, i-1, j-1)) && (matchToUpRight(pattern, matrix, i+1, j-1) || matchToUpRight(reverse(pattern), matrix, i+1, j-1))
}

func part2() {
	const Pattern = "MAS"
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	matrix := [][]string{}
	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		matrix = append(matrix, strings.Split(row, ""))
	}

	occurrences := 0
	for i, row := range matrix {
		if i == 0 || i == len(matrix)-1 {
			continue
		}
		for j, word := range row {
			if j == 0 || j == len(matrix[i])-1 {
				continue
			}
			if word == string(Pattern[1]) {
				if matchFromCenter(Pattern, matrix, i, j) {
					occurrences++
				}
			}
		}
	}

	fmt.Println("X-MAS found: ", occurrences)
}

func main() {
	part1()
	part2()
}
