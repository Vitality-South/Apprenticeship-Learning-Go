// Copyright 2023 Vitality South, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"math/rand"

	"github.com/nexidian/gocliselect"
)

// GameMode represents the game mode.
type GameMode int

const (
	Easy GameMode = iota
	Normal
	Hard
	Exit
)

// Player represents a player in the game.
type Player string

const (
	HumanPlayer Player = "X"
	AIPlayer    Player = "O"
	EmptySpot          = " "
	BoardSize          = 3
)

// Coordinate represents a coordinate on the board with 0 0 being the top left spot.
type Coordinate struct {
	X, Y int
}

// GameBoard represents the spots on the board that can be played.
type GameBoard [BoardSize][BoardSize]string

// Board represents the game board.
type Board struct {
	spaces          GameBoard
	mode            GameMode
	winningPatterns [][BoardSize]Coordinate
}

// String returns the string representation of the player.
func (p Player) String() string {
	switch p {
	case HumanPlayer:
		return "Human"
	case AIPlayer:
		return "AI"
	}

	return "Unknown"
}

// SelectGameMode displays a menu to select a game mode.
func SelectGameMode() GameMode {
	menu := gocliselect.NewMenu("Choose a game mode")

	menu.AddItem("New easy game", "Easy")
	menu.AddItem("New normal game", "Normal")
	menu.AddItem("New hard game", "Hard")
	menu.AddItem("Exit the game", "Exit")

	switch menu.Display() {
	case "Hard":
		return Hard
	case "Normal":
		return Normal
	case "Easy":
		return Easy
	}

	return Exit
}

// ShowTutorial prints the tutorial for the game.
func (b Board) ShowTutorial() {
	fmt.Println("")
	fmt.Println("        X     ")
	fmt.Println("    0   1   2 ")
	fmt.Println("      |   |   ")
	fmt.Println("  0   |   |   ")
	fmt.Println("   ___|___|___")
	fmt.Println("      |   |   ")
	fmt.Println("Y 1   |   |   ")
	fmt.Println("   ___|___|___")
	fmt.Println("      |   |   ")
	fmt.Println("  2   |   |   ")
	fmt.Println("      |   |   ")

	fmt.Println("To play, enter the Y coordinate first, then the X coordinate. An example is 2 0 for the bottom left spot.")
	fmt.Println("")
}

// Print prints the current board.
func (b Board) Print() {
	fmt.Printf("\nCurrent Game Board\n")

	for i := 0; i < BoardSize; i++ {
		fmt.Println("   |   |   ")
		fmt.Printf(" %s | %s | %s \n", b.spaces[i][0], b.spaces[i][1], b.spaces[i][2])

		if i < 2 {
			fmt.Println("___|___|___")
		}
	}

	fmt.Println("   |   |   ")
	fmt.Println("")
}

// Init initializes the board for a new game and sets every play position to EmptySpot.
func (b *Board) Init() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b.spaces[i][j] = EmptySpot
		}
	}

	b.winningPatterns = [][BoardSize]Coordinate{
		{{0, 0}, {0, 1}, {0, 2}}, // row 1
		{{1, 0}, {1, 1}, {1, 2}}, // row 2
		{{2, 0}, {2, 1}, {2, 2}}, // row 3
		{{0, 0}, {1, 0}, {2, 0}}, // column 1
		{{0, 1}, {1, 1}, {2, 1}}, // column 2
		{{0, 2}, {1, 2}, {2, 2}}, // column 3
		{{0, 0}, {1, 1}, {2, 2}}, // diagonal 1
		{{0, 2}, {1, 1}, {2, 0}}, // diagonal 2
	}
}

// NewBoard returns a new board for the given game mode.
func NewBoard(m GameMode) *Board {
	b := Board{mode: m}

	b.Init()

	return &b
}

// Play plays the given coordinates for the human player.
func (b *Board) Play(x, y int) {
	b.play(x, y, HumanPlayer)
}

// play plays the given coordinates for the given player.
func (b *Board) play(x, y int, p Player) {
	b.spaces[x][y] = string(p)
}

// firstEmptySpot returns the coordinates for the first empty spot on the board.
func (b *Board) firstEmptySpot() (int, int) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.spaces[i][j] == EmptySpot {
				return i, j
			}
		}
	}

	return -1, -1
}

