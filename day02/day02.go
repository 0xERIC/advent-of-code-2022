package main

import (
	"bufio"
	"fmt"
	"os"
)

var part1Scores = map[string]int{
	"A X": 4, // tie throwing rock
	"A Y": 8, // win throwing paper
	"A Z": 3, // loss throwing scissors
	"B X": 1, // loss throwing rock
	"B Y": 5, // tie throwing paper
	"B Z": 9, // win throwing scissors
	"C X": 7, // win throwing rock
	"C Y": 2, // loss throwing paper
	"C Z": 6, // tie throwing scissors
}

var part2Scores = map[string]int{
	"A X": 3, // loss throwing scissors
	"A Y": 4, // tie throwing rock
	"A Z": 8, // win throwing paper
	"B X": 1, // loss throwing rock
	"B Y": 5, // tie throwing paper
	"B Z": 9, // win throwing scissors
	"C X": 2, // loss throwing paper
	"C Y": 6, // tie throwing scissors
	"C Z": 7, // win throwing rock
}

func main() {
	totalScore := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nextLine := scanner.Text()
		totalScore += part2Scores[nextLine]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Print(totalScore)
}
