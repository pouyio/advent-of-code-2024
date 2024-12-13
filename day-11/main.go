package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	Pebbles "github.com/pouyio/advent-of-code-2024/day-11/models"
)

func part1() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)

	pebbles := Pebbles.NewPebble()
	for _, value := range strings.Fields(content) {
		pebbleNumber, _ := strconv.Atoi(value)
		pebbles.Add(pebbleNumber)
	}

	for i := 0; i < 25; i++ {
		pebbles.Blink()
	}

	acc := 0
	for _, value := range pebbles.PebblesMap {
		if value > 0 {
			acc += value
		}
	}
	fmt.Println("Stones:", acc)
}

func part2() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)

	pebbles := Pebbles.NewPebble()
	for _, value := range strings.Fields(content) {
		pebbleNumber, _ := strconv.Atoi(value)
		pebbles.Add(pebbleNumber)
	}

	for i := 0; i < 75; i++ {
		pebbles.Blink()
	}

	acc := 0
	for _, value := range pebbles.PebblesMap {
		if value > 0 {
			acc += value
		}
	}
	fmt.Println("Stones:", acc)
}

func main() {
	part1()
	part2()
}
