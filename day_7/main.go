package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

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
	fmt.Println("solution to part two: ", partTwo(equations))
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

func partTwo(eqns []Equation) int {
	sum := 0
	for _, eqn := range eqns {
		if evalTwo(eqn.Target, eqn.Operands) {
			sum += eqn.Target
		}
	}
	return sum
}

func evalTwo(target int, operands []int) bool {
	lastIndex := len(operands) - 1

	// no elements left
	if lastIndex < 0 {
		return false
	}

	if target < 0 {
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
		multiplySuccess = evalTwo(target/operand, operands[:lastIndex])
	}

	// eval || operator only if target ends with operand
	concatSuccess := false
	if endsWithInt(target, operand) {
		newTarget := (target - operand) / int(math.Pow10(numberOfDigits(operand)))
		concatSuccess = evalTwo(newTarget, operands[:lastIndex])
	}

	// eval + operator
	return multiplySuccess || concatSuccess || evalTwo(target-operand, operands[:lastIndex])
}

func endsWithInt(target, val int) bool {
	return strings.HasSuffix(strconv.Itoa(target), strconv.Itoa(val))
}

func numberOfDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}
