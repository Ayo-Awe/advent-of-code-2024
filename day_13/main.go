package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var (
	r = regexp.MustCompile("\\d+")
)

type Game struct {
	a struct {
		x, y int
	}
	b struct {
		x, y int
	}
	prize struct {
		x, y int
	}
}

func parseInput(rows []string) ([]Game, error) {
	var games []Game
	var curr Game
	for i, row := range rows {
		// empty row
		if i%4 == 3 {
			games = append(games, curr)
			continue
		}

		digits := r.FindAllString(row, -1)
		if len(digits) != 2 {
			return nil, fmt.Errorf("expected 2 digits but got %d", len(digits))
		}

		x, err := strconv.Atoi(digits[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(digits[1])
		if err != nil {
			return nil, err
		}

		switch i % 4 {
		case 0:
			curr.a.x, curr.a.y = x, y
		case 1:
			curr.b.x, curr.b.y = x, y
		case 2:
			curr.prize.x, curr.prize.y = x, y
		}
	}

	games = append(games, curr)
	return games, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(games))
	fmt.Println("solution to part two: ", PartTwo(games))
}

func PartOne(games []Game) int {
	var tokens int
	for _, game := range games {
		tokens += solve(game)
	}
	return tokens
}

func PartTwo(games []Game) int {
	var tokens int
	for _, game := range games {
		game.prize.x += 10000000000000
		game.prize.y += 10000000000000
		tokens += solve(game)
	}
	return tokens
}

func solve(game Game) int {

	matM := [2][2]float64{
		{float64(game.a.x), float64(game.b.x)},
		{float64(game.a.y), float64(game.b.y)},
	}
	matN := [2]float64{
		float64(game.prize.x),
		float64(game.prize.y),
	}

	// calculate solution
	// X = Minv * N
	solution := [2]float64{}
	matMInverse := inverse(matM)
	for i := range len(matMInverse) {
		solution[i] = matMInverse[i][0]*matN[0] + matMInverse[i][1]*matN[1]
	}

	a := math.Round(solution[0])
	b := math.Round(solution[1])

	// handle rounding errors
	aDiff := a - solution[0]
	bDiff := b - solution[1]

	if math.Abs(aDiff) > 1e-4 || math.Abs(bDiff) > 1e-4 {
		return 0
	}

	return int(a)*3 + int(b)*1
}

func inverse(mat [2][2]float64) [2][2]float64 {
	var matInverse [2][2]float64

	det := mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]

	for i := range len(mat) {
		for j := range len(mat[0]) {
			// multiply lagging diagonal by -1
			if i != j {
				matInverse[i][j] = -1 * mat[i][j]
			} else {
				ii, jj := (i+1)%2, (j+1)%2
				matInverse[i][j] = mat[ii][jj]
			}

			matInverse[i][j] /= det
		}
	}

	return matInverse
}
