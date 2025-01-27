package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var towelRegexp = regexp.MustCompile(`[rgbuw]+`)

func parseInput(input []string) (towels, patterns []string) {
	towels = towelRegexp.FindAllString(input[0], -1)
	patterns = input[2:]
	return
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	towels, patterns := parseInput(input)
	fmt.Println("solution to part one: ", PartOne(towels, patterns))
	fmt.Println("solution to part two: ", PartTwo(towels, patterns))
}

func PartOne(towels, patterns []string) int {
	var totalPatternable int
	memo := make(map[string]int)

	for _, pattern := range patterns {
		if patternable(towels, pattern, memo) > 0 {
			totalPatternable++
		}
	}
	return totalPatternable
}

func PartTwo(towels, patterns []string) int {
	memo := make(map[string]int)
	var totalPatternable int
	for _, pattern := range patterns {
		totalPatternable += patternable(towels, pattern, memo)
	}
	return totalPatternable
}

func patternable(towels []string, pattern string, memo map[string]int) int {
	if pattern == "" {
		return 1
	}

	if options, exists := memo[pattern]; exists {
		return options
	}

	var totalOptions int
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			subPattern := strings.TrimPrefix(pattern, towel)
			totalOptions += patternable(towels, subPattern, memo)
		}
	}

	memo[pattern] = totalOptions

	return totalOptions
}
