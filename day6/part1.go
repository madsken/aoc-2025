package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	result := 0

	m := GetMatrixFromFileString(fileString)
	mInv := InvertMatrix(m)

	for _, line := range mInv {
		result += GetMathResult(line)
	}

	fmt.Println(result)
}

func GetMatrixFromFileString(input string) [][]string {
	result := [][]string{}

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}
		l := strings.Fields(line)
		result = append(result, l)
	}

	return result
}

func InvertMatrix(m [][]string) [][]string {
	result := make([][]string, len(m[0]))
	for i := range result {
		result[i] = make([]string, len(m))
	}

	for i := range m {
		for j := range m[0] {
			result[j][i] = m[i][j]
		}
	}

	return result
}

func GetMathResult(line []string) int {
	result := 0
	operation := line[len(line)-1]
	if operation != "+" && operation != "*" {
		panic("Invalid operation")
	}
	if operation == "*" {
		result = 1
	}
	numbers := line[:len(line)-1]

	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		switch operation {
		case "+":
			result += num
		case "*":
			result = result * num
		default:
			panic("Invalid operation")
		}
	}
	return result
}
