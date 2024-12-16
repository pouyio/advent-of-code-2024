package Game

import "fmt"

const width, height = 101, 103

type State struct {
	Robots    []Robot
	cuadrants [4]int
}

type Robot struct {
	px, py, vx, vy int
}

func NewRobot(px, py, vx, vy int) *Robot {
	return &Robot{px: px, py: py, vx: vx, vy: vy}
}

func NewState() *State {
	return &State{Robots: []Robot{}, cuadrants: [4]int{0, 0, 0, 0}}
}

func (s *State) AddRobot(robot Robot) {
	s.Robots = append(s.Robots, robot)
}

func (s *State) MoveSeconds(robot *Robot, sec int) (int, int) {
	Ax := sec * robot.vx
	Ay := sec * robot.vy
	finalAx := Ax % width
	finalAy := Ay % height
	posx := robot.px + finalAx
	posy := robot.py + finalAy

	if posx < 0 {
		posx = width + posx
	} else if posx > width-1 {
		posx = posx - width
	}

	if posy < 0 {
		posy = height + posy
	} else if posy > height-1 {
		posy = posy - height
	}

	if posx < (width-1)/2 && posy < (height-1)/2 {
		s.cuadrants[0]++
	} else if posx < (width-1)/2 && posy > (height-1)/2 {
		s.cuadrants[2]++
	} else if posx > (width-1)/2 && posy < (height-1)/2 {
		s.cuadrants[1]++
	} else if posx > (width-1)/2 && posy > (height-1)/2 {
		s.cuadrants[3]++
	}

	return posx, posy
}

func (s *State) MoveRobots(sec int) map[int]map[int]bool {
	positions := map[int]map[int]bool{}
	for _, robot := range s.Robots {
		x, y := s.MoveSeconds(&robot, sec)
		if positions[y] == nil {
			positions[y] = map[int]bool{}
		}
		positions[y][x] = true
	}
	return positions
}

func (s *State) getRow(rowPositions map[int]bool) []string {
	row := []string{}
	for j := 0; j < width; j++ {
		if rowPositions[j] {
			row = append(row, "#")
		} else {
			row = append(row, ".")
		}
	}
	return row
}

func (s *State) DrawPositions(positions map[int]map[int]bool) {
	for i := 0; i < height; i++ {
		row := s.getRow(positions[i])
		fmt.Println(row)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
}

func (s *State) CouldBeTree(positions map[int]map[int]bool) bool {
	for i := 0; i < height; i++ {
		row := s.getRow(positions[i])
		for j := 0; j < len(row); j++ {
			if j < width-8 && j > 8 {
				// find a straight line and check visually
				if row[j] == row[j-1] && row[j] == row[j-2] && row[j] == row[j-3] && row[j] == row[j-4] && row[j] == row[j-5] && row[j] == row[j-6] && row[j] == row[j-7] && row[j] == row[j-8] && row[j] == "#" {
					return true
				}
			}
		}
	}
	return false
}

func (s *State) SafeFactor() int {
	acc := 1
	for _, cuadrant := range s.cuadrants {
		acc *= cuadrant
	}
	return acc
}
