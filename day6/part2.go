package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	operation string
	numbers   []int
}

func Part2(file string) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reding file")
	}
	fileString := string(fileBytes)
	fileString = strings.TrimRight(fileString, "\n")

	lines := strings.Split(fileString, "\n")
	grid := [][]string{}
	for _, l := range lines {
		grid = append(grid, strings.Split(l, ""))
	}

	equations := []Equation{}
	// iterate over cols, not rows
	for col := len(grid[0]) - 1; col >= 0; col-- {
		//ensures problems is not nothing
		if len(equations) == 0 {
			equations = append(equations, Equation{})
		}
		chunk := ""
		for row := 0; row < len(grid); row++ {
			// end of row reached, handle the chunk and potentially an operation
			if row == len(grid)-1 {
				chunk = strings.TrimSpace(chunk)

				// empty col
				if chunk == "" {
					continue
				}

				chunkInt, err := strconv.Atoi(chunk)
				if err != nil {
					panic(err)
				}

				equations[len(equations)-1].numbers = append(equations[len(equations)-1].numbers, chunkInt)

				// handle op
				if grid[row][col] != " " {
					op := grid[row][col]
					equations[len(equations)-1].operation = op

					// add new problem to begin chunking
					if col != 0 {
						equations = append(equations, Equation{})
					}
				}
			} else {
				chunk += grid[row][col]
			}
		}
	}

	fmt.Println(scoreProblems(equations))
}

func scoreProblems(problems []Equation) int {
	result := 0

	for _, problem := range problems {
		switch problem.operation {
		case "*":
			s := 1
			for _, n := range problem.numbers {
				s *= n
			}
			result += s
		case "+":
			s := 0
			for _, n := range problem.numbers {
				s += n
			}
			result += s
		default:
			panic("invalid op")
		}
	}

	return result
}
