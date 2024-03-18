package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HumanInput interface {
	Scan() string
}

type ConsoleHumanInput struct{}

//type FileHumanInput struct{}

func (in *ConsoleHumanInput) Scan() string {
	//fmt.Printf("Podaj miejsce wpisania " + p + "\n")
	input := ""
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		input = scanner.Text()
		if input != "" {
			break
		}
	}
	return input
}

func ScanMove(in HumanInput) (int, error) {
	input := in.Scan()
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
