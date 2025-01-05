package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func parseInput(input string) ([]int, error) {
	input = strings.Trim(input, "\n")

	var stones []int
	for _, strInt := range strings.Split(input, " ") {
		stone, err := strconv.Atoi(strInt)
		if err != nil {
			return nil, err
		}
		stones = append(stones, stone)
	}

	return stones, nil
}

func main() {
	input, err := aoc.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	stones, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(stones))
	fmt.Println("solution to part two: ", PartTwo(stones))
}

func PartOne(stones []int) int {
	newStones := slices.Clone(stones)

	for range 25 {
		newStones = blink(newStones)
	}

	return len(newStones)
}

func PartTwo(stones []int) int {
	var totalCount int
	memo := make(map[string]int)

	for _, stone := range stones {
		totalCount += count(stone, 75, memo)
	}

	return totalCount
}

func blink(stones []int) []int {
	var newStones []int

	for _, stone := range stones {
		newStones = append(newStones, applyRules(stone)...)
	}

	return newStones
}

func applyRules(stone int) []int {
	switch {
	case stone == 0:
		return []int{1}

	case len(strconv.Itoa(stone))%2 == 0:
		left, right := halveDigits(stone)
		return []int{left, right}

	default:
		return []int{stone * 2024}
	}
}

// assumes an even number of digits
func halveDigits(stone int) (int, int) {
	numDigits := nDigits(stone)

	d := int(math.Pow10(numDigits / 2))

	right := stone % d
	left := (stone - right) / d

	return left, right
}

func nDigits(v int) int {
	n := 0
	rem := v

	for rem != 0 {
		rem = rem / 10
		n++
	}

	return n
}

func count(stone, steps int, memo map[string]int) int {
	cachedCount, ok := memo[getMapKey(stone, steps)]
	if ok {
		return cachedCount
	}

	if steps == 0 {
		return 1
	}

	var totalCount int
	for _, s := range blink([]int{stone}) {
		totalCount += count(s, steps-1, memo)
	}

	memo[getMapKey(stone, steps)] = totalCount

	return totalCount
}

func getMapKey(stone, step int) string {
	return fmt.Sprintf("%d-%d", stone, step)
}
