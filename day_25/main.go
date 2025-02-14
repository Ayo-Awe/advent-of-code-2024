package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func parseInput(input string) (keys, locks [][5]int) {
	blocks := strings.Split(input, "\n\n")
	for _, block := range blocks {
		rows := strings.Split(strings.TrimSuffix(block, "\n"), "\n")

		var isKey bool
		var value [5]int

		for i, row := range rows {
			// it's a lock if the first row is full
			if i == 0 {
				if row == "#####" {
					isKey = false
				} else {
					isKey = true
				}
			}

			// skip first and last rows
			if i == 0 || i == len(rows)-1 {
				continue
			}

			for j, cell := range row {
				if cell == '#' {
					value[j]++
				}
			}
		}

		if isKey {
			keys = append(keys, value)
		} else {
			locks = append(locks, value)
		}
	}

	return
}

func main() {
	input, err := aoc.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	keys, locks := parseInput(input)
	fmt.Println(keys[0], locks[0])

	fmt.Println("solution to part one: ", PartOne(keys, locks))
	fmt.Println("solution to part two: ", PartTwo())
}

func PartOne(keys, locks [][5]int) int {
	var matches int
	for _, lock := range locks {
		for _, key := range keys {
			for col := range key {
				if key[col]+lock[col] > 5 {
					break
				}

				// no overlaps
				if col == len(key)-1 {
					matches++
				}
			}
		}
	}

	return matches
}

func PartTwo() int {
	return 0
}

