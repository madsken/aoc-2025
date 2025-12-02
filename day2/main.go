package main

import (
	"os"
)

func main() {
	input := ""
	if len(os.Args) < 2 {
		input = "input.txt"
	} else {
		input = os.Args[1]
	}

	Part1(input)
	Part2(input)
}
