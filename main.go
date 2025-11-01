package main

import (
	"fmt"
	"os"
)

func main() {
	b := readBoard()

	if !validBoard(b) {
		fmt.Println("Malformed board.")
		os.Exit(1)
	}

	fmt.Println("Unsolved: ")
	printBoard(b)

	sweeped, solvable := singletonSweep(b)
	if !solvable || !validBoard(sweeped) {
		fmt.Println("Board not solvable.")
		os.Exit(1)
	}

	if !solved(sweeped) {
		sweeped, solvable = phaseTwo(sweeped)
		if !solvable || !validBoard(sweeped) {
			fmt.Println("Board not solvable.")
			os.Exit(1)
		}
	}

	fmt.Println("Solved: ")
	printBoard(sweeped)
}
