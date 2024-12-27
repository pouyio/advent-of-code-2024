package main

import (
	"bufio"
	"fmt"
	Game "github.com/pouyio/advent-of-code-2024/day-19/models"
	"os"
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
	state := Game.NewState()

	line := 0
	counter := 0
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			for _, letter := range strings.Split(text, ", ") {
				state.Patterns = append(state.Patterns, letter)
			}
			line++
			continue
		}

		if len(text) > 0 && line > 0 && state.IsValidDesign(text) {
			counter++
		}
	}

	fmt.Println("Total valid designs:", counter)
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
	state := Game.NewState()

	line := 0
	counter := 0
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			for _, letter := range strings.Split(text, ", ") {
				state.Patterns = append(state.Patterns, letter)
			}
			line++
			continue
		}

		if len(text) > 0 && line > 0 {
			counter += state.TotalSolutions(text)
		}
	}

	fmt.Println("Total possible solutions:", counter)
}

func main() {
	part1()
	part2()
}
