package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func Part2(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	ranges, _ = GetRangesAndIDs(fileString)

	fmt.Println(GetCountInRanges(ranges))
}

func GetCountInRanges(ranges [][2]int) int {
	// Strategy: Merge ranges (to reduce duplicate numbers) and find the differences for each range
	if len(ranges) == 0 {
		return 0
	}
	count := 0

	// Sort, to make it easier to merge
	// Sort by lowest starting number
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	mergedRanges := make([][2]int, 0)
	current := ranges[0]

	for _, r := range ranges[1:] {
		// If the range we are looking at starts inside the previous range, we merge
		// (also the end of the range we are looking at needs to extend beyond the previous range)
		if r[0] <= current[1] {
			if r[1] > current[1] {
				current[1] = r[1]
			}
		} else {
			// When ranges dont overlap, we append the current to the merged ranges list. No more ranges will overlap (due to being sorted)
			mergedRanges = append(mergedRanges, current)
			current = r
		}
	}
	mergedRanges = append(mergedRanges, current)

	for _, r := range mergedRanges {
		count += r[1] - r[0] + 1 // +1 because the end is inclusive
	}

	return count
}
