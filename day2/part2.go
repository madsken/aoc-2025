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

	for line := range strings.SplitSeq(fileString, ",") {
		numberRanges := strings.Split(line, "-")
		start, end, err := GetStartAndEnd(numberRanges)
		if err != nil {
			log.Fatalf("Error in GetStartAndEnd: %s", err)
		}

		for i := start; i <= end; i++ {
			numberString := strconv.Itoa(i)

			if IsRepeating(numberString) {
				result += i
			}
		}
	}
	fmt.Println(result)
}

// Alternative bruteforce method:
// rotate string one space until one full rotation is one
// if after any one space rotation the string matches the original, then it is repeating
// if no matches after one full rotation, it is not repeating
// Note: This function essentially does the same, but not as bruteforce
// i.e. ss = s + s contains all the permutations of the string rotation
func IsRepeating(s string) bool {
	if len(s) < 2 {
		return false
	}

	ss := s + s
	cut := ss[1 : len(ss)-1]
	return strings.Contains(cut, s)
}
