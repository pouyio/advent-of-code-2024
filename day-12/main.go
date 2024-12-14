package main

import (
	"bufio"
	"fmt"
	"os"

	Garden "github.com/pouyio/advent-of-code-2024/day-12/models"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	garden := Garden.NewGarden()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		garden.OriginalGarden = append(garden.OriginalGarden, []string{})
		for _, letter := range scanner.Text() {
			garden.OriginalGarden[y] = append(garden.OriginalGarden[y], string(letter))
		}
		y++
	}

	for y, row := range garden.OriginalGarden {
		for x, letter := range row {
			areaIndex := garden.GetMapIndex([2]int{x, y})
			if areaIndex == -1 {
				garden.Areas[letter] = append(garden.Areas[letter], [][2]int{})
				areaIndex = len(garden.Areas[letter]) - 1
				garden.Flood([2]int{x, y}, letter, areaIndex)
			}
		}
	}

	totalPrice := 0
	for _, letters := range garden.Areas {
		for _, areas := range letters {
			area := len(areas)
			perimeter := Garden.CalculatePerimeter(areas)
			totalPrice += area * perimeter
			// fmt.Printf("Price %s = %d * %d = %d\n", key, area, perimeter, area*perimeter)
		}
	}
	fmt.Println("Total price", totalPrice)

}

func part2() {
	// Open the file for reading
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	garden := Garden.NewGarden()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		garden.OriginalGarden = append(garden.OriginalGarden, []string{})
		for _, letter := range scanner.Text() {
			garden.OriginalGarden[y] = append(garden.OriginalGarden[y], string(letter))
		}
		y++
	}

	for y, row := range garden.OriginalGarden {
		for x, letter := range row {
			areaIndex := garden.GetMapIndex([2]int{x, y})
			if areaIndex == -1 {
				garden.Areas[letter] = append(garden.Areas[letter], [][2]int{})
				areaIndex = len(garden.Areas[letter]) - 1
				garden.Flood([2]int{x, y}, letter, areaIndex)
			}
		}
	}

	totalPrice := 0
	for _, letters := range garden.Areas {
		for _, areas := range letters {
			area := len(areas)
			angles := garden.CountCorners(areas)
			totalPrice += area * angles
			// fmt.Printf("Price %s = %d * %d = %d\n", key, area, angles, area*angles)
		}
	}
	fmt.Println("Total price", totalPrice)
}

func main() {
	part1()
	part2()
}
