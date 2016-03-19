package main

type CellType int

const (
	EmptyValue CellType = iota
	XValue
	OValue
)

func (cell CellType) ToString() string {
	switch cell {
	case EmptyValue:
		return " "
	case XValue:
		return "X"
	case OValue:
		return "O"
	}
	return "?"
}

type Cells []CellType

var waysToWin = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6}}

func (cells Cells) CalcWinner() GameState {
	for _, row := range waysToWin {
		if cells[row[0]] == cells[row[1]] && cells[row[0]] == cells[row[2]] && cells[row[0]] != EmptyValue {
			if (cells[row[0]] == XValue) {
				return GameStateWinnerX
			}else {
				return GameStateWinnerO
			}
		}
	}
	for _, cell := range cells {
		if cell == EmptyValue {
			return GameStateInProgress
		}
	}
	return GameStateDraw
}

func (cells Cells) LegalMoves() []int {
	result := make([]int, 0, BoardSize)
	for i, c := range cells {
		if c == EmptyValue {
			result = append(result, i)
		}
	}
	return result
}

func (cells Cells) IsLegalMove(move int) bool {
	if move < 0 || move > BoardSize {
		return false
	}
	return cells[move] == EmptyValue
}