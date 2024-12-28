package main

import (
	"bufio"
	"fmt"
	"os"

	Game "github.com/pouyio/advent-of-code-2024/day-20/models"
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

	for scanner.Scan() {
		state.AddMapRow(scanner.Text())
	}

	state.FillPathTime()
	cheats := state.FindCheats()
	fmt.Println(len(cheats))
}

func part2() {
	// Open the file for reading
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	state := Game.NewState()

	for scanner.Scan() {
		state.AddMapRow(scanner.Text())
	}

	state.FillPathTime()
	cheats := state.FindCheats1()
	fmt.Println(len(cheats))
}

func main() {
	part1()
	// part2()
}
