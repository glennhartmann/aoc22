package main

import (
	"fmt"
	"io"
)

func main() {
	var totalScore int64
	for {
		var theirsStr, mineStr string
		n, err := fmt.Scanln(&theirsStr, &mineStr)
		if err == io.EOF {
			break
		}
		if n != 2 || err != nil {
			panic("invalid input")
		}

		theirs := inputToRPS(theirsStr)
		mine := inputToRPS(mineStr)
		totalScore += mine.score() + mine.outcome(theirs).score()
	}
	fmt.Printf("%d\n", totalScore)
}

type rps int

const (
	rock rps = iota
	paper
	scissors
)

func inputToRPS(s string) rps {
	switch s {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	}
	panic(fmt.Sprintf("unsupported input: %s", s))
}

var scores = []int64{rock: 1, paper: 2, scissors: 3}

func (r rps) score() int64 { return scores[r] }

type result int

const (
	win result = iota
	tie
	lose
)

var outcomeScores = []int64{win: 6, tie: 3, lose: 0}

func (r result) score() int64 { return outcomeScores[r] }

var outcomes = [][]result{
	rock:     {rock: tie, paper: lose, scissors: win},
	paper:    {rock: win, paper: tie, scissors: lose},
	scissors: {rock: lose, paper: win, scissors: tie},
}

func (r rps) outcome(s rps) result { return outcomes[r][s] }
