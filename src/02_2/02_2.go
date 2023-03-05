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

		theirs := theirsToRPS(theirsStr)
		mine := mineToResult(mineStr)
		totalScore += mine.score() + theirs.inverseOutcome(mine).score()
	}
	fmt.Printf("%d\n", totalScore)
}

type rps int

const (
	rock rps = iota
	paper
	scissors
)

var theirInputs = map[string]rps{"A": rock, "B": paper, "C": scissors}

func theirsToRPS(s string) rps {
	return theirInputs[s]
}

var myInputs = map[string]result{"X": lose, "Y": tie, "Z": win}

func mineToResult(s string) result {
	return myInputs[s]
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

var inverseOutcomes = [][]rps{
	rock:     {win: paper, tie: rock, lose: scissors},
	paper:    {win: scissors, tie: paper, lose: rock},
	scissors: {win: rock, tie: scissors, lose: paper},
}

func (r rps) inverseOutcome(s result) rps { return inverseOutcomes[r][s] }
