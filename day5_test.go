package aoc2022

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

const stackAmount int = 9

type platform struct {
	Stacks [stackAmount][]string
	rInst  *regexp.Regexp
}

func (p *platform) init() {
	p.rInst, _ = regexp.Compile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	for i := 0; i < stackAmount; i++ {
		p.Stacks[i] = []string{}
	}
}

func (p *platform) processStackInput(input string) {
	for i := 0; i < stackAmount; i++ {
		crate := string(input[(i*4)+1])
		if crate != " " {
			p.Stacks[i] = append(p.Stacks[i], crate)
		}
	}
}

func (p *platform) processMoveInstruction(instruction string, stacked bool) {
	m := p.rInst.FindStringSubmatch(instruction)
	if len(m) != 4 {
		panic(fmt.Sprintf("invalid move instruction: %s", instruction))
	}

	amount, _ := strconv.Atoi(m[1])
	from, _ := strconv.Atoi(m[2])
	to, _ := strconv.Atoi(m[3])

	if stacked {
		p.moveStack(from-1, to-1, amount)
	} else {
		for i := 0; i < amount; i++ {
			p.move(from-1, to-1)
		}
	}

}

func (p *platform) move(from, to int) {
	el := p.Stacks[from][0]                              // take
	p.Stacks[from] = p.Stacks[from][1:]                  // remove
	p.Stacks[to] = append([]string{el}, p.Stacks[to]...) // prepend
}

func (p *platform) moveStack(from, to, amount int) {
	el := p.Stacks[from][0:amount] // take

	// we need to copy the partial slice before we append it to it's dest.
	// if we don't, we'll keep working with a ref to a part of the first stack and all goes to hell
	c := make([]string, amount)
	copy(c, el)

	p.Stacks[from] = p.Stacks[from][amount:]  // remove
	p.Stacks[to] = append(c, p.Stacks[to]...) // prepend
}

func (p *platform) print() {
	fmt.Println("PLATFORM:")
	for _, stack := range p.Stacks {
		fmt.Printf("%v\n", stack)
	}
}

func TestDayFivePartOne(t *testing.T) {

	input := parseInput(5)

	tPlatform := &platform{}
	tPlatform.init()

	// line 1 to 8 lines are stack input
	for i := 0; i < 8; i++ {
		tPlatform.processStackInput(input[i])
	}

	// line 11 to end are move instructions
	for i := 10; i < len(input); i++ {
		//fmt.Printf("%s\n", input[i])
		tPlatform.processMoveInstruction(input[i], false)
	}

	fmt.Print("ANSWER: ")
	for _, v := range tPlatform.Stacks {
		fmt.Printf("%s", v[0])
	}
	fmt.Print("\n")

}

func TestDayFivePartTwo(t *testing.T) {

	input := parseInput(5)

	tPlatform := &platform{}
	tPlatform.init()

	// line 1 to 8 lines are stack input
	for i := 0; i < 8; i++ {
		tPlatform.processStackInput(input[i])
	}

	// line 11 to end are move instructions
	for i := 10; i < len(input); i++ {
		//fmt.Printf("%s\n", input[i])
		tPlatform.processMoveInstruction(input[i], true)
		if i > 11 {
			//break
		}
	}

	tPlatform.print()

	fmt.Print("ANSWER: ")
	for _, v := range tPlatform.Stacks {
		if len(v) == 0 {
			fmt.Print("-")
			continue
		}

		fmt.Printf("%s", v[0])
	}
	fmt.Print("\n")

}
