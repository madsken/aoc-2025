package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}

	fileString := string(fileBytes)
	currentNumber := 50
	count := 0
	for line := range strings.SplitSeq(fileString, "\n") {
		if line == "" {
			continue
		}
		direction, number, rotations, err := ProcessLineWithRounds(line)
		if err != nil {
			log.Fatal(err)
		}

		count += rotations

		switch direction {
		case "L":
			zeroCheck := currentNumber
			currentNumber = currentNumber - number
			if zeroCheck != 0 && currentNumber < 0 {
				count += 1
			}
		case "R":
			currentNumber = currentNumber + number
			if currentNumber > 100 {
				count += 1
			}
		default:
			log.Fatal("Unsupported direction: " + direction)
		}

		if currentNumber < 0 {
			currentNumber = 100 + currentNumber
		} else {
			currentNumber = currentNumber % 100
		}

		if currentNumber == 0 {
			count += 1
		}
	}

	fmt.Printf("Day 2 answer: %d\n", count)
}
