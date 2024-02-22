package main

import (
	"fmt"
)

// 01010101010101
// 0xFFFFFFFFF
// 0xABC789797EDCF

func main() {
	game := newGame()
	game.player = player1

	//TODO zadanie domowe
	//1. dokonczyc 2 istniejace unit testy
	//2. napisac unit testy pod ta petle(main)
	//3. z ksiazki przeczytac patterny: commmand, strategy, template method (rozdzial 8)
	for {
		if IsWin(game.board, game.player) {
			println("koniec gry")
			break
		}

		fmt.Printf("Podaj miejsce wpisania " + game.player + "\n")

		var number, err = ScanMove()
		if err != nil {
			println(err)
			continue
		}

		if err := BoardPlacement(game, number); err != nil {
			println(err)
			continue
		}

		PrintBoard(game)
	}

}
