package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Input struct {
	left  []int
	right []int
}

func readAndParseInput(filename string) (*Input, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	left := []int{}
	right := []int{}
	for _, line := range strings.Split(string(bytes), "\n") {
		leftAndRight := regexp.MustCompile("\\d+").FindAllString(line, -1)

		if len(leftAndRight) != 2 {
			continue
		}

		leftInt, err := strconv.Atoi(leftAndRight[0])
		if err != nil {
			return nil, err
		}

		rightInt, err := strconv.Atoi(leftAndRight[1])
		if err != nil {
			return nil, err
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	return &Input{left: left, right: right}, nil
}

func main() {
	input, err := readAndParseInput("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	solutionOne := partOne(input)
	fmt.Println("solution one:", solutionOne)

	solutionTwo := partTwo(input)
	fmt.Println("solution two:", solutionTwo)
}

func partOne(input *Input) int {
	sum := 0
	for i := range len(input.left) {
		sum += diff(input.left[i], input.right[i])
	}

	return sum
}

func partTwo(input *Input) int {
	lookup := map[int]struct{}{}
	for _, val := range input.left {
		lookup[val] = struct{}{}
	}

	// 5 * 3 is the same as 5 + 5 + 5
	similarity := 0
	for _, val := range input.right {
		_, exists := lookup[val]
		if exists {
			similarity += val
		}
	}

	return similarity
}

func diff(left, right int) int {
	sub := left - right

	if sub < 0 {
		return -sub
	}

	return sub
}
