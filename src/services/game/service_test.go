package game_test

import (
	"errors"
	"testing"

	"github.com/eduardoths/checkers-game/mockgen"
	"github.com/eduardoths/checkers-game/src/interfaces"
	"github.com/eduardoths/checkers-game/src/repositories"
	"github.com/eduardoths/checkers-game/src/services/game"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/eduardoths/checkers-game/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	PlayerOneID = uuid.MustParse("a607b8b5-e8a7-4ce9-b001-8f1b3b154490")
	PlayerTwoID = uuid.MustParse("51fa65e9-2637-4fcf-8509-04cc521809d6")
)

type mockgenContainer struct {
	controller *gomock.Controller
	repository *mockgen.MockGameRepository
}

func MakeGameService(t *testing.T) (interfaces.GameService, mockgenContainer) {
	ctrl := gomock.NewController(t)
	mc := mockgenContainer{
		controller: ctrl,
		repository: mockgen.NewMockGameRepository(ctrl),
	}
	rc := repositories.RepositoriesContainer{
		GameRepository: mc.repository,
	}
	return game.NewGameService(rc), mc
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
		service, mc := MakeGameService(t)
		defer mc.controller.Finish()
		actual, err := service.NewGame(tc.input.playerOne, tc.input.playerTwo)
		tc.assert(t, tc.input, output{game: actual, err: err})
	}
}

func TestMoveChecker(t *testing.T) {
	type input struct {
		from      int
		movements []int
		gameID    uuid.UUID
	}

	type output struct {
		game *structs.Game
		err  error
	}

	type testCase struct {
		description string
		input       input
		before      func(in input, mc *mockgenContainer)
		assert      func(t *testing.T, in input, actual output)
	}

	defaultStubs := func(in input, mc *mockgenContainer) {
		defaultGame := mocks.FakeGame()
		defaultGame.ID = in.gameID
		exp := mc.repository.EXPECT()
		stubs := []*gomock.Call{
			exp.FindGame(gomock.Any()).Return(&defaultGame, nil),
		}
		for i := range stubs {
			stubs[i].AnyTimes()
		}
	}

	testCases := []testCase{
		{
			description: "Should throw error if game's not found",
			input:       input{gameID: uuid.New()},
			before: func(in input, mc *mockgenContainer) {
				err := errors.New("not found")
				mc.repository.EXPECT().FindGame(gomock.Eq(in.gameID)).Return(nil, err)
			},
			assert: func(t *testing.T, in input, actual output) {
				assert.Nil(t, actual.game)
				assert.Equal(t, errors.New("not found"), actual.err)
			},
		},
		{
			description: "Should throw error if select checker is before board init",
			input:       input{from: structs.BOARD_INIT - 1},
			before:      defaultStubs,
			assert: func(t *testing.T, in input, actual output) {
				wantErr := errors.New("invalid_field:checker position is outside of board")
				assert.Nil(t, actual.game)
				assert.Equal(t, wantErr, actual.err)
			},
		},
		{
			description: "Should throw error if select checker is after board end",
			before:      defaultStubs,
			input:       input{from: structs.BOARD_END + 1},
			assert: func(t *testing.T, in input, actual output) {
				wantErr := errors.New("invalid_field:checker position is outside of board")
				assert.Nil(t, actual.game)
				assert.Equal(t, wantErr, actual.err)
			},
		},
		{
			description: "Should throw error if checker is nil at selected pos",
			before:      defaultStubs,
			input:       input{from: structs.BOARD_INIT},
			assert: func(t *testing.T, in input, actual output) {
				wantErr := errors.New("invalid_field:no checker at selected position")
				assert.Nil(t, actual.game)
				assert.Equal(t, wantErr, actual.err)
			},
		},
		{
			description: "Should throw error if it's not the player one turn",
			before: func(in input, mc *mockgenContainer) {
				game := mocks.FakeGame()
				checker := mocks.FakeCheckerFromPlayer(game.PlayerOne)
				board := &structs.Board{&checker}
				game.Board = board
				game.IsPlayerOneTurn = false
				game.ID = in.gameID
				mc.repository.EXPECT().FindGame(gomock.Eq(in.gameID)).Return(&game, nil)
			},
			input: input{gameID: uuid.New(), from: 0},
			assert: func(t *testing.T, in input, actual output) {
				wantErr := errors.New("invalid_field:it's not the player's turn")
				assert.Nil(t, actual.game)
				assert.Equal(t, wantErr, actual.err)
			},
		},
		{
			description: "Should throw error if it's not the player two turns",
			before: func(in input, mc *mockgenContainer) {
				game := mocks.FakeGame()
				checker := mocks.FakeCheckerFromPlayer(game.PlayerTwo)
				board := &structs.Board{&checker}
				game.Board = board
				game.ID = in.gameID
				mc.repository.EXPECT().FindGame(gomock.Eq(in.gameID)).Return(&game, nil)
			},
			input: input{gameID: uuid.New(), from: 0},
			assert: func(t *testing.T, in input, actual output) {
				wantErr := errors.New("invalid_field:it's not the player's turn")
				assert.Nil(t, actual.game)
				assert.Equal(t, wantErr, actual.err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service, mc := MakeGameService(t)
			defer mc.controller.Finish()
			tc.before(tc.input, &mc)
			game, err := service.Move(tc.input.gameID, tc.input.from, tc.input.movements)
			tc.assert(t, tc.input, output{game: game, err: err})
		})
	}
}
