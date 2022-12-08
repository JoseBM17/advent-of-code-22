package main

import (
	"fmt"
	advent_of_code_22 "github.com/JoseBM17/advent-of-code-22"
	"sort"
	"strconv"
)

type Elf struct {
	calories int
}

func parseElves(input []string) []Elf {
	cal := 0
	var elves []Elf
	for i, elem := range input {
		if elem != "" && i < (len(input)-1) {
			e, err := strconv.Atoi(elem)
			advent_of_code_22.Check(err)
			cal += e
		} else if i == (len(input) - 1) {
			e, err := strconv.Atoi(elem)
			advent_of_code_22.Check(err)
			elves = append(elves, Elf{calories: cal + e})
		} else {
			elves = append(elves, Elf{calories: cal})
			cal = 0
		}
	}

	return elves
}

func findMaxCalories(elves []Elf) int {
	max := elves[0]
	for _, elf := range elves {
		if elf.calories > max.calories {
			max = elf
		}
	}
	return max.calories
}

func find3MaxCalories(elves []Elf) int {
	var maxElves []Elf
	var totalCal int

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	maxElves = elves[:3]
	for _, e := range maxElves {
		totalCal += e.calories
	}
	return totalCal
}

func main() {
	input := advent_of_code_22.ReadLines("./day_01/input.txt")
	elves := parseElves(input)
	max := findMaxCalories(elves)
	// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
	fmt.Println(max)
	// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
	max = find3MaxCalories(elves)
	fmt.Println(max)
}
