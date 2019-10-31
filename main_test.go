package main

import "testing"

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

	if b.HasEnded() != false {
		t.Error("expected HasEnded to return false")
	}

	b.Place(0, 0, PlayerO)
	b.Place(0, 1, PlayerX)
	b.Place(0, 2, PlayerO)

	b.Place(1, 0, PlayerX)
	b.Place(1, 1, PlayerO)
	b.Place(1, 2, PlayerO)

	b.Place(2, 0, PlayerX)
	b.Place(2, 1, PlayerO)
	b.Place(2, 2, PlayerX)

	if b.HasEnded() != true {
		t.Error("expected HasEnded to return true")
	}
}
