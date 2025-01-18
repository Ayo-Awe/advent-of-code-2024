package main

import (
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type Warehouse struct {
	grid  [][]rune
	robot [2]int
	moves []rune
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
					warehouse.robot = [2]int{x, y}
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

	warehouse = parseInput(input)
	fmt.Println("solution to part two: ", PartTwo(*warehouse))
}

func PartOne(warehouse Warehouse) int {
	for _, move := range warehouse.moves {
		dir := directions[move]

		boxes := countBoxes(warehouse.robot, dir, warehouse.grid)

		// check for obstruction at the end of the box train
		targetPos := stepN(warehouse.robot, dir, boxes+1)
		if warehouse.grid[targetPos[Y]][targetPos[X]] == '#' {
			continue
		}

		// move entire train in direction dir (i.e robot and boxes) starting from the last item
		for i := boxes; i >= 0; i-- {
			// swap with the cell in front of it
			pos := stepN(warehouse.robot, dir, i)
			swap(pos, stepN(pos, dir, 1), warehouse.grid)
		}

		// update the robot's position to the next square
		warehouse.robot = stepN(warehouse.robot, dir, 1)
	}

	return GPSCoordSum(warehouse.grid)
}

func PartTwo(warehouse Warehouse) int {
	warehouse.grid, warehouse.robot = expand(warehouse.grid)

	for _, m := range warehouse.moves {
		dir := directions[m]

		// horizontal move
		if dir[Y] == 0 {
			// distance between the robot's position and
			// the last consecutive box cell in direction dir
			var dist int
			curr := stepN(warehouse.robot, dir, 1)
			for warehouse.grid[curr[Y]][curr[X]] == '[' || warehouse.grid[curr[Y]][curr[X]] == ']' {
				dist++
				curr = stepN(curr, dir, 1)
			}

			target := stepN(warehouse.robot, dir, dist+1)
			if warehouse.grid[target[Y]][target[X]] == '#' {
				continue
			}

			move(warehouse.robot, dir, warehouse.grid)
			warehouse.robot = stepN(warehouse.robot, dir, 1)
		} else {
			var isBlocked bool
			queue := [][2]int{stepN(warehouse.robot, dir, 1)}
			seen := map[[2]int]struct{}{}

			for len(queue) > 0 {
				pos := queue[0]
				queue = queue[1:]

				if _, exists := seen[pos]; exists {
					continue
				}

				cell := warehouse.grid[pos[Y]][pos[X]]
				if cell == '[' {
					// append matching bracket
					queue = append(queue, [2]int{pos[X] + 1, pos[Y]})
				} else if cell == ']' {
					// append matching bracket
					queue = append(queue, [2]int{pos[X] - 1, pos[Y]})
				} else if cell == '#' {
					isBlocked = true
					break
				} else {
					continue
				}

				// add the cell in front of the current to the queue
				queue = append(queue, stepN(pos, dir, 1))
				seen[pos] = struct{}{}
			}

			// skip there's a wall in front of at least one of the boxes
			if isBlocked {
				continue
			}

			// move all boxes forward
			move(warehouse.robot, dir, warehouse.grid)
			warehouse.robot = stepN(warehouse.robot, dir, 1)
		}
	}

	return GPSCoordSum(warehouse.grid)
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

// expands grid based on part two specs
func expand(grid [][]rune) ([][]rune, [2]int) {
	var expanded [][]rune

	var robot [2]int
	for y, row := range grid {
		var expandedRow []rune
		for x, cell := range row {
			switch cell {
			case 'O':
				expandedRow = append(expandedRow, '[', ']')
			case '@':
				expandedRow = append(expandedRow, '@', '.')
				robot = [2]int{2 * x, y}
			default:
				expandedRow = append(expandedRow, cell, cell)
			}
		}
		expanded = append(expanded, expandedRow)
	}

	return expanded, robot
}

// move blocks and robot up/down in part two
func move(cellPos, dir [2]int, grid [][]rune) {
	// skip this cell if it's an empty square
	cell := grid[cellPos[Y]][cellPos[X]]
	if cell == '.' {
		return
	}

	// move the cell in front first, then attempt to move this cell
	target := stepN(cellPos, dir, 1)
	move(target, dir, grid)
	swap(cellPos, target, grid)

	// if this is a veritical move, move the matching bracket too
	if dir[Y] != 0 {
		if cell == '[' {
			move([2]int{cellPos[X] + 1, cellPos[Y]}, dir, grid)
		} else if cell == ']' {
			move([2]int{cellPos[X] - 1, cellPos[Y]}, dir, grid)
		}
	}
}

func GPSCoordSum(grid [][]rune) int {
	var sum int
	for y, row := range grid {
		for x, cell := range row {
			if cell == '[' || cell == 'O' {
				sum += y*100 + x
			}
		}
	}

	return sum
}
