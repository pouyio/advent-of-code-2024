package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	Game "github.com/pouyio/advent-of-code-2024/day-13/models"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	state := Game.NewState()

	reA, err := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reB, err := regexp.Compile(`Button\sB:\sX\+(\d+),\sY\+(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	rePrize, err := regexp.Compile(`Prize\:\sX=(\d+),\sY=(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	game := Game.ClawMachine{Ax: 0, Ay: 0, Bx: 0, By: 0, Px: 0, Py: 0}

	for scanner.Scan() {
		matchesA := reA.FindAllStringSubmatch(scanner.Text(), -1)
		matchesB := reB.FindAllStringSubmatch(scanner.Text(), -1)
		matchesPrize := rePrize.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matchesA) > 0 {
			for _, match := range matchesA {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Ax = x
				game.Ay = y
			}
		} else if len(matchesB) > 0 {
			for _, match := range matchesB {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Bx = x
				game.By = y
			}
		} else if len(matchesPrize) > 0 {
			for _, match := range matchesPrize {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Px = x
				game.Py = y
				state.AddGame(game)
			}
		}
	}

	// LOGIC
	acc := 0
	for _, game := range state.Games {
		times, valid := state.MinimizeCost(game)
		if valid {
			acc += times
		}
	}
	fmt.Println("Times pressed:", acc)
}

func part2() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	state := Game.NewState()

	reA, err := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reB, err := regexp.Compile(`Button\sB:\sX\+(\d+),\sY\+(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	rePrize, err := regexp.Compile(`Prize\:\sX=(\d+),\sY=(\d+)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	game := Game.ClawMachine{Ax: 0, Ay: 0, Bx: 0, By: 0, Px: 0, Py: 0}

	for scanner.Scan() {
		matchesA := reA.FindAllStringSubmatch(scanner.Text(), -1)
		matchesB := reB.FindAllStringSubmatch(scanner.Text(), -1)
		matchesPrize := rePrize.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matchesA) > 0 {
			for _, match := range matchesA {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Ax = x
				game.Ay = y
			}
		} else if len(matchesB) > 0 {
			for _, match := range matchesB {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Bx = x
				game.By = y
			}
		} else if len(matchesPrize) > 0 {
			for _, match := range matchesPrize {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				game.Px = x + 10000000000000
				game.Py = y + 10000000000000
				state.AddGame(game)
			}
		}
	}

	// LOGIC
	acc := 0
	for _, game := range state.Games {
		times, valid := state.SolveMachine(game)
		if valid {
			acc += times
		}

	}
	fmt.Println("Times pressed:", acc)
}

func main() {
	part1()
	part2()
}
