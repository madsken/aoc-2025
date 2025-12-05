package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ranges [][2]int
var idsToCheck []int

func Part1(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	ranges, idsToCheck = GetRangesAndIDs(fileString)

	validIDs := make(map[int]bool)

	for _, id := range idsToCheck {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				validIDs[id] = true
			}
		}
	}

	fmt.Println(len(validIDs))
}

func GetRangesAndIDs(input string) ([][2]int, []int) {
	parts := strings.Split(input, "\n\n")
	r := [][2]int{}
	ids := []int{}

	for line := range strings.SplitSeq(parts[0], "\n") {
		splits := strings.Split(line, "-")

		start, err := strconv.Atoi(splits[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(splits[1])
		if err != nil {
			panic(err)
		}

		r = append(r, [2]int{start, end})
	}

	for line := range strings.SplitSeq(parts[1], "\n") {
		if line == "" {
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		ids = append(ids, id)
	}

	return r, ids
}
