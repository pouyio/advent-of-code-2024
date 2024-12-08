package guard

import "fmt"

type guard struct {
	originalMap      [][]string
	visitedPositions [][]string
}

func NewGuard(originalMap [][]string) *guard {
	visitedPositions := make([][]string, len(originalMap))

	for i := range originalMap {
		for _, value := range originalMap[i] {
			if value != "." && value != "#" {
				visitedPositions[i] = append(visitedPositions[i], value)
			} else {
				visitedPositions[i] = append(visitedPositions[i], "")
			}
		}
	}

	return &guard{originalMap: originalMap, visitedPositions: visitedPositions}
}

func (g *guard) getCurrentPosition() (int, int) {
	currentI, currentJ := 0, 0
	for i := range g.visitedPositions {
		for j := range g.visitedPositions[i] {
			if g.visitedPositions[i][j] == "<" || g.visitedPositions[i][j] == ">" || g.visitedPositions[i][j] == "v" || g.visitedPositions[i][j] == "^" {
				currentI = i
				currentJ = j
			}
		}
	}
	return currentI, currentJ
}

func (g *guard) isPositionFree(i, j int) bool {
	return g.originalMap[i][j] != "#"
}

func (g *guard) isOutsideOfMap(i, j int) bool {
	return i < 0 || i > len(g.originalMap)-1 || j < 0 || j > len(g.originalMap[0])
}

func (g *guard) rotate() {
	i, j := g.getCurrentPosition()
	currentValue := g.visitedPositions[i][j]
	if currentValue == "^" {
		g.visitedPositions[i][j] = ">"
	} else if currentValue == "v" {
		g.visitedPositions[i][j] = "<"
	} else if currentValue == "<" {
		g.visitedPositions[i][j] = "^"
	} else if currentValue == ">" {
		g.visitedPositions[i][j] = "v"
	}
}

func (g *guard) GoToNextPosition() {
	currentI, currentJ := g.getCurrentPosition()
	currentValue := g.visitedPositions[currentI][currentJ]
	nextI, nextJ := currentI, currentJ

	if currentValue == "^" {
		nextI = currentI - 1
		if g.isOutsideOfMap(nextI, nextJ) {
			return
		}
		if g.isPositionFree(nextI, nextJ) {
			g.visitedPositions[currentI][currentJ] = "x"
			g.visitedPositions[nextI][nextJ] = "^"
			return
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
			return
		}
	} else if currentValue == ">" {
		nextJ = currentJ + 1
		if g.isOutsideOfMap(nextI, nextJ) {
			return
		}
		if g.isPositionFree(nextI, nextJ) {
			g.visitedPositions[currentI][currentJ] = "x"
			g.visitedPositions[nextI][nextJ] = ">"
			return
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
			return
		}
	} else if currentValue == "<" {
		nextJ = currentJ - 1
		if g.isOutsideOfMap(nextI, nextJ) {
			return
		}
		if g.isPositionFree(nextI, nextJ) {
			g.visitedPositions[currentI][currentJ] = "x"
			g.visitedPositions[nextI][nextJ] = "<"
			return
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
			return
		}
	} else if currentValue == "v" {
		nextI = currentI + 1
		if g.isOutsideOfMap(nextI, nextJ) {
			return
		}
		if g.isPositionFree(nextI, nextJ) {
			g.visitedPositions[currentI][currentJ] = "x"
			g.visitedPositions[nextI][nextJ] = "v"
			return
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
			return
		}

	}

}

func (g *guard) Print() {
	for _, row := range g.visitedPositions {
		for _, value := range row {
			if value != "" {
				fmt.Print(value)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------")
}

func (g *guard) IsLastPosition() bool {
	currentI, currentJ := g.getCurrentPosition()
	nextI, nextJ := 0, 0
	if g.visitedPositions[currentI][currentJ] == "^" {
		nextI = currentI - 1
	} else if g.visitedPositions[currentI][currentJ] == ">" {
		nextJ = currentJ + 1
	} else if g.visitedPositions[currentI][currentJ] == "<" {
		nextJ = currentJ - 1
	} else if g.visitedPositions[currentI][currentJ] == "v" {
		nextI = currentI + 1

	}
	return g.isOutsideOfMap(nextI, nextJ)
}

func (g *guard) CountSteps() int {
	sum := 0
	for _, row := range g.visitedPositions {
		for _, value := range row {
			if value == "x" {
				sum++
			}
		}
	}
	return sum + 1
}
