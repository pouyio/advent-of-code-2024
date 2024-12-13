package Block

import "fmt"

type Block struct {
	OriginalDiskMap []int
	OrderedDiskMap  []int
	UnrolledBlocks  []int
	OrderedBlocks   []int
}

func NewBlocks(blocks []int) *Block {
	return &Block{OriginalDiskMap: blocks, UnrolledBlocks: []int{}, OrderedBlocks: []int{}}
}

func (b *Block) Add(value int) {
	b.OriginalDiskMap = append(b.OriginalDiskMap, value)
	b.OrderedDiskMap = append(b.OrderedDiskMap, value)
}

func (b *Block) Loop() <-chan []int {
	c := make(chan []int)
	go func() {
		for index, block := range b.OriginalDiskMap {
			c <- []int{index, block}
		}
		close(c)
	}()
	return c
}

func (b *Block) FindLastNonEmptyIndex() int {
	for i := len(b.OrderedBlocks) - 1; i >= 0; i-- {
		if b.OrderedBlocks[i] != -1 {
			return i
		}
	}
	return -1
}

func (b *Block) FindFirstEmptyIndex1(minEmpty int) (int, int) {
	startIndex := -2
	endIndex := -1
	value := -1

	for i := len(b.OrderedBlocks) - 1; i >= 0; i-- {
		if b.OrderedBlocks[i] < 0 {
			if endIndex != -1 {
				endIndex = i
				startIndex = i
				value = b.OrderedBlocks[i]
				continue
			} else if value == b.OrderedBlocks[i] {
				startIndex = i
				continue
			} else {
				return startIndex, endIndex
			}
		}
	}

	return startIndex, endIndex
}

func (b *Block) AddIndividualBlock(value int) {
	b.UnrolledBlocks = append(b.UnrolledBlocks, value)
	b.OrderedBlocks = append(b.OrderedBlocks, value)
}

func (b *Block) CalculateChecksum() int {
	checksum := 0
	for index, block := range b.OrderedBlocks {
		if block > 0 {
			checksum += index * block
		}
	}
	return checksum
}

func (b *Block) CalculateDiskIndex(mapIndex int) int {
	index := 0
	for i := 0; i < mapIndex; i++ {
		index += b.OrderedDiskMap[i]
	}
	return index
}

func (b *Block) PrintOrderedBlocks() {
	for i := 0; i < len(b.OrderedBlocks); i++ {
		if b.OrderedBlocks[i] == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(b.OrderedBlocks[i])
		}
	}
	fmt.Println()
}
