package main

import (
	//"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	//"io"
	"os"
	//"path/filepath"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
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

	fmt.Println(count)
}

func ProcessLine(line string) (string, int, error) {
	direction := string(line[0])

	numberString := string(line[1:])
	if len(numberString) > 2 {
		numberString = string(line[2:])
	}

	number, err := strconv.Atoi(numberString)
	if err != nil {
		return "", 0, err
	}

	return direction, number, nil
}
