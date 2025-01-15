package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type robot struct {
	pos [2]int
	vel [2]int
}

func (r robot) newPos(dimensions [2]int, sec int) [2]int {
	w, h := dimensions[0], dimensions[1]

	nx, ny := r.pos[0]+r.vel[0]*sec, r.pos[1]+r.vel[1]*sec

	// adjust new postion to grid dimensions
	nx, ny = ((nx%w)+w)%w, ((ny%h)+h)%h

	return [2]int{nx, ny}
}

var (
	r = regexp.MustCompile("-?\\d+")
)

func parseInput(rows []string) ([]robot, error) {
	var robots []robot
	for i, row := range rows {
		var digits []int
		for _, d := range r.FindAllString(row, -1) {
			digit, err := strconv.Atoi(d)
			if err != nil {
				return nil, err
			}
			digits = append(digits, digit)
		}

		if len(digits) != 4 {
			return nil, fmt.Errorf("expected 4 digits on line %d, but got %d", i+1, len(digits))
		}

		rbt := robot{
			pos: [2]int(digits[:2]),
			vel: [2]int(digits[2:]),
		}
		robots = append(robots, rbt)
	}
	return robots, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(robots))
}

func PartOne(robots []robot) int {
	dim := [2]int{101, 103}
	sec := 100

	quadrants := make(map[[2]int]int)
	for _, r := range robots {
		pos := r.newPos(dim, sec)
		x, y := pos[0], pos[1]

		midX, midY := (dim[0]-1)/2, (dim[1]-1)/2
		if x == midX || y == midY {
			continue
		}

		quadX, quadY := x/(midX+1), y/(midY+1)
		quadrants[[2]int{quadX, quadY}]++
	}

	total := 1
	for _, count := range quadrants {
		total *= count
	}

	return total
}
