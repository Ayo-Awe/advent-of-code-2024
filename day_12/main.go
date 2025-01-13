package main

import (
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var directions = [][2]int{
	// x, y
	{0, 1},  // down
	{1, 0},  // right
	{0, -1}, // up
	{-1, 0}, // left
}

var internalCorners = [][2][2]int{
	{
		{1, 0}, // right
		{0, 1}, // down
	},
	{
		{1, 0},  // right
		{0, -1}, // top
	},
	{
		{-1, 0}, // left
		{0, -1}, // top
	},
	{
		{-1, 0}, // left
		{0, 1},  // bottom
	},
}

func parseInput(input []string) ([][]rune, error) {
	var parsedInput [][]rune
	for _, row := range input {
		parsedInput = append(parsedInput, []rune(row))
	}
	return parsedInput, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	garden, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(garden))
	fmt.Println("solution to part one: ", ParTwo(garden))
}

func PartOne(plots [][]rune) int {
	seen := make(map[[2]int]bool)

	var price int
	for y := range plots {
		for x := range plots[y] {
			price += explore1(x, y, plots, seen)
		}
	}

	return price
}

func ParTwo(plots [][]rune) int {
	seen := make(map[[2]int]bool)

	var total int
	for y := range plots {
		for x := range plots[y] {
			price := explore2(x, y, plots, seen)
			total += price
			if price > 0 {
				fmt.Println(x, y, price)
			}
		}
	}

	return total
}

func isOutOfBounds(x, y, height, width int) bool {
	return x < 0 || x >= width || y < 0 || y >= height
}

func explore1(x, y int, plots [][]rune, seen map[[2]int]bool) int {
	var area, perimeter int
	width := len(plots[0])
	height := len(plots)

	if seen[[2]int{x, y}] {
		return 0
	}

	// intialise queue
	queue := [][2]int{{x, y}}

	// run until queue is empty
	for len(queue) > 0 {
		// pop first element
		curr := queue[0]
		queue = queue[1:]

		if seen[curr] {
			continue
		}

		for _, dir := range directions {
			nx, ny := curr[0]+dir[0], curr[1]+dir[1]

			if isOutOfBounds(nx, ny, height, width) || plots[ny][nx] != plots[y][x] {
				perimeter++
			} else {
				queue = append(queue, [2]int{nx, ny})
			}
		}

		area++
		seen[curr] = true
	}

	return area * perimeter
}

func explore2(x, y int, plots [][]rune, seen map[[2]int]bool) int {
	var area, perimeter int

	width := len(plots[0])
	height := len(plots)

	if seen[[2]int{x, y}] {
		return 0
	}

	// intialise queue
	queue := [][2]int{{x, y}}

	// run until queue is empty
	for len(queue) > 0 {
		// pop first element
		curr := queue[0]
		queue = queue[1:]

		if seen[curr] {
			continue
		}

		sides := [][2]int{}
		for _, dir := range directions {
			nx, ny := curr[0]+dir[0], curr[1]+dir[1]

			if isOutOfBounds(nx, ny, height, width) || plots[ny][nx] != plots[y][x] {
				sides = append(sides, dir)
			} else {
				queue = append(queue, [2]int{nx, ny})
			}
		}

		// count internal corners
		for _, corner := range internalCorners {
			vertical := corner[0]
			horizontal := corner[1]

			diagonalX, diagonalY := curr[0]+vertical[0]+horizontal[0], curr[1]+vertical[1]+horizontal[1]

			if isOutOfBounds(curr[0]+horizontal[0], curr[1]+horizontal[1], height, width) || plots[curr[1]+horizontal[1]][curr[0]+horizontal[0]] != plots[y][x] {
				continue
			}

			if isOutOfBounds(curr[0]+vertical[0], curr[1]+vertical[1], height, width) || plots[curr[1]+vertical[1]][curr[0]+vertical[0]] != plots[y][x] {
				continue
			}

			if plots[diagonalY][diagonalX] == plots[y][x] {
				continue
			}

			fmt.Println(curr, corner)

			perimeter++
		}

		fmt.Println(externalCorners(sides))

		perimeter += externalCorners(sides)
		area++
		seen[curr] = true
	}

	return area * perimeter
}

func externalCorners(sides [][2]int) int {
	// count corners
	switch len(sides) {
	case 4:
		return 4
	case 3:
		return 2
	case 2:
		// sides are perpendicular
		if sides[0][0] != sides[1][0] && sides[0][1] != sides[1][1] {
			return 1
		}

		return 0
	default:
		return 0
	}
}
