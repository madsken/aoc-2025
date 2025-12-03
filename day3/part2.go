package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2(file string) {
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
		number, err := strconv.Atoi(GetLargest12DigitNumber(string(bank)))
		if err != nil {
			panic(err)
		}
		result += number
	}

	fmt.Println(result)
}

func GetLargest12DigitNumber(number string) string {
	digitsToRemove := len(number) - 12
	largestNumber := make([]byte, 0, 12)

	for i := 0; i < len(number); i++ {
		digit := number[i]

		for digitsToRemove > 0 && len(largestNumber) > 0 && largestNumber[len(largestNumber)-1] < digit {
			largestNumber = largestNumber[:len(largestNumber)-1]
			digitsToRemove -= 1
		}

		largestNumber = append(largestNumber, digit)
	}

	return string(largestNumber[:12])
}
