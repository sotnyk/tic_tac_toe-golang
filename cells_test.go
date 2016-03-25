package main

import "testing"

func TestWinnerX(t *testing.T) {
	cells := Cells{
		XValue, EmptyValue, EmptyValue,
		EmptyValue, XValue, EmptyValue,
		EmptyValue, EmptyValue, XValue,
	}
	expected := GameStateWinnerX
	actual := cells.CalcWinner()
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestWinnerO(t *testing.T) {
	cells := Cells{
		XValue, EmptyValue, EmptyValue,
		EmptyValue, XValue, EmptyValue,
		OValue, OValue, OValue,
	}
	expected := GameStateWinnerO
	actual := cells.CalcWinner()
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestDraw(t *testing.T) {
	cells := Cells{
		XValue, OValue, XValue,
		XValue, XValue, OValue,
		OValue, XValue, OValue,
	}
	expected := GameStateDraw
	actual := cells.CalcWinner()
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestGameInProgress(t *testing.T) {
	cells := Cells{
		XValue, OValue, XValue,
		EmptyValue, XValue, OValue,
		OValue, XValue, OValue,
	}
	expected := GameStateInProgress
	actual := cells.CalcWinner()
	if actual != expected {
		t.Error("Test failed")
	}
}