package Game

import (
	"fmt"
	"strings"
)

func (s *State) findShortcuts(currentPosition [2]int) [][2]int {
	shorctus := [][2]int{}
	currentTime := s.PathTimeMap[[2]int{currentPosition[0], currentPosition[1]}]
	if s.PathTimeMap[[2]int{currentPosition[0], currentPosition[1] - 2}] > currentTime && strings.Contains("#E", string(s.Map[currentPosition[1]-1][currentPosition[0]])) { // up
		shorctus = append(shorctus, [2]int{currentPosition[0], currentPosition[1] - 2})
	}
	if s.PathTimeMap[[2]int{currentPosition[0], currentPosition[1] + 2}] > currentTime && strings.Contains("#E", string(s.Map[currentPosition[1]+1][currentPosition[0]])) { // down
		shorctus = append(shorctus, [2]int{currentPosition[0], currentPosition[1] + 2})
	}
	if s.PathTimeMap[[2]int{currentPosition[0] + 2, currentPosition[1]}] > currentTime && strings.Contains("#E", string(s.Map[currentPosition[1]][currentPosition[0]+1])) { // right
		shorctus = append(shorctus, [2]int{currentPosition[0] + 2, currentPosition[1]})
	}
	if s.PathTimeMap[[2]int{currentPosition[0] - 2, currentPosition[1]}] > currentTime && strings.Contains("#E", string(s.Map[currentPosition[1]][currentPosition[0]-1])) { // left
		shorctus = append(shorctus, [2]int{currentPosition[0] - 2, currentPosition[1]})
	}

	return shorctus
}

func (s *State) isValidPosition(next, current [2]int) bool {
	return s.PathTimeMap[next] > s.PathTimeMap[current]
}

func (s *State) findShortcuts1(current [2]int) map[[2]int]int {
	time := 20
	shortcuts := map[[2]int]int{}

	for cheatTime := 2; cheatTime <= time; cheatTime++ {
		for deltaI := 0; deltaI <= cheatTime; deltaI++ {
			deltaJ := cheatTime - deltaI
			if s.isValidPosition([2]int{current[0] + deltaI, current[1] + deltaJ}, current) {
				shortcuts[[2]int{current[0] + deltaI, current[1] + deltaJ}] = cheatTime
				// s.drawMap(current, [2]int{current[0] + deltaI, current[1] + deltaJ})
			}
			if s.isValidPosition([2]int{current[0] - deltaI, current[1] + deltaJ}, current) {
				shortcuts[[2]int{current[0] - deltaI, current[1] + deltaJ}] = cheatTime
				// s.drawMap(current, [2]int{current[0] - deltaI, current[1] + deltaJ})
			}
			if s.isValidPosition([2]int{current[0] + deltaI, current[1] - deltaJ}, current) {
				shortcuts[[2]int{current[0] + deltaI, current[1] - deltaJ}] = cheatTime
				// s.drawMap(current, [2]int{current[0] + deltaI, current[1] - deltaJ})
			}
			if s.isValidPosition([2]int{current[0] - deltaI, current[1] - deltaJ}, current) {
				shortcuts[[2]int{current[0] - deltaI, current[1] - deltaJ}] = cheatTime
				// s.drawMap(current, [2]int{current[0] - deltaI, current[1] - deltaJ})
			}
		}
	}

	return shortcuts
}

func (s *State) getNextPosition(current [2]int) [2]int {
	if current[1]-1 >= 0 && s.PathTimeMap[[2]int{current[0], current[1] - 1}] == 0 && strings.Contains(".E", string(s.Map[current[1]-1][current[0]])) { // up
		return [2]int{current[0], current[1] - 1}
	} else if current[1]+1 < len(s.Map) && s.PathTimeMap[[2]int{current[0], current[1] + 1}] == 0 && strings.Contains(".E", string(s.Map[current[1]+1][current[0]])) { // down
		return [2]int{current[0], current[1] + 1}
	} else if current[0]+1 < len(s.Map[0]) && s.PathTimeMap[[2]int{current[0] + 1, current[1]}] == 0 && strings.Contains(".E", string(s.Map[current[1]][current[0]+1])) { // right
		return [2]int{current[0] + 1, current[1]}
	} else if current[0]-1 >= 0 && s.PathTimeMap[[2]int{current[0] - 1, current[1]}] == 0 && strings.Contains(".E", string(s.Map[current[1]][current[0]-1])) { // left
		return [2]int{current[0] - 1, current[1]}
	}
	panic("imposibru")
}

func (s *State) GetStartPosition() [2]int {
	for i, row := range s.Map {
		for j, cell := range row {
			if string(cell) == "S" {
				return [2]int{j, i}
			}
		}
	}
	return [2]int{-1, -1}
}

func (s *State) GetEndPosition() [2]int {
	for i, row := range s.Map {
		for j, cell := range row {
			if string(cell) == "E" {
				return [2]int{j, i}
			}
		}
	}
	return [2]int{-1, -1}
}

type State struct {
	Map         []string
	PathTimeMap map[[2]int]int
	paths       [][2]int
	cheats      map[[2]int]bool
}

func NewState() *State {
	return &State{Map: []string{}, PathTimeMap: map[[2]int]int{}, paths: [][2]int{}, cheats: map[[2]int]bool{}}
}

func (s *State) AddMapRow(row string) {
	s.Map = append(s.Map, row)
}

func (s *State) FillPathTime() {
	start := s.GetStartPosition()
	end := s.GetEndPosition()
	s.PathTimeMap[start] = 0
	s.paths = append(s.paths, start)

	currentPosition := start
	nextPosition := [2]int{-1, -1}

	counter := 1
	for !(nextPosition[0] == end[0] && nextPosition[1] == end[1]) {
		nextPosition = s.getNextPosition(currentPosition)
		s.PathTimeMap[nextPosition] = counter
		s.paths = append(s.paths, nextPosition)
		counter++
		currentPosition = nextPosition
	}
}

func (s *State) CalculateTotalTime() int {
	acc := 0
	for _, value := range s.PathTimeMap {
		if value > acc {
			acc = value
		}
	}
	return acc
}

func (s *State) drawMap(position, cheat [2]int) {
	for i := range s.Map {
		fmt.Println()
		for j := range s.Map[i] {
			if i == cheat[1] && j == cheat[0] {
				fmt.Print("C")
			} else if i == position[1] && j == position[0] {
				fmt.Print("X")
			} else {
				fmt.Print(string(s.Map[i][j]))
			}
		}
	}
	fmt.Println()
}

func (s *State) FindCheats() [][2]int {
	cheats := [][2]int{}
	for _, path := range s.paths {
		shortcuts := s.findShortcuts(path)
		if len(shortcuts) != 0 {
			for _, shortcut := range shortcuts {
				if s.PathTimeMap[shortcut]-s.PathTimeMap[path]-2 >= 100 {
					cheats = append(cheats, shortcut)
				}
			}
		} else {
		}
	}
	return cheats
}

func (s *State) FindCheats1() map[[2]int]bool {
	cheats := map[[2]int]bool{}
	for idx, path := range s.paths {
		shortcuts := s.findShortcuts1(path)
		if len(shortcuts) != 0 {
			for shortcut, cheatTime := range shortcuts {
				if s.PathTimeMap[shortcut]-idx-cheatTime == 72 {
					cheats[shortcut] = true
				}
			}
		}
	}
	return cheats
}
