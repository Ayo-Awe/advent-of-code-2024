package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var (
	r = regexp.MustCompile(`\d+`)
)

type Equation struct {
	Target   int
	Operands []int
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var equations []Equation
	for _, line := range input {
		nums := []int{}

		for _, str := range r.FindAllString(line, -1) {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		equations = append(equations, Equation{
			Target:   nums[0],
			Operands: nums[1:],
		})
	}

	fmt.Println("solution to part one: ", partOne(equations))
}

func partOne(eqns []Equation) int {
	sum := 0
	for _, eqn := range eqns {
		if evalOne(eqn.Target, eqn.Operands) {
			sum += eqn.Target
		}
	}
	return sum
}

func evalOne(target int, operands []int) bool {
	lastIndex := len(operands) - 1

	// no elements left
	if lastIndex < 0 {
		return false
	}

	operand := operands[lastIndex]

	// only one element left
	if lastIndex == 0 {
		return operand == target
	}

	// eval * operator only if target is divisible by operand
	multiplySuccess := false
	if target%operand == 0 {
		multiplySuccess = evalOne(target/operand, operands[:lastIndex])
	}

	// eval + operator
	return multiplySuccess || evalOne(target-operand, operands[:lastIndex])
}
