package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	Game "github.com/pouyio/advent-of-code-2024/day-17/models"
)

func part1() {
	// Create a new Scanner for the file
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Failed to read file:", err)
	}

	reA, err := regexp.Compile(`Register A: (\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reB, err := regexp.Compile(`Register B: (\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reC, err := regexp.Compile(`Register C: (\d+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	reP, err := regexp.Compile(`(?m)^Program: (\d+(?:,\d+)*)$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	a, _ := strconv.Atoi(reA.FindAllStringSubmatch(string(content), -1)[0][1])
	b, _ := strconv.Atoi(reB.FindAllStringSubmatch(string(content), -1)[0][1])
	c, _ := strconv.Atoi(reC.FindAllStringSubmatch(string(content), -1)[0][1])
	p := reP.FindAllStringSubmatch(string(content), -1)[0]

	operations := []int{}
	p = strings.Split(p[1], ",")
	for _, program := range p {
		operation, _ := strconv.Atoi(program)
		operations = append(operations, int(operation))
	}

	state := Game.NewState(a, b, c, operations)

	iteration := 0
	for state.Pointer < len(state.OriginalProgram) {
		opcode := state.OriginalProgram[state.Pointer]
		operand := state.OriginalProgram[state.Pointer+1]
		state.Opcodes[opcode](operand)
		iteration++
	}

	var outputs []string
	for _, i := range state.Out {
		outputs = append(outputs, strconv.Itoa(i))
	}

	fmt.Println(strings.Join(outputs, ","))

}

func part2() {
}

func main() {
	part1()
	part2()
}
