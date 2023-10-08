// Copyright 2023 Vitality South, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var intro = `
Greetings, Professor Falken.

░█▀▀▀█ █──█ █▀▀█ █── █── 　 █───█ █▀▀ 　 █▀▀█ █── █▀▀█ █──█ 　 █▀▀█ 　 █▀▀▀ █▀▀█ █▀▄▀█ █▀▀ ▀█ 
─▀▀▀▄▄ █▀▀█ █▄▄█ █── █── 　 █▄█▄█ █▀▀ 　 █──█ █── █▄▄█ █▄▄█ 　 █▄▄█ 　 █─▀█ █▄▄█ █─▀─█ █▀▀ █▀ 
░█▄▄▄█ ▀──▀ ▀──▀ ▀▀▀ ▀▀▀ 　 ─▀─▀─ ▀▀▀ 　 █▀▀▀ ▀▀▀ ▀──▀ ▄▄▄█ 　 ▀──▀ 　 ▀▀▀▀ ▀──▀ ▀───▀ ▀▀▀ ▄─


,------~~v,                               _--^\
|'         Ż\   ,__/Ż||                _/    /,_
/             \,/     /        ,,   ,,/^      Ż  vŻv-__
|                    /         |'~^Ż                   Ż\
\                   |        _/                     _  /^
\                 /        /                   ,~~^/|ŻŻ
 ^Ż~_            /         |          __,,  v__\   \/
     '~~,  ,Ż~Ż\ \          ^~       /    ~Ż  //
         \/     \/            \~,  ,/         Ż
                                 ~~
  UNITED STATES                  SOVIET UNION
`

func main() {
	fmt.Println(intro)

	// main app loop
	for {
		gameMode := SelectGameMode()

		if gameMode == Exit {
			break
		}

		board := NewBoard(gameMode)

		board.ShowTutorial()

		players := []Player{HumanPlayer, AIPlayer}

		// game play loop while game is in progress
		for !board.IsGameOver() {
			// let each player take a turn
			for _, player := range players {
				switch player {
				case HumanPlayer:
					letHumanPlay(board)
				case AIPlayer:
					fmt.Printf("\nAI's turn.\n")
					board.PlayAI()
				}

				board.PrintGameState()

				if board.IsGameOver() {
					break
				}
			}
		}
	}

	goodbye()
}

// letHumanPlay asks the human player for a move and plays it.
func letHumanPlay(board *Board) {
	// keep asking for a move until a valid one is given
	for {
		var x, y int

		fmt.Printf("\nYour move (y x): ")
		fmt.Scanf("%d %d", &x, &y)

		if !board.SpotIsAllowed(x, y) {
			fmt.Println("That spot is not allowed. Try again.")
			continue
		}

		board.Play(x, y)

		break
	}
}

// goodbye prints the AI's goodbye message.
func goodbye() {
	fmt.Printf("\n\nAnalyzing game...\n")

	delay := 7

	chars := [4]byte{'|', '/', '-', '\\'}

	for i := 4; i <= delay+4; i++ {
		for j := 0; j < 75; j++ {
			fmt.Print("\b")
		}

		for j := 0; j < 80; j++ {
			fmt.Print(rand.Intn(2))
		}

		for j := 0; j < 1; j++ {
			fmt.Print(string(chars[i%4]))
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Printf(" Analysis complete.")

	fmt.Printf("\n\nStrange game. The only winning move is not to play.\n\n")
}
