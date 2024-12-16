package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	Game "github.com/pouyio/advent-of-code-2024/day-14/models"
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

	re, err := regexp.Compile(`p=(-*\d+),(-*\d+) v=(-*\d+),(-*\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			px, _ := strconv.Atoi(match[1])
			py, _ := strconv.Atoi(match[2])
			vx, _ := strconv.Atoi(match[3])
			vy, _ := strconv.Atoi(match[4])
			state.AddRobot(*Game.NewRobot(px, py, vx, vy))
		}
	}

	// LOGIC

	for _, robot := range state.Robots {
		state.MoveSeconds(&robot, 100)
	}
	fmt.Println(state.SafeFactor())
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

	re, err := regexp.Compile(`p=(-*\d+),(-*\d+) v=(-*\d+),(-*\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			px, _ := strconv.Atoi(match[1])
			py, _ := strconv.Atoi(match[2])
			vx, _ := strconv.Atoi(match[3])
			vy, _ := strconv.Atoi(match[4])
			state.AddRobot(*Game.NewRobot(px, py, vx, vy))
		}
	}

	// LOGIC
	for second := 0; ; second++ {
		if second%1000 == 0 {
			fmt.Println(second)
		}
		positions := state.MoveRobots(second)

		if state.CouldBeTree(positions) {
			fmt.Println(second)
			state.DrawPositions(positions)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
