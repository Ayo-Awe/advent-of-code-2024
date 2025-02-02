package main

import (
	"fmt"
	"log"
	"maps"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

var (
	rInput = regexp.MustCompile(`\w{3}: [0-1]{1}`)
	rGates = regexp.MustCompile(`\w{3} (XOR|OR|AND) \w{3} -> \w{3}`)

	I1, I2, GateType, O = 0, 2, 1, 3
)

func parseInput(input string) (map[string]int, [][4]string, error) {
	inputs := make(map[string]int)
	gates := [][4]string{}

	for _, rawInput := range rInput.FindAllString(input, -1) {
		parts := strings.Split(rawInput, ": ")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("unexpected input syntax")
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		inputs[parts[0]] = value
	}

	for _, gate := range rGates.FindAllString(input, -1) {
		parts := strings.Split(gate, " ")
		if len(parts) != 5 {
			return nil, nil, fmt.Errorf("unexpected gate syntax: %s", gate)
		}

		input1, gateType, input2, output := parts[0], parts[1], parts[2], parts[4]
		gates = append(gates, [4]string{input1, gateType, input2, output})
	}

	return inputs, gates, nil
}

func main() {
	input, err := aoc.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs, gates, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(inputs, gates))
	fmt.Println("solution to part two: ", PartTwo())
}

func PartOne(inputs map[string]int, gates [][4]string) int {
	inputs = maps.Clone(inputs)

	pendingGates := gates
	for len(pendingGates) > 0 {
		currPending := pendingGates
		pendingGates = nil

		for _, gate := range currPending {
			_, I1Ready := inputs[gate[I1]]
			_, I2Ready := inputs[gate[I2]]

			if !I1Ready || !I2Ready {
				pendingGates = append(pendingGates, gate)
				continue
			}

			var output int
			switch gate[GateType] {
			case "XOR":
				output = inputs[gate[I1]] ^ inputs[gate[I2]]
			case "AND":
				output = inputs[gate[I1]] & inputs[gate[I2]]
			case "OR":
				output = inputs[gate[I1]] | inputs[gate[I2]]
			default:
				panic("invalid gate type")
			}

			inputs[gate[O]] = output
		}
	}

	var outputs []string
	for wire := range inputs {
		if strings.HasPrefix(wire, "z") {
			outputs = append(outputs, wire)
		}
	}

	sort.Strings(outputs)
	slices.Reverse(outputs)

	var decimalOutput int
	for i, output := range outputs {
		val := inputs[output]
		val = val << (len(outputs) - 1 - i)
		decimalOutput += val
	}

	fmt.Println(outputs)
	fmt.Printf("%b\n", decimalOutput)

	return decimalOutput
}

func PartTwo() int {
	return 0
}

