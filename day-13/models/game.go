package Game

const aTokens, bTokens int = 3, 1

type ClawMachine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

type GameSolution struct {
	A int
	B int
}

type State struct {
	Games     []ClawMachine
	Solutions []GameSolution
}

func NewState() *State {
	return &State{Games: []ClawMachine{}}
}

func (s *State) AddGame(g ClawMachine) {
	s.Games = append(s.Games, g)
}

// Find the minimum cost to align the claw to the prize
func (s *State) MinimizeCost(cm ClawMachine) (int, bool) {
	minCost := aTokens*100 + bTokens*100
	found := false
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			if cm.Ax*x+cm.Bx*y == cm.Px && cm.Ay*x+cm.By*y == cm.Py {
				cost := 3*x + y
				if cost < minCost {
					minCost = cost
					found = true
				}
			}
		}
	}
	return minCost, found
}

func (s *State) SolveMachine(cm ClawMachine) (int, bool) {
	// Cramer rule
	A := (cm.Px*cm.By - cm.Py*cm.Bx) / (cm.Ax*cm.By - cm.Ay*cm.Bx)
	B := (cm.Ax*cm.Py - cm.Ay*cm.Px) / (cm.Ax*cm.By - cm.Ay*cm.Bx)
	return A*3 + B, A <= 10000000000000 && B <= 10000000000000 && A*cm.Ax+B*cm.Bx == cm.Px && A*cm.Ay+B*cm.By == cm.Py
}
