package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

type Computer struct {
	IP      int
	A, B, C int
	Out     []int
}

func parseInput(input []string) (Computer, []int, error) {
	if len(input) != 5 {
		return Computer{}, nil, fmt.Errorf("invalid input: expected 5 lines but got %d", len(input))
	}

	var comp Computer
	var program []int

	for i, line := range input {
		switch i {
		case 0:
			A, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				return Computer{}, nil, err
			}
			comp.A = A

		case 1:
			B, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				return Computer{}, nil, err
			}
			comp.B = B

		case 2:
			C, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				return Computer{}, nil, err
			}
			comp.C = C
		case 4:
			programStr := strings.Split(line, ": ")[1]
			for _, instr := range strings.Split(programStr, ",") {
				a, err := strconv.Atoi(instr)
				if err != nil {
					return Computer{}, nil, err
				}
				program = append(program, a)
			}
		}
	}

	return comp, program, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	computer, program, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(computer, program))
}

func PartOne(computer Computer, program []int) string {
	curr := computer
	for curr.IP >= 0 && curr.IP < len(program) {
		opcode, operand := program[curr.IP], program[curr.IP+1]
		exec(opcode, operand, &curr)
	}

	var res []string
	for _, out := range curr.Out {
		res = append(res, strconv.Itoa(out))
	}

	return strings.Join(res, ",")
}

func combo(operand int, computer *Computer) int {
	switch operand {
	case 4:
		return computer.A
	case 5:
		return computer.B
	case 6:
		return computer.C
	default:
		return operand
	}
}

func exec(opCode, operand int, computer *Computer) {
	var jumped bool

	switch opCode {
	case 0:
		numerator := computer.A
		comboOp := combo(operand, computer)
		computer.A = numerator / int(math.Pow(2, float64(comboOp)))
	case 1:
		computer.B = computer.B ^ operand
	case 2:
		computer.B = combo(operand, computer) % 8
	case 3:
		if computer.A != 0 {
			jumped = true
			computer.IP = operand
		}
	case 4:
		computer.B = computer.B ^ computer.C
	case 5:
		comboOp := combo(operand, computer) % 8
		computer.Out = append(computer.Out, comboOp)
	case 6:
		numerator := computer.A
		comboOp := combo(operand, computer)
		computer.B = numerator / int(math.Pow(2, float64(comboOp)))
	case 7:
		numerator := computer.A
		comboOp := combo(operand, computer)
		computer.C = numerator / int(math.Pow(2, float64(comboOp)))
	}

	if !jumped {
		computer.IP += 2
	}
}
