package Game1

import (
	"container/heap"
	"fmt"
	"strings"
)

type Position struct {
	x, y      int
	direction int // 0=east, 1=south, 2=west, 3=north
}

type Node struct {
	pos    Position
	fScore int
	gScore int
	path   []Position
	index  int // for heap implementation
}

// PriorityQueue implementation
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].fScore < pq[j].fScore }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type Maze struct {
	grid     [][]rune
	start    Position
	end      Position
	bestCost int
}

func ParseMaze(input string) Maze {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var start, end Position

	for y, line := range lines {
		grid[y] = []rune(line)
		for x, ch := range line {
			if ch == 'S' {
				start = Position{x, y, 0} // facing east
			} else if ch == 'E' {
				end = Position{x, y, 0}
			}
		}
	}

	return Maze{
		grid:  grid,
		start: start,
		end:   end,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(p1, p2 Position) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y)
}

func getNeighbors(pos Position, maze Maze) []struct {
	pos  Position
	cost int
} {
	neighbors := []struct {
		pos  Position
		cost int
	}{}

	// Rotations
	neighbors = append(neighbors,
		struct {
			pos  Position
			cost int
		}{Position{pos.x, pos.y, (pos.direction + 1) % 4}, 1000}, // clockwise
		struct {
			pos  Position
			cost int
		}{Position{pos.x, pos.y, (pos.direction + 3) % 4}, 1000}) // counter-clockwise

	// Forward movement
	dx := []int{1, 0, -1, 0}[pos.direction]
	dy := []int{0, 1, 0, -1}[pos.direction]
	newX, newY := pos.x+dx, pos.y+dy

	if newY >= 0 && newY < len(maze.grid) &&
		newX >= 0 && newX < len(maze.grid[0]) &&
		maze.grid[newY][newX] != '#' {
		neighbors = append(neighbors, struct {
			pos  Position
			cost int
		}{Position{newX, newY, pos.direction}, 1})
	}

	return neighbors
}

func FindOptimalPaths(maze Maze) map[Position]bool {
	// Initialize priority queue for A*
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Track visited states and optimal paths
	visited := make(map[Position]int)
	optimalTiles := make(map[Position]bool)

	// Push start state
	startNode := &Node{
		pos:    maze.start,
		gScore: 0,
		fScore: manhattanDistance(maze.start, maze.end),
		path:   []Position{maze.start},
	}
	heap.Push(&pq, startNode)

	maze.bestCost = -1 // Initialize with invalid cost

	for pq.Len() > 0 {
		fmt.Println(pq.Len())
		current := heap.Pop(&pq).(*Node)

		// If we've found the end
		if current.pos.x == maze.end.x && current.pos.y == maze.end.y {
			if maze.bestCost == -1 || current.gScore == maze.bestCost {
				maze.bestCost = current.gScore
				// Mark all tiles in this optimal path
				for _, pos := range current.path {
					optimalTiles[Position{pos.x, pos.y, 0}] = true
				}
			}
			continue
		}

		// Skip if we've found a better path to this state
		if prevCost, exists := visited[current.pos]; exists && prevCost < current.gScore {
			continue
		}

		// Skip if we've found the best cost and this path is already worse
		if maze.bestCost != -1 && current.gScore > maze.bestCost {
			continue
		}

		visited[current.pos] = current.gScore

		// Process neighbors
		for _, neighbor := range getNeighbors(current.pos, maze) {
			newGScore := current.gScore + neighbor.cost

			// Skip if we've found a better path to this state
			if prevCost, exists := visited[neighbor.pos]; exists && prevCost < newGScore {
				continue
			}

			// Skip if we've found the best cost and this path would be worse
			if maze.bestCost != -1 && newGScore > maze.bestCost {
				continue
			}

			newNode := &Node{
				pos:    neighbor.pos,
				gScore: newGScore,
				fScore: newGScore + manhattanDistance(neighbor.pos, maze.end),
				path:   append(append([]Position{}, current.path...), neighbor.pos),
			}
			heap.Push(&pq, newNode)
		}
	}

	return optimalTiles
}

func VisualizeOptimalPaths(maze Maze, optimalTiles map[Position]bool) string {
	result := make([][]rune, len(maze.grid))
	for y, row := range maze.grid {
		result[y] = make([]rune, len(row))
		copy(result[y], row)
		for x, ch := range row {
			if ch != '#' && optimalTiles[Position{x, y, 0}] {
				result[y][x] = 'O'
			}
		}
	}

	var output strings.Builder
	for _, row := range result {
		output.WriteString(string(row) + "\n")
	}
	return output.String()
}
