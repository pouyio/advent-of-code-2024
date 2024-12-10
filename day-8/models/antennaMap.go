package AntennaMap

type AntennaMap struct {
	Map          [][]string
	AntinodesMap [][]string
	AntennaPairs map[string][][2]int
	Antinodes    int
}

func NewAntennaMap(originalMap [][]string) *AntennaMap {
	antinodesMap := make([][]string, len(originalMap))
	for i := range originalMap {
		antinodesMap[i] = make([]string, len(originalMap[i]))
		for j := range originalMap[i] {
			antinodesMap[i][j] = " "
		}
	}
	return &AntennaMap{Map: originalMap, AntennaPairs: map[string][][2]int{}, Antinodes: 0, AntinodesMap: antinodesMap}
}

func (a *AntennaMap) AddAntennaPair(character string, i, j int) {
	a.AntennaPairs[character] = append(a.AntennaPairs[character], [2]int{j, i})
}

func (a *AntennaMap) CalculateAntinodes(character string, position1 [2]int, position2 [2]int) {
	antinode1 := [2]int{0, 0}
	antinode2 := [2]int{0, 0}
	distanceX := position2[0] - position1[0]
	distanceY := position2[1] - position1[1]

	antinode1[0] = position1[0] - distanceX
	antinode2[0] = position2[0] + distanceX

	antinode1[1] = position1[1] - distanceY
	antinode2[1] = position2[1] + distanceY

	if antinode1[0] >= 0 && antinode1[0] <= len(a.Map[0])-1 && antinode1[1] >= 0 && antinode1[1] <= len(a.Map)-1 {
		if a.AntinodesMap[antinode1[1]][antinode1[0]] != "#" {
			a.AntinodesMap[antinode1[1]][antinode1[0]] = "#"
			a.Antinodes++
		}
	}

	if antinode2[0] >= 0 && antinode2[0] <= len(a.Map[0])-1 && antinode2[1] >= 0 && antinode2[1] <= len(a.Map)-1 {
		if a.AntinodesMap[antinode2[1]][antinode2[0]] != "#" {
			a.AntinodesMap[antinode2[1]][antinode2[0]] = "#"
			a.Antinodes++
		}
	}
}

func (a *AntennaMap) isInsideMap(x, y int) bool {
	return x >= 0 && x <= len(a.Map[0])-1 && y >= 0 && y <= len(a.Map)-1
}

func (a *AntennaMap) CalculateAntinodes1(character string, position1 [2]int, position2 [2]int) {
	antinode1 := [2]int{0, 0}
	antinode2 := [2]int{0, 0}
	distanceX := position2[0] - position1[0]
	distanceY := position2[1] - position1[1]

	antinode1[0] = position1[0]
	antinode2[0] = position2[0]

	antinode1[1] = position1[1]
	antinode2[1] = position2[1]

	for a.isInsideMap(antinode1[0], antinode1[1]) {
		if a.AntinodesMap[antinode1[1]][antinode1[0]] != "#" {
			a.AntinodesMap[antinode1[1]][antinode1[0]] = "#"
			a.Antinodes++
		}

		antinode1[0] -= distanceX
		antinode1[1] -= distanceY

	}

	for a.isInsideMap(antinode2[0], antinode2[1]) {
		if a.AntinodesMap[antinode2[1]][antinode2[0]] != "#" {
			a.AntinodesMap[antinode2[1]][antinode2[0]] = "#"
			a.Antinodes++
		}

		antinode2[0] += distanceX
		antinode2[1] += distanceY

	}

	// while antinode1 is within bounds, add it to the antinode map
	// if a.isInsideMap(antinode1[0], antinode1[1]) {
	// 	if a.AntinodesMap[antinode1[1]][antinode1[0]] != "#" {
	// 		a.AntinodesMap[antinode1[1]][antinode1[0]] = "#"
	// 		a.Antinodes++
	// 	}
	// }
	//
	// if a.isInsideMap(antinode2[0], antinode2[1]) {
	// 	if a.AntinodesMap[antinode2[1]][antinode2[0]] != "#" {
	// 		a.AntinodesMap[antinode2[1]][antinode2[0]] = "#"
	// 		a.Antinodes++
	// 	}
	// }
	// while antinode2 is within bounds, add it to the antinode map
}
