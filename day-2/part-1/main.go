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
			myHand := hands[1]
			totalScore += getScoreForRound(opponentsHand, myHand)
		}
	}

	fmt.Println(totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getScoreForRound(opponentsHand, myHand string) int {
	won := 6
	lost := 0
	draw := 3

	rock := 1
	paper := 2
	scissors := 3

	switch opponentsHand {
	case "A": // Rock
		switch myHand {
		case "Y": // Paper - 1
			return paper + won
		case "X": // Rock - 2
			return rock + draw
		case "Z": // Scissors - 3
			return scissors + lost
		}

	case "B": // Paper
		switch myHand {
		case "Y": // Paper - 1
			return paper + draw
		case "X": // Rock - 2
			return rock + lost
		case "Z": // Scissors - 3
			return scissors + won
		}
	case "C": // Scissors
		switch myHand {
		case "Y": // Paper - 1
			return paper + lost
		case "X": // Rock - 2
			return rock + won
		case "Z": // Scissors - 3
			return scissors + draw
		}
	}
	return 0
}
