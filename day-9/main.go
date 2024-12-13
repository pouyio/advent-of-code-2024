package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	Blocks "github.com/pouyio/advent-of-code-2024/day-9/models"
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
	blocks := Blocks.NewBlocks([]int{})
	for scanner.Scan() {
		for _, letter := range scanner.Text() {
			value, _ := strconv.Atoi(string(letter))
			blocks.Add(value)
		}
	}

	for values := range blocks.Loop() {
		for i := 2; i < values[1]+2; i++ {
			if values[0]%2 == 0 {
				blocks.AddIndividualBlock(values[0] / 2)
			} else {
				blocks.AddIndividualBlock(-1)
			}
		}
	}

OutterLoop:
	for i, value := range blocks.UnrolledBlocks {
		if value < 0 {
			nonEmptyIndex := blocks.FindLastNonEmptyIndex()
			if nonEmptyIndex <= i {
				break OutterLoop
			}
			if nonEmptyIndex != -1 {
				blocks.OrderedBlocks[i] = blocks.UnrolledBlocks[nonEmptyIndex]
				blocks.OrderedBlocks[nonEmptyIndex] = -1
			}
		}
	}

	for i := 0; i < len(blocks.OrderedBlocks); i++ {
		if blocks.OrderedBlocks[i] == -1 {
			// fmt.Print(".")
		} else {
			// fmt.Print(blocks.OrderedBlocks[i])
		}
	}

	// fmt.Println()
	fmt.Println("Checksum:", blocks.CalculateChecksum())
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
	blocks := Blocks.NewBlocks([]int{})
	for scanner.Scan() {
		for _, letter := range scanner.Text() {
			value, _ := strconv.Atoi(string(letter))
			blocks.Add(value)
		}
	}

	for values := range blocks.Loop() {
		for i := 2; i < values[1]+2; i++ {
			if values[0]%2 == 0 {
				blocks.AddIndividualBlock(values[0] / 2)
			} else {
				blocks.AddIndividualBlock(-1)
			}
		}
	}

	// blocks.PrintOrderedBlocks()
	for fileI := len(blocks.OriginalDiskMap) - 1; fileI >= 0; fileI -= 2 {
		fileSize := blocks.OriginalDiskMap[fileI]
		for spaceI := 1; spaceI <= len(blocks.OrderedDiskMap)-1; spaceI += 2 {
			spaceSize := blocks.OrderedDiskMap[spaceI]
			if spaceSize >= fileSize {

				fileStartIndex := blocks.CalculateDiskIndex(fileI)
				spaceStartIndex := blocks.CalculateDiskIndex(spaceI)
				// fmt.Printf("move file with value %d (size: %d, start index: %d) to space start index: %d \n", fileI/2, fileSize, fileStartIndex, spaceStartIndex)
				blocks.OrderedDiskMap[spaceI] -= fileSize
				blocks.OrderedDiskMap[spaceI-1] += fileSize
				blocks.OrderedDiskMap[fileI] = 0

				for i := 0; i < fileSize; i++ {
					blocks.OrderedBlocks[fileStartIndex+i] = -1
					blocks.OrderedBlocks[spaceStartIndex+i] = fileI / 2
				}

				// blocks.PrintOrderedBlocks()
				break
			}
		}
	}

	fmt.Println("Checksum:", blocks.CalculateChecksum())
}

func main() {
	part1()
	part2()
}
