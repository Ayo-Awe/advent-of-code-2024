package main

import (
	"fmt"
	"log"
	"maps"
	"sort"
	"strings"

	"github.com/ayo-awe/advent-of-code-2024/aoc"
)

func main() {
	input, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	network := buildNetwork(input)

	fmt.Println("solution to part one: ", PartOne(network))
	fmt.Println("solution to part two: ", PartTwo(network))
}

func PartOne(network map[string]map[string]struct{}) int {
	matches := make(map[string]struct{})
	for node, subnet := range network {
		if !strings.HasPrefix(node, "t") {
			continue
		}

		var subnodes []string
		for subnode := range subnet {
			subnodes = append(subnodes, subnode)
		}

		for i := 0; i < len(subnodes)-1; i++ {
			for j := i + 1; j < len(subnodes); j++ {
				subA, subB := subnodes[i], subnodes[j]
				lookupKey := []string{node, subA, subB}
				sort.Strings(lookupKey)

				key := strings.Join(lookupKey, ",")

				if _, exists := matches[key]; exists {
					continue
				}

				if _, connected := network[subA][subB]; connected {
					matches[key] = struct{}{}
				}
			}
		}
	}

	return len(matches)
}

func PartTwo(network map[string]map[string]struct{}) string {
	nodes := make(map[string]struct{})
	for node := range network {
		nodes[node] = struct{}{}
	}

	clique := maxClique(map[string]struct{}{}, nodes, map[string]struct{}{}, network)
	sort.Strings(clique)

	return strings.Join(clique, ",")
}

func buildNetwork(connections []string) map[string]map[string]struct{} {
	network := make(map[string]map[string]struct{})
	for _, connection := range connections {
		nodes := strings.Split(connection, "-")
		node1, node2 := nodes[0], nodes[1]

		if _, ok := network[node1]; !ok {
			network[node1] = make(map[string]struct{})
		}

		if _, ok := network[node2]; !ok {
			network[node2] = make(map[string]struct{})
		}

		network[node1][node2] = struct{}{}
		network[node2][node1] = struct{}{}
	}
	return network
}

func intersection(a, b map[string]struct{}) map[string]struct{} {
	inter := make(map[string]struct{})

	for k := range a {
		if _, ok := b[k]; ok {
			inter[k] = struct{}{}
		}
	}

	return inter
}

// bron-kerbosch algorithm implementation
func maxClique(r, p, x map[string]struct{}, network map[string]map[string]struct{}) []string {
	if len(p) == 0 && len(x) == 0 {
		var clique []string
		for k := range r {
			clique = append(clique, k)
		}
		return clique
	}

	var largestClique []string
	for node := range p {
		rr := maps.Clone(r)
		rr[node] = struct{}{}
		pp := intersection(p, network[node])
		xx := intersection(x, network[node])

		clique := maxClique(rr, pp, xx, network)
		delete(p, node)
		x[node] = struct{}{}

		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}

	return largestClique
}
