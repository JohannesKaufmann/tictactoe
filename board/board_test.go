package board

import (
	"fmt"
	"testing"
)

func TestBoardPlace(t *testing.T) {
	var b Board

	err := b.Place(0, 0, 1)
	if err != nil {
		t.Fail()
	}

	err = b.Place(0, 1, 1)
	if err != nil {
		t.Fail()
	}

	// now placing where other player already is
	err = b.Place(0, 0, 2)
	if err == nil {
		t.Fail()
	}
}

func TestBoardHasEnded(t *testing.T) {
	var b Board

	if b.HasNoSpaceLeft() != false {
		t.Error("expected HasNoSpaceLeft to return false")
	}

	b.Place(0, 0, PlayerO)
	b.Place(0, 1, PlayerX)
	b.Place(0, 2, PlayerO)
	b.Place(0, 3, PlayerO)

	b.Place(1, 0, PlayerX)
	b.Place(1, 1, PlayerO)
	b.Place(1, 2, PlayerO)
	b.Place(1, 3, PlayerO)

	b.Place(2, 0, PlayerX)
	b.Place(2, 1, PlayerO)
	b.Place(2, 2, PlayerX)
	b.Place(2, 3, PlayerO)

	b.Place(3, 0, PlayerX)
	b.Place(3, 1, PlayerO)
	b.Place(3, 2, PlayerX)
	b.Place(3, 3, PlayerO)

	if b.HasNoSpaceLeft() != true {
		t.Error("expected HasNoSpaceLeft to return true")
	}
}

func TestBoardIsFinished(t *testing.T) {
	var tests = []struct {
		Name      string
		Board     Board
		HasWinner bool
		Winner    Player
		HasDraw   bool
	}{
		{
			Board: [4][4]Player{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			HasWinner: false,
		},
		{
			Board: [4][4]Player{
				{0, 1, 1, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			HasWinner: false,
		},
		{
			Board: [4][4]Player{
				{2, 1, 1, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			HasWinner: false,
		},
		{
			Name: "horizontal win (on the first column)",
			Board: [4][4]Player{
				{1, 1, 1, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			HasWinner: true,
			Winner:    1,
		},
		{
			Name: "horizontal win (on the last column)",
			Board: [4][4]Player{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
			},
			HasWinner: true,
			Winner:    1,
		},
		{
			Name: "vertical win",
			Board: [4][4]Player{
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			HasWinner: true,
			Winner:    1,
		},
		{
			Name: "slanted win from the top left",
			Board: [4][4]Player{
				{2, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			HasWinner: true,
			Winner:    2,
		},
		{
			Name: "slanted win from the top right",
			Board: [4][4]Player{
				{0, 0, 0, 2},
				{0, 0, 2, 0},
				{0, 2, 0, 0},
				{2, 0, 0, 0},
			},
			HasWinner: true,
			Winner:    2,
		},
		{
			Name: "real life 1: draw",
			Board: [4][4]Player{
				{1, 2, 1, 2},
				{2, 1, 1, 2},
				{2, 1, 2, 1},
				{1, 2, 1, 2},
			},
			HasDraw:   true,
			HasWinner: false,
		},
		{
			Name: "real life 2: winner X",
			Board: [4][4]Player{
				{1, 1, 1, 1}, // win
				{2, 2, 1, 2},
				{2, 0, 2, 2},
				{0, 0, 0, 0},
			},
			HasDraw:   false,
			HasWinner: true,
			Winner:    1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Board.IsDraw() != test.HasDraw {
				t.Errorf("expcted IsDraw to be %v", test.HasDraw)
			}
			hasWinner, winner := test.Board.HasWinner()

			if hasWinner != test.HasWinner {
				t.Errorf("expcted HasWinner to be %v", test.HasWinner)
			}
			if winner != test.Winner {
				t.Errorf("expcted Winner to be %v", test.Winner)
			}
		})
	}
}

func TestBoardGet(t *testing.T) {
	var b Board
	b.Place(0, 0, 5)
	b.Place(1, 1, 6)
	b.Place(2, 2, 7)

	if b.Get(0, 0) != 5 {
		fmt.Println(b.Get(0, 0))
		t.Error("expected Get to return 5")
	}
	if b.Get(1, 1) != 6 {
		t.Error("expected Get to return 6")
	}
	if b.Get(2, 2) != 7 {
		t.Error("expected Get to return 7")
	}

	if b.Get(-1, -1) != -1 {
		t.Error("expected Get to return -1 but not panic")
	}
	if b.Get(3, 3) != -1 {
		t.Error("expected Get to return -1 but not panic")
	}
}
