package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	Game "github.com/pouyio/advent-of-code-2024/day-18/models"
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

	re, err := regexp.Compile(`(\d+),(\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		px, _ := strconv.Atoi(matches[0][1])
		py, _ := strconv.Atoi(matches[0][2])
		state.AddCorruptedCell(px, py)
		if len(state.Corrupted) == 1024 {
			break
		}
	}

	cost := Game.ShortestPath(state.Corrupted, [2]int{0, 0}, [2]int{Game.SIZE, Game.SIZE})
	fmt.Printf("Total score: %d\n", cost)
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

	re, err := regexp.Compile(`(\d+),(\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		px, _ := strconv.Atoi(matches[0][1])
		py, _ := strconv.Atoi(matches[0][2])
		state.AddCorruptedCell(px, py)

		cost := Game.ShortestPath(state.Corrupted, [2]int{0, 0}, [2]int{Game.SIZE, Game.SIZE})
		if cost == -1 {
			fmt.Println("No solution on:", px, py)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
