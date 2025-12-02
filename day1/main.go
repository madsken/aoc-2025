package main

import (
	"os"
	"strconv"
)

func main() {
	input := ""
	if len(os.Args) < 2 {
		input = "input.txt"
	} else {
		input = os.Args[1]
	}

	Part1(input)
	Part2(input)
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

func ProcessLineWithRounds(line string) (string, int, int, error) {
	direction := string(line[0])
	rotations := 0
	numberString := string(line[1:])

	number, err := strconv.Atoi(numberString)
	if err != nil {
		return "", 0, 0, err
	}

	if number > 99 {
		rotations = number / 100
	}

	number = number % 100
	return direction, number, rotations, nil
}
