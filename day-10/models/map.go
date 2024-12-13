package Map

import (
	"fmt"
	"os"
)

type Map struct {
	OriginalMap            [][]int
	Trailheads             [][2]int
	TrailHeadsScores       map[[2]int]bool
	TrailHeadsScoresRating int
}

func NewMap() *Map {
	return &Map{OriginalMap: [][]int{}, Trailheads: [][2]int{}, TrailHeadsScores: map[[2]int]bool{}, TrailHeadsScoresRating: 0}
}

type increment struct {
	x, y int
}

func (m *Map) CalculateTrailheadScore(trailhead [2]int) bool {
	if m.OriginalMap[trailhead[1]][trailhead[0]] == 9 {
		_, ok := m.TrailHeadsScores[trailhead]
		if !ok {
			m.TrailHeadsScores[trailhead] = true
		}
		return true
	}

	if m.isNextPositionValid(trailhead, increment{0, -1}) {
		m.CalculateTrailheadScore([2]int{trailhead[0], trailhead[1] - 1})
	}
	if m.isNextPositionValid(trailhead, increment{0, 1}) {
		m.CalculateTrailheadScore([2]int{trailhead[0], trailhead[1] + 1})
	}
	if m.isNextPositionValid(trailhead, increment{-1, 0}) {
		m.CalculateTrailheadScore([2]int{trailhead[0] - 1, trailhead[1]})
	}
	if m.isNextPositionValid(trailhead, increment{1, 0}) {
		m.CalculateTrailheadScore([2]int{trailhead[0] + 1, trailhead[1]})
	}
	return false
}

func (m *Map) CalculateTrailheadScore1(trailhead [2]int) bool {
	if m.OriginalMap[trailhead[1]][trailhead[0]] == 9 {
		m.TrailHeadsScoresRating++
		return true
	}

	if m.isNextPositionValid(trailhead, increment{0, -1}) {
		m.CalculateTrailheadScore1([2]int{trailhead[0], trailhead[1] - 1})
	}
	if m.isNextPositionValid(trailhead, increment{0, 1}) {
		m.CalculateTrailheadScore1([2]int{trailhead[0], trailhead[1] + 1})
	}
	if m.isNextPositionValid(trailhead, increment{-1, 0}) {
		m.CalculateTrailheadScore1([2]int{trailhead[0] - 1, trailhead[1]})
	}
	if m.isNextPositionValid(trailhead, increment{1, 0}) {
		m.CalculateTrailheadScore1([2]int{trailhead[0] + 1, trailhead[1]})
	}
	return false
}

func (m *Map) drawMapWithCursor(position [2]int) {
	for y, row := range m.OriginalMap {
		for x, value := range row {
			if x == position[0] && y == position[1] {
				fmt.Fprintf(os.Stdout, "\033[0;31m%d\033[0m", value)
			} else {
				fmt.Print(value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *Map) isNextPositionValid(position [2]int, inc increment) bool {
	if position[1] == 0 && inc.y == -1 || position[1] == len(m.OriginalMap)-1 && inc.y == 1 || position[0] == 0 && inc.x == -1 || position[0] == len(m.OriginalMap[0])-1 && inc.x == 1 {
		// outside of the map
		return false
	}

	if m.OriginalMap[position[1]+inc.y][position[0]+inc.x] == m.OriginalMap[position[1]][position[0]]+1 {
		return true
	}

	return false

}
