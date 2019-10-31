package main

import "fmt"

var (
	Empty   = 0
	PlayerX = 1
	PlayerO = 2
)

// Board is a Custom Type so that we can
// attach methods to it, that make it
// easier to interact with.
type Board [3][3]int

func (b *Board) Place(a, c int, player int) error {
	b[a][c] = player
	// TODO: check wether the place is already taken
	// by another player. Then return error
	return nil

}
func (b *Board) IsDraw() bool {
	// TODO: loop over the array
	// then look if every field is not empty

	return true
}
func (b *Board) IsFinished() (bool, int) {
	// TODO: declare which person won?

	return false, 0
}

func (b *Board) Print() {
	fmt.Printf("\n\n\t0 1 2\n")
	for i, e := range b {
		fmt.Printf("\n%d\t", i)
		for _, x := range e {
			fmt.Print(x, " ")
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