// PlayAI plays the move for the AI opponent.
//
// In easy mode, the AI chooses the first empty spot.
// In normal mode, the AI guesses randomly.
// In hard mode, the AI plays with a perfect strategy.
func (b *Board) PlayAI() {
	var x, y int

	switch b.mode {
	case Easy:
		x, y = b.firstEmptySpot()
	case Normal:
		for {
			x = rand.Intn(BoardSize)
			y = rand.Intn(BoardSize)

			if b.spaces[x][y] == EmptySpot {
				break
			}
		}
	case Hard:
		x, y = b.bestMoveForAI(b.spaces)
	}

	b.play(x, y, AIPlayer)
}

// IsGameOver returns true if the game is over. This is true if the board is full or there is a winner.
func (b *Board) IsGameOver() bool {
	_, hasWinner := b.GameHasWinner()

	return b.IsBoardFull() || hasWinner
}

// isBoardFull returns true if the given board is full.
func (b *Board) isBoardFull(board GameBoard) bool {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == EmptySpot {
				return false
			}
		}
	}

	return true
}

// SpotIsAllowed returns true if the spot is empty and a player can play that spot.
func (b *Board) SpotIsAllowed(x, y int) bool {
	// spots outside of our game board are not allowed
	if x < 0 || x >= BoardSize || y < 0 || y >= BoardSize {
		return false
	}

	// spots that are already played are not allowed
	return b.spaces[x][y] == EmptySpot
}

// IsBoardFull returns true if the board is full.
func (b *Board) IsBoardFull() bool {
	return b.isBoardFull(b.spaces)
}

// gameBoardWinner returns the player who won the provided board and a boolean indicating if there is a winner for the given board.
func (b *Board) gameBoardWinner(board GameBoard) (Player, bool) {
	for _, p := range b.winningPatterns {
		if board[p[0].X][p[0].Y] == board[p[1].X][p[1].Y] &&
			board[p[1].X][p[1].Y] == board[p[2].X][p[2].Y] &&
			board[p[0].X][p[0].Y] != EmptySpot {
			if board[p[0].X][p[0].Y] == string(HumanPlayer) {
				return HumanPlayer, true
			}

			return AIPlayer, true
		}
	}

	return "", false
}

// GameHasWinner return the player who won the game and a boolean indicating if there is a winner.
func (b *Board) GameHasWinner() (Player, bool) {
	return b.gameBoardWinner(b.spaces)
}

// PrintGameState prints the current board and game state.
func (b *Board) PrintGameState() {
	b.Print()

	player, hasWinner := b.GameHasWinner()

	if hasWinner {
		fmt.Printf("Game over. %s wins!\n\n", player)
		return
	}

	if b.IsBoardFull() {
		fmt.Printf("Game over. It's a draw!\n\n")
		return
	}
}

// minimax is a recursive function that returns a score for a given board.
//
// Minimax algorithm for Tic-Tac-Toe described here: https://www.neverstopbuilding.com/blog/minimax
func (b *Board) minimax(board GameBoard, depth int, isMaximizing bool) int {
	player, hasWinner := b.gameBoardWinner(board)

	if hasWinner && player == HumanPlayer {
		return -10 + depth
	}

	if hasWinner && player == AIPlayer {
		return 10 - depth
	}

	if b.isBoardFull(board) {
		return 0
	}

	bestScore, playerToSearch, nextIsMaximizing := func() (int, Player, bool) {
		if isMaximizing {
			return -100, AIPlayer, false
		}

		return 100, HumanPlayer, true
	}()

	compFunc := func(score, bestScore int) bool {
		if isMaximizing {
			return score > bestScore
		}

		return score < bestScore
	}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == EmptySpot {
				board[i][j] = string(playerToSearch)
				score := b.minimax(board, depth+1, nextIsMaximizing)
				board[i][j] = EmptySpot

				if compFunc(score, bestScore) {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

// bestMoveForAI return the best move for the AI hard mode opponent.
func (b *Board) bestMoveForAI(board GameBoard) (int, int) {
	bestScore := -100

	var moveX, moveY int

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == EmptySpot {
				board[i][j] = string(AIPlayer)
				score := b.minimax(board, 0, false)
				board[i][j] = EmptySpot

				if score > bestScore {
					bestScore = score
					moveX, moveY = i, j
				}
			}
		}
	}

	return moveX, moveY
}
