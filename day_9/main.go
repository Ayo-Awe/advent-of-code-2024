package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var FreeSpace = Block{-1}

type Block struct {
	ID int // negative file-id indicates free space
}

func main() {
	input, err := aoc.ReadInput("input.txt")
	if err != nil {
		log.Fatal("failed to read puzzle input", err)
	}

	input = strings.TrimSuffix(input, "\n")

	fmt.Println("solution for part one: ", PartOne(input))
	fmt.Println("solution for part two: ", PartTwo(input))
}

func PartOne(input string) int {
	blocks := expandDiskMap(input)
	compactBlocks(blocks)

	// calculate checksum
	checksum := 0

	for i, block := range blocks {
		if block != FreeSpace {
			checksum += i * block.ID
		}
	}

	return checksum
}

func PartTwo(input string) int {
	blocks := expandDiskMap(input)
	compactContiguousBlocks(blocks)

	// calculate checksum
	checksum := 0

	for i, block := range blocks {
		if block != FreeSpace {
			checksum += i * block.ID
		}
	}

	return checksum
}

func expandDiskMap(diskMap string) []Block {
	var blocks []Block

	for i, c := range strings.Split(diskMap, "") {
		block := FreeSpace

		isFile := i%2 == 0
		if isFile {
			fileID := i / 2
			block = Block{fileID}
		}

		blockLength, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal("invalid character", err)
		}

		for range blockLength {
			blocks = append(blocks, block)
		}
	}

	return blocks
}

// rearranges blocks in place
func compactBlocks(blocks []Block) {
	for i, currBlock := range blocks {
		if currBlock != FreeSpace {
			continue
		}

		lastFileBlockIndex := getLastFileBlockIndex(blocks)
		if lastFileBlockIndex < 0 {
			log.Fatal("something went wrong: no file blocks found")
		}

		// all free space has been removed
		if lastFileBlockIndex < i {
			return
		}

		swap(blocks, i, lastFileBlockIndex)
	}
}

func compactContiguousBlocks(disk []Block) {
	maxFile := slices.MaxFunc(disk, func(a, b Block) int { return a.ID - b.ID })

	// loop from the max file to the first file 0
	for i := maxFile.ID; i >= 0; i-- {
		start, end := findFile(disk, i, 0)
		// some files that appear on the diskmap may not appear on disk i.e size=0
		if start == -1 || end == -1 {
			continue
		}

		fileSize := end - start + 1

		freespaceStart, _ := findFile(disk, FreeSpace.ID, fileSize)
		if freespaceStart == -1 || freespaceStart >= start {
			continue
		}

		swapN(disk, start, freespaceStart, fileSize)
	}
}

// finds a file starting from the end of the disk backwards to the beginning
func findFile(disk []Block, fileID, size int) (int, int) {
	inFile := false
	fileIndex := -1

	var i int
	for i = range disk {
		if disk[i].ID == fileID && !inFile {
			inFile = true
			fileIndex = i

		} else if disk[i].ID != fileID && inFile {
			fileSize := i - fileIndex
			if fileSize >= size {
				return fileIndex, i - 1
			}

			// reset block tracker
			inFile = false
			fileIndex = -1
		}
	}

	fileSize := i - fileIndex
	if inFile && fileSize > size {
		return fileIndex, i
	}

	return -1, -1
}

// swaps n contiguous blocks between src and dest index
func swapN(disk []Block, src, dest, n int) {
	for i := range n {
		swap(disk, src+i, dest+i)
	}
}

func getLastFileBlockIndex(blocks []Block) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] != FreeSpace {
			return i
		}
	}

	return -1
}

func swap(blocks []Block, i, j int) {
	blocks[i], blocks[j] = blocks[j], blocks[i]
}
