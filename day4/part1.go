package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var directions = [][2]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func Part1(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	grid := GetGrid(fileString)
	result := 0

	result = GetRemovedScrolls(grid)
	fmt.Println(result)
}

func GetGrid(input string) [][]string {
	lines := strings.Split(input, "\n")

	grid := make([][]string, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	return grid
}

func CountPaperRolls(grid [][]string, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	result := 0

	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]

		if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols {
			if grid[nextRow][nextCol] == "@" {
				result += 1
			}
		}
	}

	return result
}
