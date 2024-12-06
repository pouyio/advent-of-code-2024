package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calculateMul(digits [][]string) int {
	digitsSingle := digits[0][1:len(digits[0])]
	digit1, _ := strconv.Atoi(digitsSingle[0])
	digit2, _ := strconv.Atoi(digitsSingle[1])
	return digit1 * digit2
}

func part1() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)

	re, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	matches := re.FindAllString(content, -1)

	re, err = regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	sum := 0
	for _, match := range matches {
		digits := re.FindAllStringSubmatch(match, -1)
		sum += calculateMul(digits)
	}

	fmt.Println("Sum is: ", sum)
}

func part2() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)

	re, err := regexp.Compile(`(mul\(\d+,\d+\))|(don\'t\(\))|(do\(\))`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	allMatches := re.FindAllString(content, -1)

	mustExecute := true
	validMatches := []string{}
	for _, value := range allMatches {
		if value == "do()" {
			mustExecute = true
			continue
		}
		if value == "don't()" {
			mustExecute = false
			continue
		}
		if mustExecute {
			validMatches = append(validMatches, value)
		}
	}

	sum := 0
	for _, match := range validMatches {
		re, err = regexp.Compile(`mul\((\d+),(\d+)\)`)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			return
		}
		matches := re.FindAllString(match, -1)

		for _, match := range matches {
			digits := re.FindAllStringSubmatch(match, -1)
			sum += calculateMul(digits)
		}
	}

	fmt.Println("Sum is: ", sum)
}

func main() {
	part1()
	part2()
}
