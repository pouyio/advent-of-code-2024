package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func transformToOperators1(combination int, maxNumber int) []string {
	maxBase3Length := len(strconv.FormatInt(int64(maxNumber), 3))
	base3Str := fmt.Sprintf("%0*s", maxBase3Length, strconv.FormatInt(int64(combination), 3))
	operators := []string{}

	for _, digit := range base3Str {
		switch digit {
		case '0':
			operators = append(operators, "+")
		case '1':
			operators = append(operators, "*")
		case '2':
			operators = append(operators, "||")
		}
	}

	return operators
}

func calculateTotal1(numbers []int, operations []string) int {
	total := 0
	for index, number := range numbers {
		if index == 0 {
			total = number
			continue
		}
		if operations[index-1] == "+" {
			total += number
		} else if operations[index-1] == "*" {
			total *= number
		} else if operations[index-1] == "||" {
			total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(number))
		}
	}
	return total
}

func transformToOperators(combination int, maxNumber int) []string {
	maxBinaryLength := len(strconv.FormatInt(int64(maxNumber), 2))
	binaryStr := fmt.Sprintf("%0*s", maxBinaryLength, strconv.FormatInt(int64(combination), 2))
	operators := []string{}

	for _, digit := range binaryStr {
		if digit == '0' {
			operators = append(operators, "+")
		} else if digit == '1' {
			operators = append(operators, "*")
		}
	}

	return operators
}

func calculateTotal(numbers []int, operations []string) int {
	total := 0
	for index, number := range numbers {
		if index == 0 {
			total = number
			continue
		}
		if operations[index-1] == "+" {
			total += number
		} else {
			total *= number
		}
	}
	return total
}

func isEquationValid(total int, numbers []int) bool {
	totalCombinations := int(math.Pow(2, float64(len(numbers)-1)))
	for i := 0; i < totalCombinations; i++ {
		operations := transformToOperators(i, totalCombinations-1)
		calculatedTotal := calculateTotal(numbers, operations)
		if calculatedTotal == total {
			fmt.Println("total", total)
			fmt.Println("calculatedTotal", calculatedTotal)
			return true
		}
	}
	return false
}

func isEquationValid1(total int, numbers []int) bool {
	totalCombinations := int(math.Pow(3, float64(len(numbers)-1)))
	for i := 0; i < totalCombinations; i++ {
		operations := transformToOperators1(i, totalCombinations-1)
		calculatedTotal := calculateTotal1(numbers, operations)
		if calculatedTotal == total {
			fmt.Println("total", total)
			fmt.Println("calculatedTotal", calculatedTotal)
			return true
		}
	}
	return false
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
	callibrations := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		intRow := []int{}
		for _, value := range strings.Fields(strings.ReplaceAll(row, ":", " ")) {
			intValue, _ := strconv.Atoi(value)
			intRow = append(intRow, intValue)
		}
		callibrations = append(callibrations, intRow)
	}

	total := 0
	for _, callibration := range callibrations {
		if isEquationValid(callibration[0], callibration[1:]) {
			total += callibration[0]
		}
	}
	fmt.Println("Total sum: ", total)
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
	callibrations := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		intRow := []int{}
		for _, value := range strings.Fields(strings.ReplaceAll(row, ":", " ")) {
			intValue, _ := strconv.Atoi(value)
			intRow = append(intRow, intValue)
		}
		callibrations = append(callibrations, intRow)
	}

	total := 0
	for _, callibration := range callibrations {
		if isEquationValid1(callibration[0], callibration[1:]) {
			total += callibration[0]
		}
	}
	fmt.Println("Total sum: ", total)
}

func main() {
	part1()
	part2()
}
