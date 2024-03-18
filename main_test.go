package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock = cos co udaje pewne konkretne zachowanie
// unit test -> mock
type MockUserInput struct {
	input string
}

func (in *MockUserInput) Scan() string {
	return in.input
}

func TestMakeTurn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		game    *Game
		in      HumanInput
		want    bool
		expGame *Game
	}{
		{
			"make a move",
			&Game{
				board: []string{

					player1, emptyBoardField, player1,
					emptyBoardField, emptyBoardField, emptyBoardField,
					emptyBoardField, emptyBoardField, emptyBoardField,
				},
				player: player2,
			},
			&MockUserInput{
				input: "1",
			},
			false,
			&Game{
				board: []string{

					player1, player1, player1,
					emptyBoardField, emptyBoardField, emptyBoardField,
					emptyBoardField, emptyBoardField, emptyBoardField,
				},
				player: player1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MakeTurn(tt.game, tt.in), "MakeTurn(%v)", tt.game)
			assert.Equalf(t, *tt.expGame, *tt.game, "unexpected game state")
		})
	}
}
