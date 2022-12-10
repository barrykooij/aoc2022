package aoc2022

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func contains(a string, b string) bool {
	aParts := strings.Split(a, "-")
	bParts := strings.Split(b, "-")

	a0, _ := strconv.Atoi(aParts[0])
	a1, _ := strconv.Atoi(aParts[1])
	b0, _ := strconv.Atoi(bParts[0])
	b1, _ := strconv.Atoi(bParts[1])

	if (a0 >= b0 && a1 <= b1) || (b0 >= a0 && b1 <= a1) {
		return true
	}

	return false
}

func overlaps(a string, b string) bool {
	aParts := strings.Split(a, "-")
	bParts := strings.Split(b, "-")

	a0, _ := strconv.Atoi(aParts[0])
	a1, _ := strconv.Atoi(aParts[1])
	b0, _ := strconv.Atoi(bParts[0])
	b1, _ := strconv.Atoi(bParts[1])

	if (a1 >= b0 && a0 <= b1) || (a1 >= b1 && a0 <= b0) || (a1 >= b1 && a0 <= b1) || (a0 <= b0 && a1 <= b1 && a1 >= b0) {
		return true
	}

	return false
}

func TestDayFourPartOne(t *testing.T) {
	input := parseInput(4)

	c := 0
	for _, row := range input {
		parts := strings.Split(row, ",")
		if contains(parts[0], parts[1]) {
			c++
		}
	}

	fmt.Printf("Contains counter: %d\n", c)
}

func TestDayFourPartTwo(t *testing.T) {
	input := parseInput(4)

	c := 0
	for _, row := range input {
		parts := strings.Split(row, ",")
		if overlaps(parts[0], parts[1]) {
			c++
		}
	}

	fmt.Printf("Contains counter: %d\n", c)
}
