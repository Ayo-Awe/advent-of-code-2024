package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var (
	directions = [4][2]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}
	X     = 0
	Y     = 1
	Steps = 2
)

func parseInput(input []string) ([][2]int, error) {
	var bytePositions [][2]int
	for _, line := range input {
		strNums := strings.Split(line, ",")

		var pos [2]int
		for i, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, err
			}
			pos[i] = num
		}
		bytePositions = append(bytePositions, pos)
	}

	return bytePositions, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	bytePositions, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(bytePositions))
	fmt.Println("solution to part two: ", PartTwo(bytePositions))
}

func PartOne(bytePositions [][2]int) int {
	corrupted := make(map[[2]int]struct{})
	n := 1 << 10 // 1024
	gridSize := 71

	for i := range n {
		bytePos := bytePositions[i]
		corrupted[bytePos] = struct{}{}
	}

	return solve(gridSize, corrupted)
}

func PartTwo(bytePositions [][2]int) string {
	corrupted := make(map[[2]int]struct{})
	n := 1 << 10 // 1024
	gridSize := 71

	for i := range n {
		bytePos := bytePositions[i]
		corrupted[bytePos] = struct{}{}
	}

	// incrementally add the remaining bytes until we find the
	// first run without a solution
	var blockingByte [2]int
	for ; n < len(bytePositions); n++ {
		bytePos := bytePositions[n]
		corrupted[bytePos] = struct{}{}

		minSteps := solve(gridSize, corrupted)

		// the first square to block the path occurs on the first run where there's
		// no solution i.e minSteps == 0
		if minSteps == 0 {
			blockingByte = bytePos
			break
		}
	}

	return fmt.Sprintf("%d,%d", blockingByte[X], blockingByte[Y])
}

func solve(gridSize int, corrupted map[[2]int]struct{}) int {
	seen := make(map[[2]int]struct{})
	endPos := [2]int{gridSize - 1, gridSize - 1}

	startNode := [3]int{}
	queue := [][3]int{startNode}

	var minSteps int
	for len(queue) > 0 {
		currNode := queue[0]
		curPos := [2]int{currNode[X], currNode[Y]}
		queue = queue[1:]

		if _, ok := seen[curPos]; ok {
			continue
		}

		if _, ok := corrupted[curPos]; ok {
			continue
		}

		if curPos[X] < 0 || curPos[X] >= gridSize || curPos[Y] < 0 || curPos[Y] >= gridSize {
			continue
		}

		if curPos == endPos {
			minSteps = currNode[Steps]
			break
		}

		for _, dir := range directions {
			queue = append(queue, [3]int{currNode[X] + dir[X], currNode[Y] + dir[Y], currNode[Steps] + 1})
		}

		seen[curPos] = struct{}{}
	}

	return minSteps
}
