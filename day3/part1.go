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

	for bank := range strings.SplitSeq(fileString, "\n") {
		if bank == "" {
			continue
		}
		largestSoFar := 0
		largestIdx := 0

		for idx, charge := range bank[:len(bank)-1] {
			cNumber, err := strconv.Atoi(string(charge))
			if err != nil {
				panic(err)
			}

			if cNumber > largestSoFar {
				largestSoFar = cNumber
				largestIdx = idx
			}
		}

		largestSoFar2 := 0
		for _, charge := range bank[largestIdx+1:] {
			cNumber, err := strconv.Atoi(string(charge))
			if err != nil {
				panic(err)
			}

			if cNumber > largestSoFar2 {
				largestSoFar2 = cNumber
			}
		}

		result += (largestSoFar * 10) + largestSoFar2
	}

	fmt.Println(result)
}
