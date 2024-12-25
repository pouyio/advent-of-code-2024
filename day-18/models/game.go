package Game

import (
	"container/heap"
)

const SIZE = 70

type Node struct {
	cost      int
	x, y      int
	direction int
}
type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func ShortestPath(corrupted map[[2]int]bool, start, end [2]int) int {
	rows := SIZE + 1
	cols := SIZE + 1
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{cost: 0, x: start[0], y: start[1], direction: 0})
	visited := make(map[[3]int]bool)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)

		if curr.x == end[0] && curr.y == end[1] {
			return curr.cost
		}

		key := [3]int{curr.x, curr.y, curr.direction}
		if visited[key] {
			continue
		}
		visited[key] = true

		for i, d := range directions {
			nx, ny := curr.x+d[0], curr.y+d[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !corrupted[[2]int{nx, ny}] {
				newCost := curr.cost + 1
				heap.Push(pq, Node{
					cost:      newCost,
					x:         nx,
					y:         ny,
					direction: i,
				})
			}
		}
	}

	return -1
}

type State struct {
	Corrupted map[[2]int]bool
}

func NewState() *State {
	return &State{Corrupted: map[[2]int]bool{}}
}

func (s *State) AddCorruptedCell(x, y int) {
	s.Corrupted[[2]int{x, y}] = true
}
