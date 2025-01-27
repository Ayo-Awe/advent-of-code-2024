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
	end := [2]int{gridSize - 1, gridSize - 1}

	for i := range n {
		bytePos := bytePositions[i]
		corrupted[bytePos] = struct{}{}
	}

	seen := make(map[[2]int]struct{})
	queue := [][3]int{{}}

	var minSteps int
	for len(queue) > 0 {
		curr := queue[0]
		pos := [2]int{curr[X], curr[Y]}
		queue = queue[1:]

		if _, ok := seen[pos]; ok {
			continue
		}

		if _, ok := corrupted[pos]; ok {
			continue
		}

		if pos[X] < 0 || pos[X] >= gridSize || pos[Y] < 0 || pos[Y] >= gridSize {
			continue
		}

		if pos == end {
			minSteps = curr[Steps]
			break
		}

		for _, dir := range directions {
			queue = append(queue, [3]int{curr[X] + dir[X], curr[Y] + dir[Y], curr[Steps] + 1})
		}

		seen[pos] = struct{}{}
	}

	return minSteps
}

func PartTwo(bytePositions [][2]int) string {
	corrupted := make(map[[2]int]struct{})
	n := 1 << 10 // 1024
	gridSize := 71
	end := [2]int{gridSize - 1, gridSize - 1}

	for i := range n {
		bytePos := bytePositions[i]
		corrupted[bytePos] = struct{}{}
	}

	var blockingByte [2]int
	for ; n < len(bytePositions); n++ {
		seen := make(map[[2]int]struct{})
		queue := [][3]int{{}}
		bytePos := bytePositions[n]
		corrupted[bytePos] = struct{}{}

		var minSteps int
		for len(queue) > 0 {
			curr := queue[0]
			pos := [2]int{curr[X], curr[Y]}
			queue = queue[1:]

			if _, ok := seen[pos]; ok {
				continue
			}

			if _, ok := corrupted[pos]; ok {
				continue
			}

			if pos[X] < 0 || pos[X] >= gridSize || pos[Y] < 0 || pos[Y] >= gridSize {
				continue
			}

			if pos == end {
				minSteps = curr[Steps]
				break
			}

			for _, dir := range directions {
				queue = append(queue, [3]int{curr[X] + dir[X], curr[Y] + dir[Y], curr[Steps] + 1})
			}

			seen[pos] = struct{}{}
		}

		// we never reached the end square
		if minSteps == 0 {
			blockingByte = bytePos
			break
		}
	}

	return fmt.Sprintf("%d,%d", blockingByte[X], blockingByte[Y])
}
