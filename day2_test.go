package aoc2022

import (
	"fmt"
	"testing"
)

const (
	Draw      int = 0
	PlayerOne     = 1
	PlayerTwo     = 2
)

const (
	Rock     int = 0
	Paper        = 1
	Scissors     = 2
)

type Round struct {
	PlayerA int
	PlayerB int
}

func (r *Round) GetHandOfPlayer(player int) string {
	hand := ""

	move := 0
	if player == PlayerOne {
		move = r.PlayerA
	} else if player == PlayerTwo {
		move = r.PlayerB
	}

	switch move {

	case Rock:
		hand = "Rock"
		break
	case Paper:
		hand = "Paper"
		break
	case Scissors:
		hand = "Scissors"
		break
	}

	return hand
}

func (r *Round) GetScore() int {
	score := r.PlayerB + 1

	winner := getWinner(r.PlayerA, r.PlayerB)

	if winner == 0 { // draw
		score += 3
	} else if winner == 2 { // player 2 (me) winner
		score += 6
	}

	return score
}

func getWinner(playerA int, playerB int) int {

	m := 3

	outcome := (playerB - playerA) % m

	if outcome < 0 {
		outcome += m
	}

	if outcome == 0 {
		return Draw
	} else if outcome == 1 {
		return PlayerTwo
	} else if outcome == 2 {
		return PlayerOne
	}

	panic(fmt.Sprintf("incorrect outcome: %d", outcome))
}

func getPlayByInstruction(instruction string) int {
	if instruction == "A" || instruction == "X" {
		return Rock
	}

	if instruction == "B" || instruction == "Y" {
		return Paper
	}

	if instruction == "C" || instruction == "Z" {
		return Scissors
	}

	panic("getPlayByInstruction(): invalid instruction")
}

func getResultByInstruction(instruction string) int {
	if instruction == "X" {
		return PlayerOne
	}
	if instruction == "Y" {
		return Draw
	}
	if instruction == "Z" {
		return PlayerTwo
	}

	panic("getResultByInstruction(): invalid instruction ")
}

func getMoveLabel(move int) string {
	if move == Rock {
		return "Rock"
	} else if move == Paper {
		return "Paper"
	} else if move == Scissors {
		return "Scissors"
	}
	return "NOTHING"
}

func TestSingle(t *testing.T) {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			round := &Round{i, j}
			fmt.Printf("%s vs %s, winner: Player %d, score: %d\n", getMoveLabel(round.PlayerA), getMoveLabel(round.PlayerB), getWinner(round.PlayerA, round.PlayerB), round.GetScore())
		}
	}

}

func TestDayTwoPartOne(t *testing.T) {

	input := parseInput(2)

	score := 0
	for _, r := range input {
		round := &Round{getPlayByInstruction(string(r[0])), getPlayByInstruction(string(r[2]))}
		score += round.GetScore()
	}

	fmt.Printf("Total score: %d\n", score)

}

func TestDayTwoPartTwo(t *testing.T) {

	input := parseInput(2)

	score := 0
	for _, r := range input {
		desiredResult := getResultByInstruction(string(r[2]))
		playerAMove := getPlayByInstruction(string(r[0]))

		for i := 0; i <= 2; i++ {
			if getWinner(playerAMove, i) == desiredResult {
				round := &Round{playerAMove, i}
				score += round.GetScore()
				break
			}
		}
	}

	fmt.Printf("Total score: %d\n", score)

}
