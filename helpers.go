package aoc2022

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput(day int) []string {

	var lines = make([]string, 0)

	f, err := os.OpenFile(fmt.Sprintf("day%d_input", day), os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return lines
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return lines
	}
	return lines
}
