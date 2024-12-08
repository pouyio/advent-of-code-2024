package path

import (
	"strings"
)

type guard1 struct {
	OriginalMap      [][]string
	visitedPositions [][]string
	currentPosition  [2]int
	currentDirection string
}

func NewGuard(originalMap [][]string) *guard1 {
	visitedPositions := make([][]string, len(originalMap))
	currentPosition := [2]int{}
	currentDirection := ""

	for i := range originalMap {
		for j, value := range originalMap[i] {
			if value != "." && value != "#" {
				currentPosition = [2]int{i, j}
				currentDirection = value
			}
			visitedPositions[i] = append(visitedPositions[i], "")
		}
	}

	return &guard1{OriginalMap: originalMap, visitedPositions: visitedPositions, currentPosition: currentPosition, currentDirection: currentDirection}
}

func (g *guard1) GetCurrentPosition() (int, int) {
	return g.currentPosition[0], g.currentPosition[1]
}

func (g *guard1) ResetVisited() {
	visitedPositions := make([][]string, len(g.OriginalMap))

	for i := range g.OriginalMap {
		for j, value := range g.OriginalMap[i] {
			if value != "." && value != "#" && len(value) > 0 {
				g.currentPosition = [2]int{i, j}
				g.currentDirection = value
			}
			visitedPositions[i] = append(visitedPositions[i], "")
		}
	}
	g.visitedPositions = visitedPositions
}

func (g *guard1) IsPositionFree(i, j int) bool {
	return g.OriginalMap[i][j] != "#"
}

func (g *guard1) isPositionVisited(i, j int, direction string) bool {
	// 10 is magic number, I need to move on, if I have crossed that path
	// more times its probably a loop
	return strings.Contains(g.visitedPositions[i][j], direction) || len(g.visitedPositions[i][j]) > 10
}

func (g *guard1) isOutsideOfMap(i, j int) bool {
	return i < 0 || i > len(g.OriginalMap)-1 || j < 0 || j > len(g.OriginalMap[0])-1
}

func (g *guard1) rotate() {
	if g.currentDirection == "^" {
		g.currentDirection = ">"
	} else if g.currentDirection == "v" {
		g.currentDirection = "<"
	} else if g.currentDirection == "<" {
		g.currentDirection = "^"
	} else if g.currentDirection == ">" {
		g.currentDirection = "v"
	}
}

func (g *guard1) GoToNextPosition() bool {
	currentI, currentJ := g.currentPosition[0], g.currentPosition[1]
	nextI, nextJ := currentI, currentJ

	if g.currentDirection == "^" {
		nextI = currentI - 1
		if g.IsPositionFree(nextI, nextJ) {
			if g.isPositionVisited(nextI, nextJ, "↑") {
				return true
			}
			g.visitedPositions[currentI][currentJ] += "↑"
			g.currentPosition[0], g.currentPosition[1] = nextI, nextJ
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
		}
		return false
	} else if g.currentDirection == ">" {
		nextJ = currentJ + 1
		if g.IsPositionFree(nextI, nextJ) {
			if g.isPositionVisited(nextI, nextJ, "→") {
				return true
			}
			g.visitedPositions[currentI][currentJ] += "→"
			g.currentPosition[0], g.currentPosition[1] = nextI, nextJ
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
		}
		return false
	} else if g.currentDirection == "<" {
		nextJ = currentJ - 1
		if g.IsPositionFree(nextI, nextJ) {
			if g.isPositionVisited(nextI, nextJ, "←") {
				return true
			}
			g.visitedPositions[currentI][currentJ] += "←"
			g.currentPosition[0], g.currentPosition[1] = nextI, nextJ
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
		}
		return false
	} else if g.currentDirection == "v" {
		nextI = currentI + 1
		if g.IsPositionFree(nextI, nextJ) {
			if g.isPositionVisited(nextI, nextJ, "↓") {
				return true
			}
			g.visitedPositions[currentI][currentJ] += "↓"
			g.currentPosition[0], g.currentPosition[1] = nextI, nextJ
		} else {
			g.visitedPositions[nextI][nextJ] = "#"
			g.rotate()
		}
	}
	return false
}

func (g *guard1) IsLastPosition() bool {
	currentI, currentJ := g.currentPosition[0], g.currentPosition[1]
	nextI, nextJ := 0, 0
	if g.currentDirection == "^" {
		nextI = currentI - 1
	} else if g.currentDirection == ">" {
		nextJ = currentJ + 1
	} else if g.currentDirection == "<" {
		nextJ = currentJ - 1
	} else if g.currentDirection == "v" {
		nextI = currentI + 1
	}
	return g.isOutsideOfMap(nextI, nextJ)
}

func (g *guard1) CountSteps() int {
	sum := 0
	for _, row := range g.visitedPositions {
		for _, value := range row {
			if strings.Contains(value, "↑") || strings.Contains(value, "→") || strings.Contains(value, "↓") || strings.Contains(value, "←") {
				sum++
			}
		}
	}
	return sum + 1
}

func (g *guard1) CreatesLoop() bool {
	for !g.IsLastPosition() {
		if g.GoToNextPosition() {
			return true
		}
	}
	return false
}
