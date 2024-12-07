package main

import (
	"slices"
)

func partTwo(input [][]string) int {
	totalMatches := 0
	for y, row := range input {
		for x, c := range row {
			if c == "A" && checkForX_MAS(input, x, y) {
				totalMatches++
			}
		}
	}

	return totalMatches
}

func checkForX_MAS(input [][]string, x, y int) bool {
	topBound := 0
	bottomBound := len(input) - 1
	rightBound := len(input[y]) - 1
	leftBound := 0

	leftDiagonalVal := ""
	rightDiagonalVal := ""

	for i := -1; i <= 1; i++ {
		rx, ry := x+i, y-i
		lx, ly := x-i, y-i

		outOfBounds := rx > rightBound || rx < leftBound ||
			lx > rightBound || lx < leftBound ||
			ry > bottomBound || ry < topBound ||
			ly > bottomBound || ly < topBound

		if outOfBounds {
			return false
		}

		leftDiagonalVal += input[ly][lx]
		rightDiagonalVal += input[ry][rx]
	}

	return isMAS(leftDiagonalVal) && isMAS(rightDiagonalVal)
}

func isMAS(s string) bool {
	reversed := []rune(s)
	slices.Reverse(reversed)

	return string(reversed) == "MAS" || s == "MAS"
}
