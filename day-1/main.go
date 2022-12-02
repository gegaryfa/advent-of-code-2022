package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var elfCalories []int
	currentElfCal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			currentFoodCal, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during parsing calories")
				continue
			}
			currentElfCal += currentFoodCal
			continue
		}
		elfCalories = append(elfCalories, currentElfCal)
		currentElfCal = 0
	}

	sort.Ints(elfCalories)

	fmt.Println(elfCalories[len(elfCalories)-1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
