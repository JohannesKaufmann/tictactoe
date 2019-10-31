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
			// since a place is still empty,
			// the game can't have ended yet.
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

	// if there is no winner, even though
	// the game has ended, we have a draw.
	hasWinner, _ := b.IsFinished()
	if !hasWinner {
		return true
	}

	return false
}

// TODO: rename to HasWinner?
func (b *Board) IsFinished() (bool, int) {
	for column := 0; column < 3; column++ {
		for row := 0; row < 3; row++ {
			placed := b[column][row]

			if placed == Empty {
				continue
			}

			// we are looking at all neighbors
			// to see wether they have the same
			// player on the field.

			var left int
			var right int
			if row > 0 {
				left = b[column][row-1]
			}
			if row < 2 {
				right = b[column][row+1]
			}

			var up int
			var down int
			if column > 0 {
				up = b[column-1][row]
			}
			if column < 2 {
				down = b[column+1][row]
			}

			// fmt.Printf("c:%d r:%d -> %d \n", column, row, placed)
			// fmt.Printf("\t%d _%d_ %d \n", left, placed, right)

			// If the left and right place is taken
			// by the same player, he has won.
			if left == placed && right == placed {
				return true, placed
			}
			// If the up and down place is taken
			// by the same player, he has won.
			if up == placed && down == placed {
				return true, placed
			}

		}
	}

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
