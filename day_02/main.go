package main

import (
	"fmt"
	advent_of_code_22 "github.com/JoseBM17/advent-of-code-22"
	"strconv"
	"strings"
)

type Round struct {
	opponentChoice string
	myChoice       string
}

type Rounds []Round

func parseInput(input []string) Rounds {
	//lines := strings.Split(strings.TrimSpace(raw), "\n")
	result := Rounds{}
	for _, round := range input {
		line := strings.Split(round, " ")
		r := Round{
			opponentChoice: line[0],
			myChoice:       line[1],
		}
		result = append(result, r)
	}

	return result
}

func getRoundScore(r Round) int {
	switch r.opponentChoice {
	case "A": // rock
		switch r.myChoice {
		case "X": // rock
			return 1 + 3
		case "Y":
			return 2 + 6
		case "Z":
			return 3 + 0
		}
	case "B": // paper
		switch r.myChoice {
		case "X": // rock
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C": // scissors
		switch r.myChoice {
		case "X": // rock
			return 1 + 6
		case "Y":
			return 2 + 0
		case "Z":
			return 3 + 3
		}
	}
	panic("Unreachable code")
}

func getRoundScorePart2(r Round) int {
	switch r.opponentChoice {
	case "A": // rock
		switch r.myChoice {
		case "X": // lose
			return 3 + 0 //scissors
		case "Y":
			return 1 + 3
		case "Z":
			return 2 + 6
		}
	case "B": // paper
		switch r.myChoice {
		case "X": // lose
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C": // scissors
		switch r.myChoice {
		case "X": // lose
			return 2 + 0
		case "Y":
			return 3 + 3
		case "Z":
			return 1 + 6
		}
	}
	panic("Unreachable code")
}

func main() {
	input := advent_of_code_22.ReadLines("./day_02/input.txt")
	rounds := parseInput(input)

	totalScore := 0
	for _, round := range rounds {
		totalScore += getRoundScore(round)
	}
	fmt.Println("My total score is: " + strconv.Itoa(totalScore))

	totalScorePart2 := 0
	for _, round := range rounds {
		totalScorePart2 += getRoundScorePart2(round)
	}
	fmt.Println("Part 2 total score is: " + strconv.Itoa(totalScorePart2))
}
