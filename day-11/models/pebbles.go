package Pebbles

import (
	"strconv"
)

type Pebbles struct {
	OriginalPebbles []int
	PebblesMap      map[int]int
}

func NewPebble() *Pebbles {
	return &Pebbles{OriginalPebbles: []int{}, PebblesMap: map[int]int{}}
}

func splitInTwo(value int) (int, int) {
	valueString := strconv.Itoa(value)
	length := len(valueString)
	firstHalf, _ := strconv.Atoi(valueString[:length/2])
	secondHalf, _ := strconv.Atoi(valueString[length/2:])
	return firstHalf, secondHalf
}

func (p *Pebbles) Add(pebbleNumber int) {
	p.PebblesMap[pebbleNumber]++
}

func (p *Pebbles) Blink() {
	newPebblesMap := make(map[int]int)

	for key, value := range p.PebblesMap {
		newPebblesMap[key] = value
	}

	for key, value := range p.PebblesMap {
		if value <= 0 {
			continue
		}
		if key == 0 {
			newPebblesMap[1] += value
			newPebblesMap[0] -= value
		} else if len(strconv.Itoa(key))%2 == 0 {
			value1, value2 := splitInTwo(key)
			newPebblesMap[key] -= value
			newPebblesMap[value1] += value
			newPebblesMap[value2] += value
		} else {
			newPebblesMap[key] -= value
			newPebblesMap[key*2024] += value
		}
	}
	p.PebblesMap = newPebblesMap
}
