package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_Game_IsWin_CheckAdditionalCases(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		game   Game
		expWin bool
	}{
		//testing table (remisy, przegrane)
		{
			name: "XXX________",
			game: Game{
				board: []string{

					player1, player1, player1,
					emptyBoardField, emptyBoardField, emptyBoardField,
					emptyBoardField, emptyBoardField, emptyBoardField,
				},
				player: player1,
			},
			expWin: true,
		},
		{
			name: "XYX________",
			game: Game{
				board: []string{

					player1, player2, player1,
					emptyBoardField, emptyBoardField, emptyBoardField,
					emptyBoardField, emptyBoardField, emptyBoardField,
				},
				player: player1,
			},
			expWin: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsWin(tt.game.board, tt.game.player); got != tt.expWin {
				assert.True(t, got, "unexpected game result")
			}
		})
	}
}

func Test_Game_IsWin_CheckWinningCoordinates(t *testing.T) {
	t.Parallel()

	for idx, coordinates := range winnigCoordinates {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			t.Parallel()
			board := make([]string, 9)
			board[coordinates[0]] = player1
			board[coordinates[1]] = player1
			board[coordinates[2]] = player1
			assert.True(t, IsWin(board, player1), "unexpected game result")
		})
	}
}
