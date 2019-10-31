package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/JohannesKaufmann/tictactoe/board"
)

// printBoard prints the current board in a grid
func printBoard(b *board.Board) {
	fmt.Printf("\n\nrow\t\t0 1 2\n")
	for i, column := range b {
		fmt.Printf("\ncolumn %d\t", i)
		for _, row := range column {
			fmt.Print(row, " ")
		}
	}
	fmt.Println()
}

// askForPlacement is responsible for prompting the
// user for input and returning the values in a
// usable format.
func askForPlacement(reader *bufio.Reader) (int, int, error) {
	fmt.Print("column: ")
	columnString, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}

	fmt.Print("row: ")
	rowString, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}

	// make sure to trim all whitespace
	columnString = strings.TrimSpace(columnString)
	rowString = strings.TrimSpace(rowString)

	column, err := strconv.Atoi(columnString)
	if err != nil {
		return 0, 0, err
	}
	row, err := strconv.Atoi(rowString)
	if err != nil {
		log.Fatal(err)
	}

	return column, row, nil
}

// askForPlacementAndPlace is responsible for handling errors
// and letting the user try again.
func askForPlacementAndPlace(reader *bufio.Reader, b *board.Board, player board.Player) {
	fmt.Println("\n\tPlayer", player, "it is your turn!")

	var needValidInput = true
	for needValidInput {
		column, row, err := askForPlacement(reader)
		if err != nil {
			fmt.Println("we encountered an unexpected error, try again:", err)
			continue
		}

		err = b.Place(column, row, player)
		if err != nil && err == board.PlaceTakenErr {
			fmt.Println("the place that you wanted is already taken, please try again!")
			continue
		} else if err != nil {
			fmt.Println("we encountered an unexpected error, try again:", err)
			continue
		}

		needValidInput = false
	}
}

// checkForEnd checks wether the game has ended and
// then quits the program. Quitting should probably
// be done in the main function...
func checkForEnd(b *board.Board) {
	if hasWinner, winner := b.HasWinner(); hasWinner {
		fmt.Println("\n\nWE HAVE A WINNER!")
		fmt.Println("it is", winner, "!")

		os.Exit(0)
	}
	if b.IsDraw() {
		fmt.Println("\n\nWE HAVE A DRAW!")
		fmt.Println("try again :D")

		os.Exit(0)
	}
}

func main() {
	var b board.Board
	reader := bufio.NewReader(os.Stdin)

	var round int
	var gameInProgress = true
	for gameInProgress {
		round++

		fmt.Println("\n\t\tROUND", round)

		// Ask Player X
		printBoard(&b)
		askForPlacementAndPlace(reader, &b, board.PlayerX)
		checkForEnd(&b)

		// Ask Player O
		printBoard(&b)
		askForPlacementAndPlace(reader, &b, board.PlayerO)
		checkForEnd(&b)
	}
}
