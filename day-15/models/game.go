package Game

import (
	"fmt"
	"os"
)

var incrementsMap = map[string][2]int{
	"<": {-1, 0},
	">": {1, 0},
	"v": {0, 1},
	"^": {0, -1},
}

type State struct {
	Map       []string
	Movements string
}

func (s *State) GetRobotPosition() [2]int {
	for i, row := range s.Map {
		for j, cell := range row {
			if string(cell) == "@" {
				return [2]int{j, i}
			}
		}
	}
	panic("Imposibru")
}

func NewState() *State {
	return &State{Map: []string{}, Movements: ""}
}

func (s *State) AddMapRow(row string) {
	s.Map = append(s.Map, row)
}

func (s *State) AddMapRow1(row string) {
	newRow := ""
	for _, cell := range row {
		if string(cell) == "#" {
			newRow = newRow + "##"
		} else if string(cell) == "O" {
			newRow = newRow + "[]"
		} else if string(cell) == "." {
			newRow = newRow + ".."
		} else if string(cell) == "@" {
			newRow = newRow + "@."
		}
	}
	s.Map = append(s.Map, newRow)
}

func (s *State) DrawMap(boxMap []string) {
	for index, row := range boxMap {
		position := s.GetRobotPosition()
		if index == position[1] {
			fmt.Print(row[:position[0]])
			fmt.Fprintf(os.Stdout, "\033[0;31m%s\033[0m", "▮")
			fmt.Println(row[position[0]+1:])
		} else {
			fmt.Println(row)
		}
	}
	fmt.Println("-------------")
}

func (s *State) replaceAt(y int, index int, replacement string) string {
	str := s.Map[y]
	if index < 0 || index >= len(str) {
		return str
	}
	return str[:index] + replacement + str[index+1:]
}

func oppositeBoxSide(side string) string {
	if side == "[" {
		return "]"
	} else if side == "]" {
		return "["
	} else if side == "@" {
		return "."
	}
	return side
}

func (s *State) moveBox(position [2]int, increment string) {
	nextPosition := [2]int{position[0] + incrementsMap[increment][0], position[1] + incrementsMap[increment][1]}
	nextElement := string(s.Map[nextPosition[1]][nextPosition[0]])

	if nextElement == "O" {
		s.moveBox(nextPosition, increment)
	} else if nextElement == "." {
		s.Map[nextPosition[1]] = s.replaceAt(nextPosition[1], nextPosition[0], "O")
	} else if nextElement == "#" {
		return
	} else {
		panic("Imposibruu")
	}
}

func (s *State) moveBoxHorizontal(side string, nextPosition [2]int, increment string) {
	nextElement := string(s.Map[nextPosition[1]][nextPosition[0]])

	s.Map[nextPosition[1]] = s.replaceAt(nextPosition[1], nextPosition[0], side)

	if increment == "<" || increment == ">" {
		if nextElement == "[" || nextElement == "]" {
			s.moveBoxHorizontal(oppositeBoxSide(side), [2]int{nextPosition[0] + incrementsMap[increment][0], nextPosition[1] + incrementsMap[increment][1]}, increment)
			return
		} else if nextElement == "." || nextElement == "#" {
			return
		}
	}
}

func (s *State) moveVertical(positions map[[2]int]string, increment int) {
	newPositions := map[[2]int]string{}

	for position, element := range positions {
		replacedValue := string(s.Map[position[1]+increment][position[0]])
		s.Map[position[1]+increment] = s.replaceAt(position[1]+increment, position[0], element)

		if replacedValue != "." {
			newPositions[[2]int{position[0], position[1] + increment}] = replacedValue
		}

		if replacedValue == "[" {
			replacedValue2 := string(s.Map[position[1]+increment][position[0]+1])
			s.Map[position[1]+increment] = s.replaceAt(position[1]+increment, position[0]+1, ".")
			if replacedValue2 != "." {
				newPositions[[2]int{position[0] + 1, position[1] + increment}] = replacedValue2
			}
		} else if replacedValue == "]" {
			replacedValue2 := string(s.Map[position[1]+increment][position[0]-1])
			s.Map[position[1]+increment] = s.replaceAt(position[1]+increment, position[0]-1, ".")
			if replacedValue2 != "." {
				newPositions[[2]int{position[0] - 1, position[1] + increment}] = replacedValue2
			}
		}
	}

	if len(newPositions) != 0 {
		s.moveVertical(newPositions, increment)
	}
}

func (s *State) ConsumeMovement(increment string) {
	currentPosition := s.GetRobotPosition()
	if s.canMove(currentPosition, increment) {
		nextPosition := [2]int{currentPosition[0] + incrementsMap[increment][0], currentPosition[1] + incrementsMap[increment][1]}
		nextElement := string(s.Map[nextPosition[1]][nextPosition[0]])

		s.Map[currentPosition[1]] = s.replaceAt(currentPosition[1], currentPosition[0], ".")
		s.Map[nextPosition[1]] = s.replaceAt(nextPosition[1], nextPosition[0], "@")

		if nextElement == "O" {
			s.moveBox(nextPosition, increment)
		}
	}

	if len(s.Movements) > 0 {
		s.Movements = s.Movements[1 : len(s.Movements)-1]
	}
}

