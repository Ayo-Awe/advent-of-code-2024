package main

import (
	"fmt"
	"log"
	"math"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var directions = [][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

var (
	X = 0
	Y = 1
)

func parseInput(input []string) (grid [][]rune, start, end [2]int) {
	for y, row := range input {
		for x, cell := range row {
			switch cell {
			case 'S':
				start = [2]int{x, y}
			case 'E':
				end = [2]int{x, y}
			}
		}
		grid = append(grid, []rune(row))
	}

	return
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid, start, end := parseInput(input)
	fmt.Println("solution to part one: ", PartOne(grid, start, end))
	fmt.Println("solution to part two: ", PartTwo(grid, start, end))
}

func PartOne(grid [][]rune, start, end [2]int) int {
	path := plotPath(grid, start, end)

	var count int
	for i, pos := range path[:len(path)-102] {
		for _, target := range path[102+i:] {
			dx, dy := pos[X]-target[X], pos[Y]-target[Y]

			if dx != 0 && dy != 0 {
				continue
			}

			if math.Abs(float64(dx)) == 2 || math.Abs(float64(dy)) == 2 {
				count++
			}
		}
	}

	return count
}

func PartTwo(grid [][]rune, start, end [2]int) int {
	path := plotPath(grid, start, end)
	minSavings := 100

	// it's impossible to jump a 100 squares with 0 steps and
	// if the squares were only 1 step apart the grid would be invalid
	// as there's only one path to the end
	minCheatSteps := 2
	maxCheatSteps := 20

	// this is the minimum amount of steps between two squares required to reach
	// our savings goals
	minStepsApart := minSavings + minCheatSteps

	var numberOfCheats int
	for i, pos := range path {
		for j := minStepsApart + i; j < len(path); j++ {
			target := path[j]
			dx, dy := pos[X]-target[X], pos[Y]-target[Y]

			cheatSteps := int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
			originalSteps := j - i
			savings := originalSteps - cheatSteps

			if cheatSteps > maxCheatSteps {
				continue
			}

			if savings >= minSavings {
				numberOfCheats++
			}
		}
	}

	return numberOfCheats
}

func plotPath(grid [][]rune, start, end [2]int) [][2]int {
	var path [][2]int
	curr := start

	for curr != end {
		// find the next open position and take it
		var nextPos [2]int

		for _, dir := range directions {
			newPos := [2]int{curr[X] + dir[X], curr[Y] + dir[Y]}

			// ensure we don't go backwards in the path
			if len(path) > 0 && newPos == path[len(path)-1] {
				continue
			}

			if grid[newPos[Y]][newPos[X]] == '#' {
				continue
			}

			nextPos = newPos
			break
		}

		path = append(path, curr)
		curr = nextPos
	}
	path = append(path, end)

	return path
}
