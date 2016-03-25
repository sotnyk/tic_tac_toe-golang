package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
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