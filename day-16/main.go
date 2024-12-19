package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	Game "github.com/pouyio/advent-of-code-2024/day-16/models"
	Game1 "github.com/pouyio/advent-of-code-2024/day-16/modesl1"
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

	cost := Game.ShortestPathWithLeastTurns(state.Map, state.GetStartPosition(), state.GetEndPosition())
	fmt.Printf("Total score: %d\n", cost)
}

func part2() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	maze := Game1.ParseMaze(string(content))
	optimalTiles := Game1.FindOptimalPaths(maze)

	count := 0
	for _, isOptimal := range optimalTiles {
		if isOptimal {
			count++
		}
	}

	fmt.Printf("Number of tiles on optimal paths: %d\n", count)
	fmt.Println("\nVisualized optimal paths:")
	fmt.Print(Game1.VisualizeOptimalPaths(maze, optimalTiles))
}

func main() {
	// chatGPT: couldnt make part 2 work
	part1()
	// claude: takes long but works
	part2()
}
