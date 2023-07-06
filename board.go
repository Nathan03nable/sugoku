package main

import "fmt"

type Game struct {
	Board [9][9]int
}

type Coordinates struct {
	x int // X is the column
	y int // Y is the row
}

type Move struct {
	Coordinates Coordinates
	Value       int
}

func printBoard(board *[9][9]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
