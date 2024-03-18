package main

// 01010101010101
// 0xFFFFFFFFF
// 0xABC789797EDCF

func main() {
	game := newGame()
	game.player = player1
	in := &ConsoleHumanInput{}

	//TODO zadanie domowe
	//1. dokonczyc test make turn z main_test
	//2. zaimplementowac HumanOutput
	//3. pomyslec, jak zaimplementowac AIPlayer
	for {
		if MakeTurn(game, in) {
			break
		}
	}

}

func MakeTurn(game *Game, in HumanInput) bool {
	if IsWin(game.board, game.player) {
		println("koniec gry")
		return true
	}

	var number, err = ScanMove(in)
	if err != nil {
		println(err)
		return false
	}

	if err := BoardPlacement(game, number); err != nil {
		println(err)
		return false
	}

	PrintBoard(game)
	return false
}
