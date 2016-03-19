package main

type GameState int

const (
	GameStateInProgress GameState = iota
	GameStateWinnerX
	GameStateWinnerO
	GameStateDraw
)

