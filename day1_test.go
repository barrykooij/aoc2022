package aoc2022

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

type Elf struct {
	id        int
	totalCal  int
	foodItems []int
}

func (e *Elf) AddItem(i int) {
	e.foodItems = append(e.foodItems, i)
	e.totalCal += i
}

func parseElfsFromFile() []*Elf {
	var elfs = make([]*Elf, 0)

	lines := parseInput(1)

	curId := 1
	curElf := &Elf{id: curId}
	for _, line := range lines {
		if len(line) == 0 {
			// new line, next elf
			elfs = append(elfs, curElf)
			curId++
			curElf = &Elf{id: curId}
			continue
		}

		cal, e := strconv.Atoi(line)
		if e != nil {
			panic(e.Error())
		}

		curElf.AddItem(cal)
	}
	elfs = append(elfs, curElf)

	fmt.Printf("scanned %d lines \n", len(lines))

	return elfs
}

func DayOne() {
	elfs := parseElfsFromFile()

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].totalCal > elfs[j].totalCal
	})

	fmt.Printf("Winning elf %d with %d calories and %d items \n", elfs[0].id, elfs[0].totalCal, len(elfs[0].foodItems))
	fmt.Printf("Top 3 cals: %d \n", (elfs[0].totalCal + elfs[1].totalCal + elfs[2].totalCal))
}

func TestDayOne(t *testing.T) {
	DayOne()
}
