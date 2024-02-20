package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	emptyBoardField = "" //const - stala wartosc, ktorej nie mozna zmienic
	player1         = "X"
	player2         = "O"
	boardSize       = 9
)

type Game struct {
	board      []string //wskaznik by default
	player     string
	turnNumber int
}

// 01010101010101
// 0xFFFFFFFFF
// 0xABC789797EDCF

func main() {
	game := newGame()
	game.player = "X"

	for i := 0; i < 9; i++ {
		if IsWin(game) == false {
			fmt.Printf("Podaj miejsce wpisania " + game.player + "\n")

			var number, err = ScanMove()
			println(err)

			BoardPlacement(game, number)

			PrintBoard(game)

			fmt.Printf("\n")
		} else {
			println("koniec gry")
			break
		}
	}

}

func ScanMove() (int, error) {
	//os - pakiet systemowy
	//Stdin - klawiatura
	//Stdout - consola (w linux 1)
	//Stderr - wyjscie na bledy (w linux 2)
	//cat file.txt 2>/dev/null (2=etderr)
	//cat file.txt | grep "dupa"

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	input := scanner.Text()
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("input musi byc liczba: %w", err)
	}
	if number < 0 || number > boardSize-1 {
		return 0, fmt.Errorf("ruch musi byc liczba 0-%d", boardSize)
	}
	return number, nil
}

// b []string === c *[]string --> []string to zawsze pointer do tablicy
func PrintBoard(game *Game) {

	for i, v := range game.board {
		if v != emptyBoardField { //tutaj "" to "magic value"
			fmt.Printf(game.board[i])
		} else {
			fmt.Printf("_")
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}

	}

}

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

func BoardPlacement(game *Game, number int) {
	if game.board[number] == player1 || game.board[number] == player2 {
		println("miejsce jest juz zajete")
	} else {
		game.player = SwitchPlayer(game)
		game.board[number] = game.player
	}
}

func SwitchPlayer(game *Game) string {

	if game.player == player1 {
		return player2
	} else {
		return player1
	}

}

func IsWin(game *Game) bool {
	if game.board[0] == player1 && game.board[4] == player1 && game.board[8] == player1 || game.board[0] == player2 && game.board[4] == player2 && game.board[8] == player2 {
		return true
	}
	if game.board[2] == player1 && game.board[4] == player1 && game.board[6] == player1 || game.board[2] == player2 && game.board[4] == player2 && game.board[6] == player2 {
		return true
	}
	if game.board[0] == player1 && game.board[1] == player1 && game.board[2] == player1 || game.board[0] == player2 && game.board[1] == player2 && game.board[2] == player2 {
		return true
	}
	if game.board[3] == player1 && game.board[4] == player1 && game.board[5] == player1 || game.board[3] == player2 && game.board[4] == player2 && game.board[5] == player2 {
		return true
	}
	if game.board[6] == player1 && game.board[7] == player1 && game.board[8] == player1 || game.board[6] == player2 && game.board[7] == player2 && game.board[8] == player2 {
		return true
	}
	if game.board[0] == player1 && game.board[3] == player1 && game.board[6] == player1 || game.board[0] == player2 && game.board[3] == player2 && game.board[6] == player2 {
		return true
	}
	if game.board[1] == player1 && game.board[4] == player1 && game.board[7] == player1 || game.board[1] == player2 && game.board[4] == player2 && game.board[7] == player2 {
		return true
	}
	if game.board[2] == player1 && game.board[5] == player1 && game.board[8] == player1 || game.board[2] == player2 && game.board[5] == player2 && game.board[8] == player2 {
		return true
	}
	return false
}
