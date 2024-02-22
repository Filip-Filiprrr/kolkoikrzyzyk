package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Printf("\n")

}
