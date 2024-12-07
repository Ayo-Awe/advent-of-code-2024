package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func main() {
	var input [][]string

	lines, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		input = append(input, strings.Split(line, ""))
	}

	fmt.Println("solution to part one:", partOne(input))
}
