package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func parseInput(input []string) ([][]int, error) {
	var topographicalMap [][]int
	for i := range input {
		var row []int

		for _, c := range input[i] {
			cell, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}
			row = append(row, cell)
		}

		topographicalMap = append(topographicalMap, row)
	}

	return topographicalMap, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	topographicalMap, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(topographicalMap))
	fmt.Println("solution to part two: ", PartTwo(topographicalMap))
}

func PartOne(topoMap [][]int) int {
	var totalScore int

	for y := range topoMap {
		for x := range topoMap[y] {
			// is trailhead
			seenTrails := make(map[string]bool)
			if topoMap[y][x] == 0 {
				score := calculateScore(topoMap, seenTrails, x, y)
				totalScore += score
			}
		}
	}

	return totalScore
}

func PartTwo(topoMap [][]int) int {
	var totalScore int

	for y := range topoMap {
		for x := range topoMap[y] {
			if topoMap[y][x] == 0 {
				score := calculateScore(topoMap, nil, x, y)
				totalScore += score
			}
		}
	}

	return totalScore
}

func calculateScore(topo [][]int, seenTrails map[string]bool, x, y int) int {
	var score int
	curr := topo[y][x]

	// end of trail
	if curr == 9 {
		// trail tracking is disabled
		if seenTrails == nil {
			return 1
		}

		// trail tracking enabled
		if seenTrails[getPosKey(x, y)] {
			return 0
		} else {
			seenTrails[getPosKey(x, y)] = true
			return 1
		}
	}

	// left
	if isPassible(topo, curr, x-1, y) {
		score += calculateScore(topo, seenTrails, x-1, y)
	}

	// right
	if isPassible(topo, curr, x+1, y) {
		score += calculateScore(topo, seenTrails, x+1, y)
	}

	// up
	if isPassible(topo, curr, x, y-1) {
		score += calculateScore(topo, seenTrails, x, y-1)
	}

	// down
	if isPassible(topo, curr, x, y+1) {
		score += calculateScore(topo, seenTrails, x, y+1)
	}

	return score
}

func isWithinBounds(topo [][]int, x, y int) bool {
	// x out of range
	if x > len(topo[0])-1 || x < 0 {
		return false
	}

	// y out of range
	if y > len(topo)-1 || y < 0 {
		return false
	}

	return true
}

func isPassible(topo [][]int, currVal, nextX, nextY int) bool {
	if !isWithinBounds(topo, nextX, nextY) {
		return false
	}

	return topo[nextY][nextX]-currVal == 1
}

func getPosKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
