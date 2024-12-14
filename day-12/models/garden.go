package Garden

import (
	"slices"
)

type Garden struct {
	OriginalGarden [][]string
	Areas          map[string][][][2]int
}

func NewGarden() *Garden {
	return &Garden{OriginalGarden: [][]string{}, Areas: map[string][][][2]int{}}
}

func (g *Garden) Flood(p [2]int, originalLetter string, areaIndex int) {

	currentLetter := g.OriginalGarden[p[1]][p[0]]
	if currentLetter != originalLetter {
		return
	}

	if g.GetMapIndex([2]int{p[0], p[1]}) > -1 {
		return
	}

	g.Areas[originalLetter][areaIndex] = append(g.Areas[originalLetter][areaIndex], p)

	if g.isNextPositionSameLetter(p, [2]int{0, -1}) {
		g.Flood([2]int{p[0], p[1] - 1}, originalLetter, areaIndex)
	}
	if g.isNextPositionSameLetter(p, [2]int{0, 1}) {
		g.Flood([2]int{p[0], p[1] + 1}, originalLetter, areaIndex)
	}
	if g.isNextPositionSameLetter(p, [2]int{-1, 0}) {
		g.Flood([2]int{p[0] - 1, p[1]}, originalLetter, areaIndex)
	}
	if g.isNextPositionSameLetter(p, [2]int{1, 0}) {
		g.Flood([2]int{p[0] + 1, p[1]}, originalLetter, areaIndex)
	}
}

func (m *Garden) isNextPositionSameLetter(position [2]int, inc [2]int) bool {
	if position[1] == 0 && inc[1] == -1 || position[1] == len(m.OriginalGarden)-1 && inc[1] == 1 || position[0] == 0 && inc[0] == -1 || position[0] == len(m.OriginalGarden[0])-1 && inc[0] == 1 {
		// outside of the map
		return false
	}
	if m.OriginalGarden[position[1]+inc[1]][position[0]+inc[0]] == m.OriginalGarden[position[1]][position[0]] {
		return true
	}
	return false
}

func (g *Garden) GetMapIndex(p [2]int) int {
	letter := g.OriginalGarden[p[1]][p[0]]
	areas := g.Areas[letter]

	for index, area := range areas {
		found := slices.IndexFunc(area, func(point [2]int) bool {
			return point[0] == p[0] && point[1] == p[1]
		})

		if found > -1 {
			return index
		}
	}

	return -1
}

func CalculatePerimeter(points [][2]int) int {
	pointSet := make(map[[2]int]bool)

	// Store all points in a set for quick lookup
	for _, p := range points {
		pointSet[[2]int{p[0], p[1]}] = true
	}

	// Define the 4 possible directions (neighbors)
	directions := [][2]int{
		{0, 1},  // up
		{1, 0},  // right
		{0, -1}, // down
		{-1, 0}, // left
	}

	perimeter := 0

	// For each point, check how many sides are exposed
	for _, p := range points {
		for _, d := range directions {
			neighbor := [2]int{p[0] + d[0], p[1] + d[1]}
			if !pointSet[neighbor] {
				perimeter++
			}
		}
	}

	return perimeter
}

func (g *Garden) safeValue(x, y int) string {
	if x < 0 || y < 0 || y > len(g.OriginalGarden)-1 || x > len(g.OriginalGarden[0])-1 {
		return "."
	}
	return g.OriginalGarden[y][x]
}

func (g *Garden) corners(point [2]int) int {
	currentWord := g.OriginalGarden[point[1]][point[0]]
	acc := 0

	// top left
	if ((currentWord != g.safeValue(point[0], point[1]-1)) && (currentWord != g.safeValue(point[0]-1, point[1]))) || ((currentWord == g.safeValue(point[0]-1, point[1])) && (currentWord == g.safeValue(point[0], point[1]-1)) && (currentWord != g.safeValue(point[0]-1, point[1]-1))) {
		acc++
	}

	// top right
	if ((currentWord != g.safeValue(point[0], point[1]-1)) && (currentWord != g.safeValue(point[0]+1, point[1]))) || ((currentWord == g.safeValue(point[0]+1, point[1])) && (currentWord == g.safeValue(point[0], point[1]-1)) && (currentWord != g.safeValue(point[0]+1, point[1]-1))) {
		acc++
	}

	// bottom left
	if ((currentWord != g.safeValue(point[0], point[1]+1)) && (currentWord != g.safeValue(point[0]-1, point[1]))) || ((currentWord == g.safeValue(point[0]-1, point[1])) && (currentWord == g.safeValue(point[0], point[1]+1)) && (currentWord != g.safeValue(point[0]-1, point[1]+1))) {
		acc++
	}

	// bottom right
	if ((currentWord != g.safeValue(point[0], point[1]+1)) && (currentWord != g.safeValue(point[0]+1, point[1]))) || ((currentWord == g.safeValue(point[0]+1, point[1])) && (currentWord == g.safeValue(point[0], point[1]+1)) && (currentWord != g.safeValue(point[0]+1, point[1]+1))) {
		acc++
	}

	return acc
}

func (g *Garden) CountCorners(points [][2]int) int {
	corners := 0

	for _, p := range points {
		corners += g.corners(p)
	}

	return corners
}
