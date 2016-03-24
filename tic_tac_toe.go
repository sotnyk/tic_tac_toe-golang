package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

var cells_diag = Cells{
XValue, EmptyValue, EmptyValue,
EmptyValue, OValue, EmptyValue,
EmptyValue, EmptyValue, XValue,
}

func main() {
	//TestWeightFor8_tst()
	//TestWeightFor4_tst()
	//TestWeightFor_tst()
	//TestWeight_tst4()
//	TestFitnessCompMove()
/*	TestSort()
	TestFitnessHumanMove()
	test_field(
		cells_diag,
		[]float32{
			4, 0, -1,
			0, 4, 0,
			-1, 0, 4,
		})
*/
	displayInstructions()
	for {
		board := initBoard()
		board.StartGame(askYesNo("Do you wannna play as " + XValue.ToString() + "? (y/n): "))
		if !askYesNo("Do you wanna repeat? (y/n)") {
			break
		}
	}
	fmt.Println()
	fmt.Println("Good luck!")
}

func initBoard() Board {
	result := NewBoard()
	return result
}

func askYesNo(message string) bool {
	for {
		fmt.Print(message)
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(response)
		if response == "y" {
			return true
		}
		if response == "n" {
			return false
		}
	}
}

func askNumber(message string, min int, max int) int {
	for {
		fmt.Print(message)
		var response string
		fmt.Scanln(&response)
		result, err := strconv.Atoi(response)
		if err == nil {
			if result >= min && result <= max {
				return result
			} else {
				fmt.Println("Please, enter value in range %d...%d", min, max)
			}
		} else {
			fmt.Println("Please, enter valid integer")
		}
	}
}

func displayInstructions() {
	fmt.Println(
		`    Welcome to the tic-tac-toe game.
		To make a move, enter a digit from 1 to 9. This digit accords
		to the board's fields as it defined below:

		1 | 2 | 3
		--+---+--
		4 | 5 | 6
		--+---+--
		7 | 8 | 9`)
}

func constructBoardForSecondMove_tst() Board {
	b := NewBoard()
	b.initGame(true)
	b.cells[0] = XValue
	return b
}

func TestWeightFor1_tst() {
	b := NewBoard()
	b.initGame(true)
	copy(b.cells, cells_diag)
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 1)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestFitnessHumanMove(){
/*
	position := Cells{

		XValue, OValue, EmptyValue,
		EmptyValue, OValue, EmptyValue,
		EmptyValue, EmptyValue, XValue,
	}
*/
	position := Cells{

		XValue, OValue, EmptyValue,
		EmptyValue, OValue, EmptyValue,
		OValue, XValue, XValue,
	}

	b := NewBoard()
	b.initGame(true)
	copy(b.cells, position)
	var expected float32 = 0
	actual := calcFitnessHumanMove(&b, 2)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestFitnessCompMove(){
	position := Cells{
		XValue, OValue, XValue,
		EmptyValue, OValue, EmptyValue,
		OValue, XValue, XValue,
	}
	b := NewBoard()
	b.initGame(true)
	copy(b.cells, position)
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 5)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeightFor8_tst() {
	b := constructBoardForSecondMove_tst()
	var expected float32 = -1
	actual := calcFitnessCompMove(&b, 8)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeightFor4_tst() {
	b := constructBoardForSecondMove_tst()
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 4)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeightFor_tst() {
	b := NewBoard()
	b.initGame(true)
	b.cells = Cells{
		OValue, OValue, XValue,
		XValue, XValue, EmptyValue,
		OValue, XValue, EmptyValue,
	}
	var expected float32 = -1
	actual := calcFitnessCompMove(&b, 8)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
	expected = 0
	actual = calcFitnessCompMove(&b, 5)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeight_tst3() {
	b := NewBoard()
	b.initGame(true)
	initCells := func() Cells{
		return Cells{
			XValue, XValue, EmptyValue,
			OValue, OValue, EmptyValue,
			XValue, OValue, EmptyValue,
		}
	}
	b.cells = initCells()
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 2)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}

	b.cells = initCells()
	expected = 1
	actual = calcFitnessCompMove(&b, 5)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}

	b.cells = initCells()
	expected = -1
	actual = calcFitnessCompMove(&b, 8)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func TestWeight_tst4() {
	b := NewBoard()
	b.initGame(true)
	initCells := func() Cells{
		return Cells{
			XValue, XValue, EmptyValue,
			OValue, OValue, EmptyValue,
			XValue, EmptyValue, EmptyValue,
		}
	}
	b.cells = initCells()
	var expected float32 = 0
	actual := calcFitnessCompMove(&b, 2)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}

	b.cells = initCells()
	expected = 1
	actual = calcFitnessCompMove(&b, 5)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}

	b.cells = initCells()
	expected = -1
	actual = calcFitnessCompMove(&b, 8)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}

	b.cells = initCells()
	expected = -1
	actual = calcFitnessCompMove(&b, 7)
	if (expected != actual) {
		fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
	}
}

func test_field(position Cells, weights []float32){
	for i, c := range position{
		if c==EmptyValue {
			b := NewBoard()
			b.initGame(true)
			copy(b.cells, position)
			var expected float32 = weights[i]
			actual := calcFitnessCompMove(&b, i)
			if (expected != actual) {
				fmt.Errorf("Test failed, expected '%f', got '%f'", expected, actual)
			}
		}
	}
}

func TestSort()  {
	wm := weightedMoves{
		weightedMove{move:0, weight:0},
		weightedMove{move:1, weight:1},
		weightedMove{move:2, weight:-1},
	}
	fmt.Println(wm)
	sort.Sort(wm)
	fmt.Println(wm)
	sort.Sort(sort.Reverse(wm))
	fmt.Println(wm)
}