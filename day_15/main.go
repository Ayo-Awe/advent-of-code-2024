package main

import (
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type Warehouse struct {
	grid   [][]rune
	player [2]int
	moves  []rune
}

var (
	robotSymbol = '@'
	X           = 0
	Y           = 1
	directions  = map[rune][2]int{
		'^': {0, -1},
		'v': {0, 1},
		'>': {1, 0},
		'<': {-1, 0},
	}
)

func parseInput(lines []string) *Warehouse {
	inGridSection := true
	var warehouse Warehouse

	for y, line := range lines {
		if len(line) == 0 {
			inGridSection = false
		}

		if inGridSection {
			warehouse.grid = append(warehouse.grid, []rune(line))

			// find robot position
			for x, c := range line {
				if c == robotSymbol {
					warehouse.player = [2]int{x, y}
				}
			}
		} else {
			warehouse.moves = append(warehouse.moves, []rune(line)...)
		}
	}

	return &warehouse
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	warehouse := parseInput(input)
	fmt.Println("solution to part one: ", PartOne(*warehouse))
}

func PartOne(warehouse Warehouse) int {
	for _, move := range warehouse.moves {
		dir := directions[move]

		boxes := countBoxes(warehouse.player, dir, warehouse.grid)

		// check for obstruction at the end of the box train
		targetPos := stepN(warehouse.player, dir, boxes+1)
		if warehouse.grid[targetPos[Y]][targetPos[X]] == '#' {
			continue
		}

		// move entire train in direction dir (i.e robot and boxes) starting from the last item
		for i := boxes; i >= 0; i-- {
			// swap with the cell in front of it
			pos := stepN(warehouse.player, dir, i)
			swap(pos, stepN(pos, dir, 1), warehouse.grid)
		}

		// update the robot's position to the next square
		warehouse.player = stepN(warehouse.player, dir, 1)
	}

	var sum int
	for y, row := range warehouse.grid {
		for x, cell := range row {
			if cell == 'O' {
				sum += y*100 + x
			}
		}
	}

	return sum
}

// counts consecutive boxes after position pos in direction dir
func countBoxes(pos, dir [2]int, grid [][]rune) int {
	var count int
	currX, currY := pos[X]+dir[X], pos[Y]+dir[Y]

	for grid[currY][currX] == 'O' {
		count++

		// update position
		currX += dir[X]
		currY += dir[Y]
	}

	return count
}

func swap(a, b [2]int, grid [][]rune) {
	grid[a[Y]][a[X]], grid[b[Y]][b[X]] = grid[b[Y]][b[X]], grid[a[Y]][a[X]]
}

// take n steps in a given direct
func stepN(pos, dir [2]int, n int) [2]int {
	return [2]int{pos[X] + dir[X]*n, pos[Y] + dir[Y]*n}
}
