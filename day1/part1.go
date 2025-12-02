package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1(file string) {
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
		direction, number, err := ProcessLine(line)
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "L":
			currentNumber = currentNumber - number
		case "R":
			currentNumber = currentNumber + number
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

	fmt.Printf("Day 1 answer: %d\n", count)
}
