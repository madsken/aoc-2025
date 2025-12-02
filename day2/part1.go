package main

import (
	"fmt"
	"log"
	"math"
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

	for line := range strings.SplitSeq(fileString, ",") {
		numberRanges := strings.Split(line, "-")
		start, end, err := GetStartAndEnd(numberRanges)
		if err != nil {
			log.Fatalf("Error in GetStartAndEnd: %s", err)
		}

		for i := start; i <= end; i++ {
			numberString := strconv.Itoa(i)
			if len(numberString)%2 != 0 {
				continue
			}

			if numberString[:len(numberString)/2] == numberString[len(numberString)/2:] {
				result += i
			}
		}

	}
	fmt.Println(result)
}

func GetStartAndEnd(ranges []string) (int, int, error) {
	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		return 0, 0, err
	}
	end, err := strconv.Atoi(strings.TrimSpace(ranges[1]))
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func GetDigitCount(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	return int(math.Log10(float64(n))) + 1
}
