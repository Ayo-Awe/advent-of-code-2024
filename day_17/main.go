package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var (
	rRegisters = regexp.MustCompile(`[A,B,C]:\s\d+`)
	rProgram   = regexp.MustCompile(`Program: (\d+(?:,\d+)*)`)
)

type Computer struct {
	IP      int
	A, B, C int
	Out     []int
}

func (c *Computer) String() string {
	return fmt.Sprintf("Register A: %b\nRegister B: %b\nRegister C: %b", c.A, c.B, c.C)
}

func (c *Computer) Run(program []int) {
	for c.IP >= 0 && c.IP < len(program) {
		opcode, operand := program[c.IP], program[c.IP+1]
		exec(opcode, operand, c)
	}
}

func parseInput(input string) (Computer, []int, error) {
	var comp Computer
	var registers, program []int

	for _, r := range rRegisters.FindAllString(input, -1) {
		rInt, err := strconv.Atoi(strings.Split(r, " ")[1])
		if err != nil {
			return Computer{}, nil, err
		}
		registers = append(registers, rInt)
	}

	for _, v := range strings.Split(rProgram.FindStringSubmatch(input)[1], ",") {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return Computer{}, nil, err
		}
		program = append(program, vInt)
	}

	comp.A = registers[0]
	comp.B = registers[1]
	comp.C = registers[2]

	return comp, program, nil
}

func main() {
	input, err := aoc.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	computer, program, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(computer, program))
	fmt.Println("solution to part two: ", PartTwo(computer, program))
}

func PartOne(comp Computer, program []int) string {
	comp.Run(program)

	var res []string
	for _, out := range comp.Out {
		res = append(res, strconv.Itoa(out))
	}

	return strings.Join(res, ",")
}

func PartTwo(computer Computer, program []int) int {
	// Note: you'll need to walkthrough your program to know how many bits each segment should
	// be.

	// we reverse engineer the segments of the A register 3-bits at a time,
	// the first 3 bits of the A register are responsible for producing
	// the last output — in my case, 0. Then we expand by an extra 3 bits, searching through
	// values 0 (0b000) to 7 (0b111), noting values that produce the last 2 outputs and so on...

	// from hacking around, 0b111 or 7 is known to produce the last output 0
	// at least, based on my given input
	knownSegments := []int{0b111}

	// we try to find the remaining segments of the output, we skip the last output 0, since
	// we already know that
	for i := 2; i <= len(program); i++ {
		var expandedSegments []int
		expectedOutput := program[len(program)-i:]

		for _, f := range knownSegments {
			for i := 0b000; i <= 0b111; i++ {
				regA := f<<3 + i
				comp := Computer{A: regA}
				comp.Run(program)

				if slices.Equal(comp.Out, expectedOutput) {
					expandedSegments = append(expandedSegments, regA)
				}
			}
		}

		knownSegments = expandedSegments
	}

	return slices.Min(knownSegments)
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
		computer.A = numerator / (1 << comboOp)
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
		computer.B = numerator / (1 << comboOp)
	case 7:
		numerator := computer.A
		comboOp := combo(operand, computer)
		computer.C = numerator / (1 << comboOp)
	}

	if !jumped {
		computer.IP += 2
	}
}
