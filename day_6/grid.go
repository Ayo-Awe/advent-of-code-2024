package main

import "slices"

type GridArea [][]string

func (g GridArea) FindGuard() *Guard {
	for y, row := range g {
		for x, val := range row {
			if isGuard(val) {
				return &Guard{
					x:         x,
					y:         y,
					direction: Direction(val),
				}
			}
		}
	}

	return nil
}

func (g GridArea) IsOutOfBounds(x, y int) bool {
	return x < 0 || x > len(g[0])-1 || y < 0 || y > len(g)-1
}

func (g GridArea) Clone() GridArea {
	clone := GridArea{}
	for _, row := range g {
		clone = append(clone, slices.Clone(row))
	}
	return clone
}
