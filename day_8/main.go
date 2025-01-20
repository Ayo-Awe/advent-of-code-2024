package main

import (
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type pos [2]int

func (p pos) x() int { return p[0] }
func (p pos) y() int { return p[1] }

func parseInput(input []string) [][]rune {
	var cells [][]rune
	for _, line := range input {
		cells = append(cells, []rune(line))
	}
	return cells
}

func PartOne(grid [][]rune) int {
	antennaMap := make(map[rune][]pos)
	seenAntinodes := make(map[pos]struct{})
	dimensions := [2]int{len(grid[0]), len(grid)}

	for y, row := range grid {
		for x, antenna := range row {
			if antenna == '.' {
				continue
			}

			for _, node := range antennaMap[antenna] {
				if node.y() == y {
					continue
				}

				dx, dy := node.x()-x, node.y()-y

				// upper antinode
				antinodeA := pos{node.x() + dx, node.y() + dy}
				if !antinodeA.isOutOfBounds(dimensions) {
					seenAntinodes[antinodeA] = struct{}{}
				}

				// lower anitode
				antinodeB := pos{node.x() - 2*dx, node.y() - 2*dy}
				if !antinodeB.isOutOfBounds(dimensions) {
					seenAntinodes[antinodeB] = struct{}{}
				}
			}

			antennaMap[antenna] = append(antennaMap[antenna], pos{x, y})
		}
	}

	return len(seenAntinodes)
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := parseInput(input)
	fmt.Println("solution to part one: ", PartOne(grid))
}

func (p pos) isOutOfBounds(dimensions [2]int) bool {
	return p.x() < 0 || p.x() >= dimensions[0] || p.y() < 0 || p.y() >= dimensions[1]
}
