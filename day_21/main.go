package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
	"golang.org/x/exp/maps"
)

var numericalKeypad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'#', '0', 'A'},
}

var directionalKeypad = [][]rune{
	{'#', '^', 'A'},
	{'<', 'v', '>'},
}

var directions = map[rune][2]int{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

const (
	X = iota
	Y
)

var precomputedSeqs = precomputeSeq()

// TODO: cleanup

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(input))
	fmt.Println("solution to part two: ", PartTwo(input))
}

func PartOne(codes []string) int {
	memo := map[string]int{}

	var complexitySum int
	for _, code := range codes {
		inputs := keypresses(code, numericalKeypad, [2]int{2, 3})

		optimal := math.MaxInt
		for _, input := range inputs {
			input = "A" + input

			var length int
			for i := 0; i < len(input)-1; i++ {
				x, y := rune(input[i]), rune(input[i+1])
				length += computeLength(x, y, 2, memo)
			}

			if length < optimal {
				optimal = length
			}
		}

		intCode, err := strconv.Atoi(code[:3])
		if err != nil {
			log.Fatal(err)
		}

		complexitySum += intCode * optimal
	}

	return complexitySum
}

func getMemoKey(x, y rune, depth int) string {
	return fmt.Sprintf("%c:%c:%d", x, y, depth)
}

func computeLength(x, y rune, depth int, memo map[string]int) int {
	if depth == 1 {
		return len(precomputedSeqs[[2]rune{x, y}][0])
	}

	if length, seen := memo[getMemoKey(x, y, depth)]; seen {
		return length
	}

	optimal := math.MaxInt
	for _, seq := range precomputedSeqs[[2]rune{x, y}] {
		seq = append([]rune{'A'}, seq...)

		var length int
		for i := 0; i < len(seq)-1; i++ {
			a, b := seq[i], seq[i+1]
			val := computeLength(a, b, depth-1, memo)
			length += val
		}
		if length < optimal {
			optimal = length
		}
	}

	memo[getMemoKey(x, y, depth)] = optimal
	return optimal
}

func precomputeSeq() map[[2]rune][][]rune {
	cache := make(map[[2]rune][][]rune)

	kw, kh := len(directionalKeypad[0]), len(directionalKeypad)
	for i := range kw * kh {
		srcX, srcY := i%kw, i/kw
		for j := range kw * kh {
			destX, destY := j%kw, j/kw

			if i == 0 || j == 0 {
				continue
			}

			// find the shortest paths between the two keypads
			queue := []state{{nil, [2]int{srcX, srcY}}}
			seen := map[[2]int]int{}
			minPathLength := math.MaxInt
			var minPaths [][]rune

			for len(queue) > 0 {

				// pop element of the queue
				curr := queue[0]
				queue = queue[1:]

				// this path costs more that our existing solution
				if len(curr.sequence) > minPathLength {
					continue
				}

				// out of bounds
				if curr.pos[Y] < 0 || curr.pos[Y] >= kh || curr.pos[X] < 0 || curr.pos[X] >= kw {
					continue
				}

				// ignore invalid square
				if directionalKeypad[curr.pos[Y]][curr.pos[X]] == '#' {
					continue
				}

				// skip — we've seen this square before but with a shorter sequence
				if cost, ok := seen[curr.pos]; ok && len(curr.sequence) > cost {
					continue
				}

				if directionalKeypad[curr.pos[Y]][curr.pos[X]] == directionalKeypad[destY][destX] {
					// reset the minimum sequences seen
					if len(curr.sequence) < minPathLength {
						minPaths = nil
						minPathLength = len(curr.sequence)
					}

					curr.sequence = append(curr.sequence, 'A')

					minPaths = append(minPaths, curr.sequence)
					continue
				}

				for seq, dir := range directions {
					sequence := slices.Clone(curr.sequence)
					sequence = append(sequence, seq)
					queue = append(queue, state{sequence, [2]int{curr.pos[X] + dir[X], curr.pos[Y] + dir[Y]}})
				}

				seen[curr.pos] = len(curr.sequence)
			}

			srcRune, targetRune := directionalKeypad[srcY][srcX], directionalKeypad[destY][destX]
			cache[[2]rune{srcRune, targetRune}] = minPaths
		}
	}

	return cache
}

func PartTwo(codes []string) int {
	memo := map[string]int{}

	var complexitySum int
	for _, code := range codes {
		inputs := keypresses(code, numericalKeypad, [2]int{2, 3})

		optimal := math.MaxInt
		for _, input := range inputs {
			input = "A" + input

			var length int
			for i := 0; i < len(input)-1; i++ {
				x, y := rune(input[i]), rune(input[i+1])
				length += computeLength(x, y, 25, memo)
			}

			if length < optimal {
				optimal = length
			}
		}

		intCode, err := strconv.Atoi(code[:3])
		if err != nil {
			log.Fatal(err)
		}

		complexitySum += intCode * optimal
	}

	return complexitySum
}

type state struct {
	sequence []rune
	pos      [2]int
}

func keypresses(code string, keypad [][]rune, start [2]int) []string {
	queue := []state{{nil, start}}
	seen := make(map[[2]int]int)

	var targetIndex int
	var nextQueueItems []state
	minSequenceLength := math.MaxInt

	for targetIndex < len(code) {
		// all paths have been explored for the current target
		// reset search parameters
		if len(queue) == 0 {
			maps.Clear(seen)
			targetIndex++
			minSequenceLength = math.MaxInt
			queue = nextQueueItems
			continue
		}

		// pop element of the queue
		curr := queue[0]
		queue = queue[1:]

		// this path costs more that our existing solution
		if len(curr.sequence) > minSequenceLength {
			continue
		}

		// out of bounds
		if curr.pos[Y] < 0 || curr.pos[Y] >= len(keypad) || curr.pos[X] < 0 || curr.pos[X] >= len(keypad[0]) {
			continue
		}

		// ignore invalid square
		if keypad[curr.pos[Y]][curr.pos[X]] == '#' {
			continue
		}

		// skip — we've seen this square before but with a shorter sequence
		if cost, ok := seen[curr.pos]; ok && len(curr.sequence) > cost {
			continue
		}

		if keypad[curr.pos[Y]][curr.pos[X]] == rune(code[targetIndex]) {
			// reset the minimum sequences seen
			if len(curr.sequence) < minSequenceLength {
				nextQueueItems = nil
			}

			minSequenceLength = len(curr.sequence)

			// button press
			curr.sequence = append(curr.sequence, 'A')
			nextQueueItems = append(nextQueueItems, curr)
			continue
		}

		for seq, dir := range directions {
			sequence := slices.Clone(curr.sequence)
			sequence = append(sequence, seq)
			queue = append(queue, state{sequence, [2]int{curr.pos[X] + dir[X], curr.pos[Y] + dir[Y]}})
		}

		seen[curr.pos] = len(curr.sequence)
	}

	var shortestKeypresses []string
	for _, item := range nextQueueItems {
		shortestKeypresses = append(shortestKeypresses, string(item.sequence))
	}
	return shortestKeypresses
}
