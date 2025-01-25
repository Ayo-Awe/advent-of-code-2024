package main

import (
	"fmt"
	"log"
	"math"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type pos [2]int

func (p pos) x() int { return p[0] }
func (p pos) y() int { return p[1] }

type direction int

func (d direction) clockwise() direction {
	return (d + 1) % 4
}

func (d direction) anitClockwise() direction {
	return (d + 3) % 4
}

func (d direction) delta() [2]int {
	switch d {
	case North:
		return [2]int{0, -1}
	case South:
		return [2]int{0, 1}
	case East:
		return [2]int{1, 0}
	case West:
		return [2]int{-1, 0}
	default:
		return [2]int{}
	}
}

const (
	// the order matters for the rotate function
	North direction = iota
	East
	South
	West
)

func parseInput(input []string) (grid [][]rune) {
	for _, row := range input {
		grid = append(grid, []rune(row))
	}
	return
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := parseInput(input)
	fmt.Println("solution to part one: ", PartOne(grid))
	fmt.Println("solution to part two: ", PartTwo(grid))
}

func PartOne(grid [][]rune) int {
	type state struct {
		pos
		score int
		dir   direction
	}

	start := pos{1, len(grid) - 2}
	seen := make(map[string]int)
	queue := []state{{start, 0, East}}
	minScore := math.MaxInt

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.score > minScore {
			continue
		}

		// we differentiate states only by their positions and directions. The score is ignored
		stateKey := newStateKey(currState.pos, currState.dir)
		if prevScore, ok := seen[stateKey]; ok && prevScore < currState.score {
			continue
		}

		if grid[currState.y()][currState.x()] == 'E' {
			minScore = currState.score
			continue
		}

		// forward
		delta := currState.dir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x() + delta[0], currState.y() + delta[1]}
			queue = append(queue, state{newPos, currState.score + 1, currState.dir})
		}

		// only turn clockwise if there's an open square to your right
		newDir := currState.dir.clockwise()
		delta = newDir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x(), currState.y()}
			queue = append(queue, state{newPos, currState.score + 1000, newDir})
		}

		// only turn anit-clockwise if there's an open square to your right
		newDir = currState.dir.anitClockwise()
		delta = newDir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x(), currState.y()}
			queue = append(queue, state{newPos, currState.score + 1000, newDir})
		}

		seen[stateKey] = currState.score
	}

	return minScore
}

func PartTwo(grid [][]rune) int {
	type state struct {
		pos
		score int
		dir   direction
		path  []pos
	}

	var minScoreTiles []pos
	start := pos{1, len(grid) - 2}
	seen := make(map[string]int)
	queue := []state{{start, 0, East, nil}}
	minScore := math.MaxInt

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.score > minScore {
			continue
		}

		// we differentiate states only by their positions and directions. The score and path are ignored
		stateKey := newStateKey(currState.pos, currState.dir)
		if score, ok := seen[stateKey]; ok && score < currState.score {
			continue
		}

		if grid[currState.y()][currState.x()] == 'E' {
			// reset the seen min score tiles
			if currState.score < minScore {
				minScoreTiles = nil
			}

			minScoreTiles = append(minScoreTiles, currState.path...)
			minScore = currState.score
			continue
		}

		// forward
		delta := currState.dir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x() + delta[0], currState.y() + delta[1]}
			path := append([]pos{currState.pos}, currState.path...)
			queue = append(queue, state{newPos, currState.score + 1, currState.dir, path})
		}

		// only explore clockwise if there's an open square to your right
		newDir := currState.dir.clockwise()
		delta = newDir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x(), currState.y()}
			path := append([]pos{currState.pos}, currState.path...)
			queue = append(queue, state{newPos, currState.score + 1000, newDir, path})
		}

		// only explore anti-clockwise if there's an open square to your left
		newDir = currState.dir.anitClockwise()
		delta = newDir.delta()
		if grid[currState.y()+delta[1]][currState.x()+delta[0]] != '#' {
			newPos := pos{currState.x(), currState.y()}
			path := append([]pos{currState.pos}, currState.path...)
			queue = append(queue, state{newPos, currState.score + 1000, newDir, path})
		}

		seen[stateKey] = currState.score
	}

	uniqueTiles := make(map[[2]int]struct{})
	for _, tile := range minScoreTiles {
		uniqueTiles[tile] = struct{}{}
	}

	// we add 1 to account for the end tile
	return len(uniqueTiles) + 1
}

func newStateKey(p pos, d direction) string {
	return fmt.Sprintf("%d:%d:%d", p.x(), p.y(), d)
}
