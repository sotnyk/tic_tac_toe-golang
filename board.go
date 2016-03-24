package main

import (
	"fmt"
	"sort"
	"math/rand"
)

type weightedMove struct {
	move   int
	weight float32
}
type weightedMoves []weightedMove

func (slice weightedMoves) Len() int {
	return len(slice)
}

func (slice weightedMoves) Less(i, j int) bool {
	return slice[i].weight < slice[j].weight;
}

func (slice weightedMoves) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type Board struct {
	cells               Cells
	humanCell           CellType
	computerCell        CellType
	humanWinnerState    GameState
	computerWinnerState GameState
	isHumanOrder        bool
}

const BoardSize = 9

func NewBoard() Board {
	result := Board{}
	result.cells = make([]CellType, BoardSize, BoardSize)
	return result
}

func (board *Board) Display() {
	cells := board.cells
	fmt.Println()
	fmt.Println("\t" + cells[0].ToString() + " | " + cells[1].ToString() + " | " + cells[2].ToString())
	fmt.Println("\t--+---+--")
	fmt.Println("\t" + cells[3].ToString() + " | " + cells[4].ToString() + " | " + cells[5].ToString())
	fmt.Println("\t--+---+--")
	fmt.Println("\t" + cells[6].ToString() + " | " + cells[7].ToString() + " | " + cells[8].ToString())
	fmt.Println()
}

func (board *Board) StartGame(humanFirst bool) {
	board.initGame(humanFirst)
	cells := board.cells

	for cells.CalcWinner() == GameStateInProgress {
		board.Display()
		if board.isHumanOrder {
			cells[makeHumanMove(board)] = board.humanCell
		} else {
			cells[makeComputerMove(board)] = board.computerCell
		}
		board.isHumanOrder = !board.isHumanOrder
	}
	board.Display()

	board.reportTheGameResult()
}

func makeHumanMove(board *Board) int {
	cells := board.cells
	move := -1
	for !cells.IsLegalMove(move) {
		move = askNumber("Your move. Please, select one of empty field (1...9): ", 1, BoardSize + 1) - 1
	}
	return move
}

var bestMoves = [][]int{
	{4, 0, 2, 6, 8, 1, 3, 5, 7},
	{8, 6, 2, 0, 4, 7, 3, 5, 1},
	{6, 2, 0, 8, 4, 7, 3, 5, 1},
	{4, 6, 2, 8, 0, 7, 3, 5, 1},
	{2, 6, 0, 8, 4, 5, 7, 1, 3}}

func makeComputerMove(board *Board) int {
	cells := board.cells
	fmt.Print("Computer's move is ")
	legal := cells.LegalMoves()
	// Check if one-move win position
	for _, move := range legal {
		cells[move] = board.computerCell
		if cells.CalcWinner() == board.computerWinnerState {
			return move
		}
		cells[move] = EmptyValue
	}
	// Check if one-move defeat position
	for _, move := range legal {
		cells[move] = board.humanCell
		if cells.CalcWinner() == board.humanWinnerState {
			return move
		}
		cells[move] = EmptyValue
	}

	possibleMoves := make(weightedMoves, 0, BoardSize)
	for _, m := range bestMoves[rand.Intn(len(bestMoves))] {
		if cells.IsLegalMove(m) {
			possibleMoves = append(possibleMoves, weightedMove{m, calcFitnessCompMove(board, m)})
		}
	}
	sort.Sort(sort.Reverse(possibleMoves))
	//sort.Sort(possibleMoves)
	move := possibleMoves[0].move
	fmt.Println(move)
	return move
}

func calcFitnessCompMove(board *Board, move int) float32 {
	var result float32
	cells := board.cells

	cells[move] = board.computerCell
	switch cells.CalcWinner() {
	case board.computerWinnerState:
		result = 1
	case board.humanWinnerState:
		panic("Something wrong! This state should be unreachable")
	case GameStateDraw:
		result = 0
	case GameStateInProgress:
		possibleMoves := make(weightedMoves, 0, BoardSize)
		for i, c := range cells {
			if c == EmptyValue {
				possibleMoves = append(possibleMoves, weightedMove{i, calcFitnessHumanMove(board, i)})
			}
		}
		sort.Sort(possibleMoves)
		//sort.Sort(sort.Reverse(possibleMoves))
		result = possibleMoves[0].weight
	default:
		panic("Something wrong! This state should be unreachable")
	}
	cells[move] = EmptyValue
	return result
}

func calcFitnessHumanMove(board *Board, move int) float32 {
	var result float32
	cells := board.cells

	cells[move] = board.humanCell
	switch cells.CalcWinner() {
	case board.computerWinnerState:
		panic("Something wrong! This state should be unreachable")
	case board.humanWinnerState:
		result = -1
	case GameStateDraw:
		result = 0
	case GameStateInProgress:
		possibleMoves := make(weightedMoves, 0, BoardSize)
		for i, c := range cells {
			if c == EmptyValue {
				possibleMoves = append(possibleMoves, weightedMove{i, calcFitnessCompMove(board, i)})
			}
		}
		sort.Sort(sort.Reverse(possibleMoves))
		//sort.Sort(possibleMoves)
		result = possibleMoves[0].weight
	default:
		panic("Something wrong! This state should be unreachable")
	}
	cells[move] = EmptyValue
	return result
}

func (board *Board) reportTheGameResult() {
	switch board.cells.CalcWinner() {
	case GameStateDraw:
		fmt.Println("Game result is DRAW. Nobody won.")
	case board.humanWinnerState:
		fmt.Println("You won! Congratulates!")
	case board.computerWinnerState:
		fmt.Println("I am winner! Don't worry, baby!")
	default:
		panic("Something wrong! This state should be unreachable")
	}
}

func (board *Board) initGame(humanFirst bool) {
	board.isHumanOrder = humanFirst
	if humanFirst {
		board.humanCell = XValue
		board.computerCell = OValue
		board.humanWinnerState = GameStateWinnerX
		board.computerWinnerState = GameStateWinnerO
	} else {
		board.humanCell = OValue
		board.computerCell = XValue
		board.humanWinnerState = GameStateWinnerO
		board.computerWinnerState = GameStateWinnerX
	}
}