package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			hands := strings.Split(line, " ")
			if err != nil {
				fmt.Println("Error during parsing hands of the game")
				continue
			}

			opponentsHand := hands[0]
			roundResult := hands[1]
			totalScore += getScoreForRound(opponentsHand, roundResult)
		}
	}

	fmt.Println(totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getScoreForRound(opponentsHand, roundResult string) int {
	won := 6
	lost := 0
	draw := 3

	rock := 1
	paper := 2
	scissors := 3

	switch opponentsHand {
	case "A": // Rock
		switch roundResult {
		case "Y": // draw
			return draw + rock
		case "X": // lose
			return lost + scissors
		case "Z": // win
			return won + paper
		}
	case "B": // Paper
		switch roundResult {
		case "Y": // draw
			return draw + paper
		case "X": // lose
			return lost + rock
		case "Z": // win
			return won + scissors
		}
	case "C": // Scissors
		switch roundResult {
		case "Y": // draw
			return draw + scissors
		case "X": // lose
			return lost + paper
		case "Z": // win
			return won + rock
		}
	}
	return 0
}
