package main

import "errors"

const (
	emptyBoardField = "" //const - stala wartosc, ktorej nie mozna zmienic
	player1         = "X"
	player2         = "O"
	boardSize       = 9
)

type (
	Player = string
	Board  = []string
	Game   struct {
		board  Board //wskaznik by default
		player Player
		//turnNumber int
	}
)

var (
	errBoardPlaceTaken = errors.New("board place already used by another player")
	winnigCoordinates  = [][]int{
		{0, 1, 2}, //1 rzad
		{0, 3, 6}, //1 columna
		{0, 4, 8}, //diagonal 1
	}
)

// constructor zwracajacy zawsze nowa instancje dla struktury game (nowy adres miejsca w pamieci)
// sa tez dekonstruktory
func newGame() *Game {
	//var g Game
	//g.board --> nil
	//g.board=append(g.board, "X")
	//var a [9]string
	//a[2] = "x"
	return &Game{
		board: make([]string, 9), //alokacja tablicy string z 9 elementami
	}
}

func BoardPlacement(game *Game, number int) error {
	if game.board[number] == player1 || game.board[number] == player2 {
		return errBoardPlaceTaken
	} else {
		game.player = SwitchPlayer(game)
		game.board[number] = game.player
		return nil
	}
}

func SwitchPlayer(game *Game) string {

	if game.player == player1 {
		return player2
	} else {
		return player1
	}

}

func IsWin(board Board, p Player) bool {
	for _, coordinates := range winnigCoordinates {
		//1 rzad 0,1,2
		if board[coordinates[0]] == p && board[coordinates[1]] == p && board[coordinates[2]] == p {
			return true
		}
	}
	// TODO wypelnic coordinates
	//if game.board[0] == player1 && game.board[4] == player1 && game.board[8] == player1 || game.board[0] == player2 && game.board[4] == player2 && game.board[8] == player2 {
	//	return true
	//}
	//if game.board[2] == player && game.board[4] == player && game.board[6] == player || game.board[2] == player2 && game.board[4] == player2 && game.board[6] == player2 {
	//	return true
	//}
	//if game.board[0] == player1 && game.board[1] == player1 && game.board[2] == player1 || game.board[0] == player2 && game.board[1] == player2 && game.board[2] == player2 {
	//	return true
	//}
	//if game.board[3] == player1 && game.board[4] == player1 && game.board[5] == player1 || game.board[3] == player2 && game.board[4] == player2 && game.board[5] == player2 {
	//	return true
	//}
	//if game.board[6] == player1 && game.board[7] == player1 && game.board[8] == player1 || game.board[6] == player2 && game.board[7] == player2 && game.board[8] == player2 {
	//	return true
	//}
	//if game.board[0] == player1 && game.board[3] == player1 && game.board[6] == player1 || game.board[0] == player2 && game.board[3] == player2 && game.board[6] == player2 {
	//	return true
	//}
	//if game.board[1] == player1 && game.board[4] == player1 && game.board[7] == player1 || game.board[1] == player2 && game.board[4] == player2 && game.board[7] == player2 {
	//	return true
	//}
	//if game.board[2] == player1 && game.board[5] == player1 && game.board[8] == player1 || game.board[2] == player2 && game.board[5] == player2 && game.board[8] == player2 {
	//	return true
	//}

	return false
}
