package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readAndParseInput(filename string) ([][]int, error) {
	reports := [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, err
		}

		report := []int{}
		for _, c := range strings.Split(scanner.Text(), " ") {
			level, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func main() {
	input, err := readAndParseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one is: ", partOne(input))
	fmt.Println("solution to part two is: ", partTwo(input))
}

func partOne(input [][]int) int {
	count := 0
	for _, report := range input {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func partTwo(input [][]int) int {
	count := 0
	for _, report := range input {
		if isSafeWithTolerance(report) {
			count++
		}
	}
	return count
}

func isSafe(report []int) bool {
	prevIsDesc := report[0] > report[1]

	for i := range len(report) - 1 {
		absDiff := abs(report[i] - report[i+1])

		if absDiff > 3 || absDiff < 1 {
			return false
		}

		isDesc := report[i] > report[i+1]
		if isDesc != prevIsDesc {
			return false
		}

		prevIsDesc = isDesc
	}

	return true
}

func isSafeWithTolerance(report []int) bool {

	for i := range len(report) {
		isSafeWithoutIndex := isSafe(omitIndex(report, i))
		if isSafeWithoutIndex {
			return true
		}
	}

	return false
}

func omitIndex(s []int, index int) []int {
	res := []int{}
	for i, val := range s {
		if i == index {
			continue
		}

		res = append(res, val)
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
