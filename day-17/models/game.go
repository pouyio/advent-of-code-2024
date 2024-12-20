package Game

type State struct {
	originalA, originalB, originalC int
	A, B, C                         int
	OriginalProgram                 []int
	Out                             []int
	Pointer                         int
	Opcodes                         map[int]func(int)
}

func NewState(a, b, c int, program []int) *State {
	computer := &State{originalA: a, A: a, originalB: b, B: b, originalC: c, C: c, OriginalProgram: program, Out: []int{}, Pointer: 0}

	computer.Opcodes = map[int]func(operand int){
		0: func(operand int) {
			result := computer.A / (1 << computer.GetOperandCombo(operand))
			computer.A = result
			computer.Pointer += 2
		},
		1: func(operand int) {
			computer.B ^= operand
			computer.Pointer += 2
		},
		2: func(operand int) {
			result := computer.GetOperandCombo(operand) % 8
			computer.B = result
			computer.Pointer += 2
		},
		3: func(operand int) {
			if computer.A == 0 {
				computer.Pointer += 2
			} else {
				computer.Pointer = operand
			}
		},
		4: func(operand int) {
			computer.B ^= computer.C
			computer.Pointer += 2
		},
		5: func(operand int) {
			result := computer.GetOperandCombo(operand) % 8
			result = result & 7
			computer.Out = append(computer.Out, result)
			computer.Pointer += 2
		},
		6: func(operand int) {
			result := computer.A / (1 << computer.GetOperandCombo(operand))
			computer.B = int(result)
			computer.Pointer += 2
		},
		7: func(operand int) {
			result := computer.A / (1 << computer.GetOperandCombo(operand))
			computer.C = (result)
			computer.Pointer += 2
		},
	}

	return computer
}

func (s *State) GetOperandCombo(operand int) int {
	if operand <= 3 {
		return operand
	}

	if operand == 4 {
		return s.A
	} else if operand == 5 {
		return s.B
	} else if operand == 6 {
		return s.C
	} else {
		panic("Imposibru")
	}
}
