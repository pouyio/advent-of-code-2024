package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	Game "github.com/pouyio/advent-of-code-2024/day-15/models"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reVerticalLimits, err := regexp.Compile(`^##+#$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reMap, err := regexp.Compile(`^#\S+#$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reMovements, err := regexp.Compile(`[<v>\^]+`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	state := Game.NewState()

	for scanner.Scan() {
		matchesVerticalLimits := reVerticalLimits.FindAllStringSubmatch(scanner.Text(), -1)
		matchesMap := reMap.FindAllStringSubmatch(scanner.Text(), -1)
		matchesMovements := reMovements.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matchesVerticalLimits) > 0 {
			state.AddMapRow(matchesVerticalLimits[0][0])
		} else if len(matchesMap) > 0 {
			state.AddMapRow(matchesMap[0][0])
		} else if len(matchesMovements) > 0 {
			state.Movements = state.Movements + strings.Join(matchesMovements[0], "")
		}
	}

	// LOGIC
	for _, movement := range state.Movements {
		state.ConsumeMovement(string(movement))
		// state.DrawMap(state.Map)
	}

	fmt.Println("GPS:", state.CalculateBoxesGPS())
}

func part2() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reVerticalLimits, err := regexp.Compile(`^##+#$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reMap, err := regexp.Compile(`^#\S+#$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reMovements, err := regexp.Compile(`[<v>\^]+`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	state := Game.NewState()

	for scanner.Scan() {
		matchesVerticalLimits := reVerticalLimits.FindAllStringSubmatch(scanner.Text(), -1)
		matchesMap := reMap.FindAllStringSubmatch(scanner.Text(), -1)
		matchesMovements := reMovements.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matchesVerticalLimits) > 0 {
			state.AddMapRow1(matchesVerticalLimits[0][0])
		} else if len(matchesMap) > 0 {
			state.AddMapRow1(matchesMap[0][0])
		} else if len(matchesMovements) > 0 {
			state.Movements = state.Movements + strings.Join(matchesMovements[0], "")
		}
	}

	// LOGIC
	for _, movement := range state.Movements {
		// fmt.Println("Index:", index, "movement", string(movement))
		state.ConsumeMovement1(string(movement))
		// state.DrawMap(state.Map)
		// time.Sleep(70000000)
	}

	fmt.Println("GPS:", state.CalculateBoxesGPS1())
}

func main() {
	part1()
	part2()
}
