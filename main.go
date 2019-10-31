package main

import (
	"errors"
	"fmt"
)

var (
	Empty   = 0
	PlayerX = 1
	PlayerO = 2
)

var PlaceTakenErr = errors.New("the place is already taken")

// Board is a Custom Type so that we can
// attach methods to it, that make it
// easier to interact with.
type Board [3][3]int

// Place places the player on the Board. Returns an error
// if the place is already taken by either player.
func (b *Board) Place(column, row int, player int) error {
	val := b[column][row]
	if val != 0 {
		return PlaceTakenErr
	}

	b[column][row] = player

	return nil
}

func (b *Board) HasEnded() bool {
	for _, column := range b {
		for _, row := range column {
			if row == Empty {
				return false
			}
		}
	}

	return true
}

func (b *Board) IsDraw() bool {
	// A draw can only occur when all
	// places on the board are filled.
	if !b.HasEnded() {
		return false
	}

	// TODO: loop over the array
	// then look if every field is not empty

	return true
}
func (b *Board) IsFinished() (bool, int) {
	// TODO: declare which person won?

	return false, 0
}

func (b *Board) Print() {
	fmt.Printf("\n\nrow\t\t0 1 2\n")
	for i, column := range b {
		fmt.Printf("\ncolumn %d\t", i)
		for _, row := range column {
			fmt.Print(row, " ")
		}
	}
	fmt.Println()
}

func main() {
	// TODO: place game locic in own package
	fmt.Println("- - - Tic-Tac-Toe - - -")

	var b Board
	b.Print()
	b.Place(0, 1, PlayerX)
	b.Print()

	fmt.Println(b.IsDraw())

	// TODO: game loop
	// TODO: prompt for user action
}
