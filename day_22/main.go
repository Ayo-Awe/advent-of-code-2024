package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func parseInput(input []string) ([]int, error) {
	var secrets []int
	for _, line := range input {
		secret, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, secret)
	}
	return secrets, nil
}

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	secrets, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(secrets))
	fmt.Println("solution to part two: ", PartTwo(secrets))
}

func PartOne(secrets []int) int {
	var sum int
	for _, secret := range secrets {
		sum += nthSecret(secret, 2000)[1999]
	}
	return sum
}

func PartTwo(buySecrets []int) int {
	sequencePrices := make(map[[4]int]int)

	for _, intialSecret := range buySecrets {
		secrets := nthSecret(intialSecret, 2000)
		secrets = append([]int{intialSecret}, secrets...)

		// map of a set of change seqs to their max obtainable value in for this buyer
		seen := make(map[[4]int]struct{})
		for i := 1; i < len(secrets); i++ {
			var sequence [4]int
			start, end := i, i+3

			if end >= len(secrets) {
				continue
			}

			for j := start; j <= end; j++ {
				change := (secrets[j] % 10) - (secrets[j-1] % 10)
				sequence[j-i] = change
			}

			if _, ok := seen[sequence]; ok {
				continue
			}

			price := secrets[end] % 10
			sequencePrices[sequence] += price
			seen[sequence] = struct{}{}
		}
	}

	var maxBananas int
	for _, price := range sequencePrices {
		maxBananas = max(maxBananas, price)
	}

	return maxBananas
}

func nthSecret(secret, n int) []int {
	var secrets []int
	curr := secret

	for range n {
		curr = curr ^ (curr * 64)

		curr = curr % (1 << 24)

		curr = curr ^ (curr / 32)

		curr = curr % (1 << 24)

		curr = curr ^ (curr * 2048)

		curr = curr % (1 << 24)
		secrets = append(secrets, curr)
	}
	return secrets
}
