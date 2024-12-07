package main

func partOne(input [][]string) int {
	totalMatches := 0
	for y, row := range input {
		for x, c := range row {
			if c == "X" {
				totalMatches += evaluatePosition(input, x, y)
			}
		}
	}

	return totalMatches
}

func checkHorizontalF(input [][]string, x, y int) bool {
	target := "XMAS"
	rightBound := len(input[y]) - 1

	found := ""
	for i := x; i < x+len(target); i++ {
		if i > rightBound {
			break
		}

		found += input[y][i]
	}

	return found == target
}

func checkHorizontalB(input [][]string, x, y int) bool {
	target := "XMAS"
	found := ""

	for i := x; i > x-len(target); i-- {
		if i < 0 {
			break
		}

		found += input[y][i]
	}

	return found == target
}

func checkVerticalF(input [][]string, x, y int) bool {
	target := "XMAS"
	lowerBound := len(input) - 1

	found := ""
	for i := y; i < y+len(target); i++ {
		if i > lowerBound {
			break
		}

		found += input[i][x]
	}

	return found == target
}

func checkVerticalB(input [][]string, x, y int) bool {
	target := "XMAS"
	found := ""
	for i := y; i > y-len(target); i-- {
		if i < 0 {
			break
		}

		found += input[i][x]
	}

	return found == target
}
func checkRightDiagonalF(input [][]string, x, y int) bool {
	target := "XMAS"
	topBound := 0
	rightBound := len(input[y]) - 1

	found := ""
	for i := 0; i < len(target); i++ {
		tx, ty := x+i, y-i
		if tx > rightBound || ty < topBound {
			break
		}

		found += input[ty][tx]
	}

	return found == target
}

func checkRightDiagonalB(input [][]string, x, y int) bool {
	target := "XMAS"
	bottomBound := len(input) - 1 //
	leftBound := 0

	found := ""
	for i := 0; i < len(target); i++ {
		tx, ty := x-i, y+i
		if tx < leftBound || ty > bottomBound {
			break
		}

		found += input[ty][tx]
	}

	return found == target
}

func checkLeftDiagonalF(input [][]string, x, y int) bool {
	target := "XMAS"
	bottomBound := len(input) - 1
	rightBound := len(input[y]) - 1

	found := ""
	for i := 0; i < len(target); i++ {
		tx, ty := x+i, y+i
		if tx > rightBound || ty > bottomBound {
			break
		}

		found += input[ty][tx]
	}

	return found == target
}

func checkLeftDiagonalB(input [][]string, x, y int) bool {
	target := "XMAS"
	bottomBound := 0
	rightBound := 0

	found := ""
	for i := 0; i < len(target); i++ {
		tx, ty := x-i, y-i
		if tx < rightBound || ty < bottomBound {
			break
		}

		found += input[ty][tx]
	}

	return found == target
}

// returns number of matches found at the given postion
func evaluatePosition(input [][]string, x, y int) int {
	evaluators := []func([][]string, int, int) bool{
		checkHorizontalF,
		checkHorizontalB,
		checkLeftDiagonalF,
		checkLeftDiagonalB,
		checkRightDiagonalB,
		checkRightDiagonalF,
		checkVerticalB,
		checkVerticalF,
	}

	matches := 0
	for _, e := range evaluators {
		if e(input, x, y) {
			matches++
		}
	}

	return matches
}
