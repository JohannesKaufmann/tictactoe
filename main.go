package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// Get returns the element at this position
// but returning -1 if the place is outisde the board.
// Makes it safer to access places on the board
func (b *Board) Get(column, row int) int {
	// check wether the index is inside the bounds
	if column < 0 || column > 2 {
		return -1
	}
	if row < 0 || row > 2 {
		return -1
	}

	return b[column][row]
}

// TODO: rename to HasWinner?
func (b *Board) IsFinished() (bool, int) {
	// This function acts on the assumption
	// that the only way to win, is to be in
	// the "Middle" and the surrounding place
	// is taken by the same player.

	for column := 0; column < 3; column++ {
		for row := 0; row < 3; row++ {
			placed := b[column][row]

			if placed == Empty {
				continue
			}

			// we are looking at all neighbors
			// to see wether they have the same
			// player on the field.
			var left = b.Get(column, row-1)
			var right = b.Get(column, row+1)

			var up = b.Get(column-1, row)
			var down = b.Get(column+1, row)

			var topLeft = b.Get(column-1, row-1)
			var topRight = b.Get(column-1, row+1)
			var bottomLeft = b.Get(column+1, row-1)
			var bottomRight = b.Get(column+1, row+1)

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

			// slanted from the topLeft
			if topLeft == placed && bottomRight == placed {
				return true, placed
			}
			// slanted from the topRight
			if topRight == placed && bottomLeft == placed {
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
	reader := bufio.NewReader(os.Stdin)

	var b Board
	for !b.HasEnded() {
		fmt.Println("- - ROUND number - -")
		b.Print()

		if hasWinner, winner := b.IsFinished(); hasWinner {
			fmt.Println("THERE IS A WINNER")
			fmt.Println("it is player", winner)
		}

		fmt.Print("PLAYER X(1) column: ")
		column, _ := reader.ReadString('\n')
		fmt.Print("PLAYER X(1) row: ")
		row, _ := reader.ReadString('\n')

		fmt.Println(column, row)

		co, err := strconv.Atoi(strings.TrimSpace(column))
		if err != nil {
			log.Fatal(err)
		}
		ro, err := strconv.Atoi(strings.TrimSpace(row))
		if err != nil {
			log.Fatal(err)
		}

		b.Place(co, ro, PlayerX)
	}
	if b.IsDraw() {
		fmt.Println("THERE IS A DRAW")
	}

	// b.Print()
	// b.Place(0, 1, PlayerX)
	// b.Print()
	// fmt.Println(b.IsDraw())

	// TODO: game loop
	// TODO: prompt for user action
}
