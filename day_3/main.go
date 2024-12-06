package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	sum        int  // current operation sum
	mulEnabled bool // dictates if mul instructions are enabled
}

var (
	instructionHandlers = map[string]func(string, *State) *State{
		"mul":   mulHandler,
		"do":    doHandler,
		"don't": dontHandler,
	}
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one is: ", partOne(string(input)))
	fmt.Println("solution to part two is: ", partTwo(string(input)))
}
func partOne(input string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	instructions := r.FindAllString(string(input), -1)

	sum := 0
	for _, instruction := range instructions {
		r := regexp.MustCompile(`\d+`)
		operands := r.FindAllString(instruction, -1)

		leftOperand, err := strconv.Atoi(operands[0])
		if err != nil {
			log.Fatal(err)
		}

		rightOperand, err := strconv.Atoi(operands[1])
		if err != nil {
			log.Fatal(err)
		}

		sum += leftOperand * rightOperand
	}

	return sum
}

func partTwo(input string) int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	instructions := r.FindAllString(input, -1)

	state := &State{sum: 0, mulEnabled: true}
	for _, instruction := range instructions {
		operator, args := splitInstruction(instruction)
		handler, exists := instructionHandlers[operator]
		if !exists {
			log.Fatal("instruction not supported")
		}

		handler(args, state)
	}

	return state.sum
}

func splitInstruction(instruction string) (string, string) {
	index := strings.Index(instruction, "(")
	operator := instruction[:index]
	args := instruction[index+1:]
	return operator, args
}

// args: "(...)"
func mulHandler(args string, state *State) *State {
	if !state.mulEnabled {
		return state
	}

	r := regexp.MustCompile(`\d+`)
	operands := r.FindAllString(args, -1)

	leftOperand, err := strconv.Atoi(operands[0])
	if err != nil {
		log.Fatal(err)
	}

	rightOperand, err := strconv.Atoi(operands[1])
	if err != nil {
		log.Fatal(err)
	}

	state.sum += leftOperand * rightOperand

	return state
}

func dontHandler(_ string, state *State) *State {
	state.mulEnabled = false
	return state
}

func doHandler(_ string, state *State) *State {
	state.mulEnabled = true
	return state
}
