package board

import (
	"errors"
)

type Player int

func (p Player) String() string {
	if p == 1 {
		return "X"
	} else if p == 2 {
		return "O"
	}
	return "_"
}

var (
	Empty   Player = 0
	PlayerX Player = 1
	PlayerO Player = 2
)

var PlaceTakenErr = errors.New("the place is already taken")

// Board is a Custom Type so that we can
// attach methods to it, that make it
// easier to interact with.
type Board [4][4]Player

// Place places the player on the Board. Returns an error
// if the place is already taken by either player.
func (b *Board) Place(column, row int, player Player) error {
	val := b[column][row]
	if val != 0 {
		return PlaceTakenErr
	}

	b[column][row] = player

	return nil
}

func (b *Board) HasNoSpaceLeft() bool {
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
	if !b.HasNoSpaceLeft() {
		return false
	}

	// if there is no winner, even though
	// the game has ended, we have a draw.
	hasWinner, _ := b.HasWinner()
	if !hasWinner {
		return true
	}

	return false
}

// Get returns the element at this position
// but returning -1 if the place is outisde the board.
// Makes it safer to access places on the board
func (b *Board) Get(column, row int) Player {
	// check wether the index is inside the bounds
	if column < 0 || column > 2 {
		return -1
	}
	if row < 0 || row > 2 {
		return -1
	}

	return b[column][row]
}

func areSamePlayers(elements []Player) bool {
	var lastElement = elements[0]

	// Empty is not a real Player
	if lastElement == Empty {
		return false
	}

	for _, elem := range elements {
		if elem != lastElement {
			return false
		}
	}

	return true
}

func (b *Board) HasWinner() (bool, Player) {
	// horizontal check
	for _, row := range b {
		// `row[:]` is a trick to convert a fixed size
		// array into a slice. That way I don't have to
		// hardcode the size of the array.
		if areSamePlayers(row[:]) {
			return true, row[0]
		}
	}

	// vertical check
	var inverted [4][4]Player
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			inverted[j][i] = b[i][j]
		}
	}
	for _, elems := range inverted {
		if areSamePlayers(elems[:]) {
			return true, elems[0]
		}
	}

	// diagonal check
	// TODO: implement diagonal check

	return false, 0
}
