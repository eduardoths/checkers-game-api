package game_test

import (
	"errors"
	"testing"

	"github.com/eduardoths/checkers-game/src/services/game"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	PlayerOneID = uuid.MustParse("a607b8b5-e8a7-4ce9-b001-8f1b3b154490")
	PlayerTwoID = uuid.MustParse("51fa65e9-2637-4fcf-8509-04cc521809d6")
)

func MakeGameService() *game.GameService {
	return game.NewGameService()
}

func TestNewGame(t *testing.T) {
	makePlayerOne := func() *structs.Player {
		return &structs.Player{ID: PlayerOneID}
	}
	makePlayerTwo := func() *structs.Player {
		return &structs.Player{ID: PlayerTwoID}
	}

	type input struct {
		playerOne *structs.Player
		playerTwo *structs.Player
	}

	type output struct {
		game *structs.Game
		err  error
	}

	type testCase struct {
		description string
		input       input
		assert      func(*testing.T, input, output)
	}

	testCases := []testCase{
		{
			description: "Should return non nil game",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				assert.NoError(t, actual.err)
				assert.NotNil(t, actual.game)
			},
		},
		{
			description: "Should have correct player one set",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, in input, actual output) {
				assert.NoError(t, actual.err)
				assert.Equal(t, in.playerOne, actual.game.PlayerOne)
			},
		},
		{
			description: "Should have correct player two set",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, in input, actual output) {
				assert.NoError(t, actual.err)
				assert.Equal(t, in.playerTwo, actual.game.PlayerTwo)
			},
		},
		{
			description: "Should have a new board",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, in input, actual output) {
				wantBoard := structs.NewBoard(in.playerOne, in.playerTwo)
				assert.NoError(t, actual.err)
				assert.Equal(t, wantBoard, actual.game.Board)
			},
		},
		{
			description: "Should be player one's turn",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				assert.NoError(t, actual.err)
				assert.Equal(t, true, actual.game.IsPlayerOneTurn)
			},
		},
		{
			description: "Should return new id for game",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				assert.NoError(t, actual.err)
				assert.NotEqual(t, uuid.Nil, actual.game.ID)
			},
		},
		{
			description: "Should throw error if playerOne is nil",
			input: input{
				playerOne: nil,
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:player is nil")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should throw error if playerTwo is nil",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: nil,
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:player is nil")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should throw error if player one and player two are the same",
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerOne(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:both players are the same")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
	}
	for _, tc := range testCases {
		service := MakeGameService()
		actual, err := service.NewGame(tc.input.playerOne, tc.input.playerTwo)
		tc.assert(t, tc.input, output{game: actual, err: err})
	}
}
