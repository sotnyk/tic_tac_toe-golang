package main

import "testing"

func constructBoardForSecondMove() Board{
	b := NewBoard()
	b.initGame(true)
	b.cells[0] = XValue
	return b
}

func TestSecondMove(t *testing.T){
	b := constructBoardForSecondMove()
	expected := 4
	actual := makeComputerMove(&b)
	if (expected != actual){
		t.Errorf("Test failed, expected '%d', got '%d'", expected, actual)
	}
}

func TestWeightFor4(t *testing.T){
	b := constructBoardForSecondMove()
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 4)
	if (expected != actual){
		t.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeightFor8(t *testing.T){
	b := constructBoardForSecondMove()
	var expected float32 = -1
	actual := calcFitnessCompMove(&b, 8)
	if (expected != actual){
		t.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}