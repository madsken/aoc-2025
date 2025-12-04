package main

import (
	"fmt"
	"log"
	"os"
)

func Part2(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	grid := GetGrid(fileString)

	result := 0
	prevResult := -1
	for result != prevResult {
		prevResult = result
		result += GetRemovedScrolls(grid)
	}

	fmt.Println(result)
}

func GetRemovedScrolls(grid [][]string) int {
	result := 0
	scrollsToRemove := [][2]int{}
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "." {
				continue
			}
			rolls := CountPaperRolls(grid, i, j)
			if rolls < 4 {
				result += 1
				scrollsToRemove = append(scrollsToRemove, [2]int{i, j})
			}
		}
	}

	for _, scrollToRemove := range scrollsToRemove {
		grid[scrollToRemove[0]][scrollToRemove[1]] = "."
	}

	return result
}