func (s *State) ConsumeMovement1(increment string) {
	currentPosition := s.GetRobotPosition()
	if s.canMove1(currentPosition, increment) {
		nextPosition := [2]int{currentPosition[0] + incrementsMap[increment][0], currentPosition[1] + incrementsMap[increment][1]}
		nextElement := string(s.Map[nextPosition[1]][nextPosition[0]])

		if increment == "<" || increment == ">" {
			s.Map[currentPosition[1]] = s.replaceAt(currentPosition[1], currentPosition[0], ".")
			s.Map[nextPosition[1]] = s.replaceAt(nextPosition[1], nextPosition[0], "@")

			if nextElement == "[" || nextElement == "]" {
				nextBoxPosition := [2]int{nextPosition[0] + incrementsMap[increment][0], nextPosition[1] + incrementsMap[increment][1]}
				if increment == "<" || increment == ">" {
					s.moveBoxHorizontal(nextElement, nextBoxPosition, increment)
				}
			}
		} else {
			if nextElement == "[" || nextElement == "]" {
				s.Map[currentPosition[1]] = s.replaceAt(currentPosition[1], currentPosition[0], ".")
				positionMap := map[[2]int]string{currentPosition: "@"}
				if increment == "v" {
					if nextElement == "[" {
						positionMap[[2]int{currentPosition[0] + 1, currentPosition[1]}] = "."
						s.moveVertical(positionMap, 1)
					} else {
						positionMap[[2]int{currentPosition[0] - 1, currentPosition[1]}] = "."
						s.moveVertical(positionMap, 1)
					}
				} else {
					if nextElement == "[" {
						positionMap[[2]int{currentPosition[0] + 1, currentPosition[1]}] = "."
						s.moveVertical(positionMap, -1)
					} else if nextElement == "]" {
						positionMap[[2]int{currentPosition[0] - 1, currentPosition[1]}] = "."
						s.moveVertical(positionMap, -1)
					}
				}
			} else {
				s.Map[currentPosition[1]] = s.replaceAt(currentPosition[1], currentPosition[0], ".")
				s.Map[nextPosition[1]] = s.replaceAt(nextPosition[1], nextPosition[0], "@")
			}
		}
	}

	if len(s.Movements) > 0 {
		s.Movements = s.Movements[1 : len(s.Movements)-1]
	}
}

func (s *State) canMove(currentPosition [2]int, increment string) bool {
	nextPosition := [2]int{currentPosition[0] + incrementsMap[increment][0], currentPosition[1] + incrementsMap[increment][1]}

	if nextPosition[0] < 1 || nextPosition[0] > len(s.Map[0])-2 {
		return false
	}
	if nextPosition[1] < 1 || nextPosition[1] > len(s.Map)-2 {
		return false
	}

	elementNextPosition := string(s.Map[nextPosition[1]][nextPosition[0]])

	if elementNextPosition == "#" {
		return false
	}

	if elementNextPosition == "." {
		return true
	}

	if elementNextPosition == "O" {
		return s.canMove(nextPosition, increment)
	}

	return false
}
func (s *State) canMove1(currentPosition [2]int, increment string) bool {
	nextPosition := [2]int{currentPosition[0] + incrementsMap[increment][0], currentPosition[1] + incrementsMap[increment][1]}

	if nextPosition[0] < 1 || nextPosition[0] > len(s.Map[0])-2 {
		return false
	}
	if nextPosition[1] < 1 || nextPosition[1] > len(s.Map)-2 {
		return false
	}

	elementNextPosition := string(s.Map[nextPosition[1]][nextPosition[0]])

	if elementNextPosition == "#" {
		return false
	}

	if elementNextPosition == "." {
		return true
	}

	if elementNextPosition == "[" || elementNextPosition == "]" {
		if increment == "<" || increment == ">" {
			return s.canMove1(nextPosition, increment)
		}

		canMove1 := s.canMove1(nextPosition, increment)
		canMove2 := false
		if elementNextPosition == "[" {
			canMove2 = s.canMove1([2]int{nextPosition[0] + 1, nextPosition[1]}, increment)
		} else {
			canMove2 = s.canMove1([2]int{nextPosition[0] - 1, nextPosition[1]}, increment)
		}
		return canMove1 && canMove2
	}

	return false
}

func (s *State) CalculateBoxesGPS() int {
	acc := 0
	for i, row := range s.Map {
		for j, cell := range row {
			if string(cell) == "O" {
				acc += 100*i + j
			}
		}
	}
	return acc
}

func (s *State) CalculateBoxesGPS1() int {
	acc := 0
	for i, row := range s.Map {
		for j, cell := range row {
			if string(cell) == "[" {
				acc += 100*i + j
			}
		}
	}
	return acc
}
