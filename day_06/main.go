package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var gridArea GridArea
	for _, line := range input {
		gridArea = append(gridArea, strings.Split(line, ""))
	}

	fmt.Println("solution to part one: ", PartOne(gridArea))
	fmt.Println("solution to part two: ", PartTwo(gridArea))
}

func PartOne(gridArea GridArea) int {

	guard := gridArea.FindGuard()
	if guard == nil {
		log.Fatal("invalid input")
	}
	exploredPositions := make(map[string]struct{})

	for {
		// mark current position as seen
		key := fmt.Sprintf("%d,%d", guard.x, guard.y)
		exploredPositions[key] = struct{}{}

		// get coordinates of the next position
		x, y := guard.LookAhead()

		// next position is out of bounds
		if gridArea.IsOutOfBounds(x, y) {
			break

			// next position is blocked
		} else if gridArea[y][x] == "#" {
			guard.Rotate90Clockwise()

			// otherwise move forward
		} else {
			guard.Forward()
		}
	}

	return len(exploredPositions)
}

func PartTwo(gridArea GridArea) int {
	loops := 0
	for y, row := range gridArea {
		for x, pos := range row {
			if isGuard(pos) || pos == "#" {
				continue
			}

			// replace pos with an obstacle and test for loop
			newGridArea := gridArea.Clone()
			newGridArea[y][x] = "#"

			if checkForLoop(newGridArea) {
				loops++
			}
		}
	}

	return loops
}

func checkForLoop(g GridArea) bool {
	guard := g.FindGuard()
	if guard == nil {
		return false
	}

	seenTurns := make(map[string]struct{})
	for {
		x, y := guard.LookAhead()

		if g.IsOutOfBounds(x, y) {
			break

			// is a turn
		} else if g[y][x] == "#" {
			key := fmt.Sprintf("%d,%d,%s", guard.x, guard.y, guard.direction)

			// if a guard passes a turn twice, he's in a loop
			_, seen := seenTurns[key]
			if seen {
				return true
			}

			seenTurns[key] = struct{}{}
			guard.Rotate90Clockwise()
			// otherwise move forward
		} else {
			guard.Forward()
		}
	}

	return false
}
