package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partitionfIndex := slices.Index(input, "")
	if partitionfIndex < 0 {
		log.Fatal("invalid input")
	}

	rules := input[:partitionfIndex]
	pageUpdates := input[partitionfIndex+1:]
	rulesLookup := constructRulesLookup(rules)

	fmt.Println("solution to part one: ", PartOne(rulesLookup, pageUpdates))
	fmt.Println("solution to part two: ", PartTwo(rulesLookup, pageUpdates))
}

func PartOne(rulesLookup map[string]struct{}, updates []string) int {
	sumOfMids := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")
		if isValidPageOrder(pages, rulesLookup) {
			mid := pages[len(pages)/2]
			midAsInt, err := strconv.Atoi(mid)
			if err != nil {
				log.Fatal(err)
			}

			sumOfMids += midAsInt
		}
	}

	return sumOfMids
}

func PartTwo(rulesLookup map[string]struct{}, updates []string) int {
	sumOfMids := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")
		if !isValidPageOrder(pages, rulesLookup) {
			sortPages(pages, rulesLookup)
			mid := pages[len(pages)/2]
			midAsInt, err := strconv.Atoi(mid)
			if err != nil {
				log.Fatal(err)
			}

			sumOfMids += midAsInt
		}
	}

	return sumOfMids
}

// converts an update to a list of rules that must exist for the update to be valid
// 15,16,18 -> 15|16, 15|18, 16|18
func expandToRules(pages []string) []string {
	var requiredRules []string
	for i, page := range pages[:len(pages)-1] {
		for _, subPage := range pages[i+1:] {
			rule := fmt.Sprintf("%s|%s", page, subPage)
			requiredRules = append(requiredRules, rule)
		}
	}

	return requiredRules
}

func constructRulesLookup(rules []string) map[string]struct{} {
	lookup := make(map[string]struct{})
	for _, rule := range rules {
		lookup[rule] = struct{}{}
	}
	return lookup
}

func isValidPageOrder(pages []string, rulesLookup map[string]struct{}) bool {
	requiredRules := expandToRules(pages)

	for _, rule := range requiredRules {
		_, exists := rulesLookup[rule]
		if !exists {
			return false
		}
	}

	return true
}

// sorts pages in place using the order provided in the rules lookup
func sortPages(pages []string, rulesLookup map[string]struct{}) {
	// sort pages using rules order
	slices.SortFunc(pages, func(a, b string) int {
		// returns -1 if a should appear before b and 1 if otherwise

		rule := a + "|" + b
		_, exists := rulesLookup[rule]
		if exists {
			return -1
		}

		return 1
	})
}
