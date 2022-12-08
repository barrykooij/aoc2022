package aoc2022

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode"
)

type group struct {
	rucksacks []*rucksack
}

func (g *group) FindBadge() string {
	m1 := getMapOfItemSlice(g.rucksacks[0].AllItems)
	m2 := getMapOfItemSlice(g.rucksacks[1].AllItems)
	for _, item := range g.rucksacks[2].AllItems {
		_, f1 := m1[item]
		if !f1 {
			continue
		}
		_, f2 := m2[item]
		if !f2 {
			continue
		}
		return item
	}
	panic("no badge found")
}

type compartment struct {
	Items       []string
	MappedItems map[string]bool
}

type rucksack struct {
	AllItems                   []string
	CompartmentA, CompartmentB *compartment
}

func (r *rucksack) FindMutualCompartmentItem() string {

	for _, item := range r.CompartmentB.Items {
		if _, found := r.CompartmentA.MappedItems[item]; found {
			return item
		}
	}

	return ""
}

func makeRucksack(data string) *rucksack {
	h := int(math.Round(float64(len(data)) * .5))

	return &rucksack{
		AllItems:     removeDuplicates(strings.Split(data, "")),
		CompartmentA: makeCompartment(strings.Split(data[0:h], "")),
		CompartmentB: makeCompartment(strings.Split(data[h:len(data)], "")),
	}
}

func makeCompartment(items []string) *compartment {

	return &compartment{removeDuplicates(items), getMapOfItemSlice(items)}
}

func getMapOfItemSlice(items []string) map[string]bool {
	itemsMap := make(map[string]bool)
	for _, v := range items {
		itemsMap[v] = true
	}
	return itemsMap
}

func removeDuplicates(s []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, v := range s {
		if _, r := keys[v]; !r {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}

func getPriorityOfLetter(letter string) int {
	if len(letter) != 1 {
		panic(fmt.Sprintf("getPriorityOfLetter() expects a single letter, given: %s", letter))
	}

	r := []rune(letter)[0]

	if unicode.IsUpper(r) {
		r -= 38
	} else if unicode.IsLower(r) {
		r -= 96
	}

	return int(r)
}

func TestDayThreePartOne(t *testing.T) {

	input := parseInput(3)

	totalPrio := 0

	for _, d := range input {
		sack := makeRucksack(d)
		totalPrio += getPriorityOfLetter(sack.FindMutualCompartmentItem())
	}

	fmt.Printf("total prio: %d\n", totalPrio)

}

func TestDayThreePartTwo(t *testing.T) {

	input := parseInput(3)

	groups := make([]*group, 0)

	grp := &group{}

	for _, d := range input {
		grp.rucksacks = append(grp.rucksacks, makeRucksack(d))
		if len(grp.rucksacks) > 2 {
			groups = append(groups, grp)
			grp = &group{}
		}
	}

	totalPrio := 0
	for _, grp = range groups {
		totalPrio += getPriorityOfLetter(grp.FindBadge())
	}

	fmt.Printf("total group/badge prio: %d\n", totalPrio)

}
