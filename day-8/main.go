package main

import (
	"bufio"
	"fmt"
	AtennaMap "github.com/pouyio/advent-of-code-2024/day-8/models"
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
	originalMap := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		originalMap = append(originalMap, strings.Split(row, ""))
	}
	antennaMap := AtennaMap.NewAntennaMap(originalMap)

	// extract pairs of antennas
	for indexI, row := range antennaMap.Map {
		for indexJ, value := range row {
			if value != "." {
				antennaMap.AddAntennaPair(value, indexI, indexJ)
			}
		}
	}

	for key, pairs := range antennaMap.AntennaPairs {
		for index := 0; index <= len(pairs)-1; index++ {
			for nextIndex := index + 1; nextIndex <= len(pairs)-1; nextIndex++ {
				position1 := pairs[index]
				position2 := pairs[nextIndex]
				antennaMap.CalculateAntinodes(key, position1, position2)
			}
		}
	}

	for _, row := range antennaMap.AntinodesMap {
		fmt.Println(row)
	}
	fmt.Println("Antinodes: ", antennaMap.Antinodes)
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
	originalMap := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		originalMap = append(originalMap, strings.Split(row, ""))
	}
	antennaMap := AtennaMap.NewAntennaMap(originalMap)

	// extract pairs of antennas
	for indexI, row := range antennaMap.Map {
		for indexJ, value := range row {
			if value != "." {
				antennaMap.AddAntennaPair(value, indexI, indexJ)
			}
		}
	}

	for key, pairs := range antennaMap.AntennaPairs {
		for index := 0; index <= len(pairs)-1; index++ {
			for nextIndex := index + 1; nextIndex <= len(pairs)-1; nextIndex++ {
				position1 := pairs[index]
				position2 := pairs[nextIndex]
				antennaMap.CalculateAntinodes1(key, position1, position2)
			}
		}
	}

	for _, row := range antennaMap.AntinodesMap {
		fmt.Println(row)
	}
	fmt.Println("Antinodes: ", antennaMap.Antinodes)
}

func main() {
	// part1()
	part2()
}
