package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./input_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	left := []int{}
	right := []int{}
	for _, line := range strings.Split(string(bytes), "\n") {
		leftAndRight := regexp.MustCompile("\\d+").FindAllString(line, -1)

		if len(leftAndRight) != 2 {
			continue
		}

		leftInt, err := strconv.Atoi(leftAndRight[0])
		if err != nil {
			log.Fatal(err)
		}

		rightInt, err := strconv.Atoi(leftAndRight[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range len(left) {
		sum += diff(left[i], right[i])
	}

	fmt.Println(sum)
}

func diff(left, right int) int {
	sub := left - right

	if sub < 0 {
		return -sub
	}

	return sub
}
