package main

import (
	"fmt"
	"log"
	"math"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type pos [3]int

func (p pos) x() int         { return p[0] }
func (p pos) y() int         { return p[1] }
func (p pos) dir() direction { return direction(p[2]) }

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
	fmt.Println(PartOne(grid))
}

type state struct {
	pos
	score int
}

func PartOne(grid [][]rune) int {
	start := pos{1, len(grid) - 2, int(East)}
	seen := make(map[pos]int)
	queue := []state{{start, 0}}
	minScore := math.MaxInt

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.score > minScore {
			continue
		}

		if score, ok := seen[curr.pos]; ok && score < curr.score {
			continue
		}

		if grid[curr.y()][curr.x()] == 'E' {
			minScore = curr.score
			continue
		}

		// forward
		delta := curr.pos.dir().delta()
		if grid[curr.y()+delta[1]][curr.x()+delta[0]] != '#' {
			newPos := pos{curr.x() + delta[0], curr.y() + delta[1], int(curr.dir())}
			queue = append(queue, state{newPos, curr.score + 1})
		}

		// only turn clockwise if there's an open square to your right
		delta = curr.dir().clockwise().delta()
		if grid[curr.y()+delta[1]][curr.x()+delta[0]] != '#' {
			newPos := pos{curr.x(), curr.y(), int(curr.dir().clockwise())}
			queue = append(queue, state{newPos, curr.score + 1000})
		}

		delta = curr.dir().anitClockwise().delta()
		if grid[curr.y()+delta[1]][curr.x()+delta[0]] != '#' {
			newPos := pos{curr.x(), curr.y(), int(curr.dir().anitClockwise())}
			queue = append(queue, state{newPos, curr.score + 1000})
		}

		seen[curr.pos] = curr.score
	}

	return minScore
}
