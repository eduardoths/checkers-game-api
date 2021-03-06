package game_test

import (
	"errors"
	"testing"

	"github.com/eduardoths/checkers-game/mockgen"
	"github.com/eduardoths/checkers-game/src/domain"
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

func defaultStubs(mc *mockgenContainer) {
	defaultGame := mocks.FakeGame()
	exp := mc.repository.EXPECT()
	stubs := []*gomock.Call{
		exp.FindGame(gomock.Any()).Return(&defaultGame, nil),
		exp.SaveGame(gomock.Any()).Return(nil),
	}
	for i := range stubs {
		stubs[i].AnyTimes()
	}
}

func MakeGameUseCases(t *testing.T) (interfaces.GameService, mockgenContainer) {
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
		before      func(in input, mc *mockgenContainer)
		input       input
		assert      func(*testing.T, input, output)
	}

	defaultBeforeCallback := func(_ input, mc *mockgenContainer) {
		defaultStubs(mc)
	}

	testCases := []testCase{
		{
			description: "Should return non nil game",
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
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
			before:      defaultBeforeCallback,
			input: input{
				playerOne: nil,
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:player")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should throw error if playerTwo is nil",
			before:      defaultBeforeCallback,
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: nil,
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:player")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should throw error if player one and player two are the same",
			before:      defaultBeforeCallback,
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerOne(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("invalid_field:player")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should throw error if saving to repository returns error",
			before: func(in input, mc *mockgenContainer) {
				mc.repository.EXPECT().SaveGame(gomock.Any()).Return(errors.New("repository error"))
			},
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				wantErr := errors.New("repository error")
				assert.Equal(t, wantErr, actual.err)
				assert.Nil(t, actual.game)
			},
		},
		{
			description: "Should return game if saving to repository is ok",
			before: func(in input, mc *mockgenContainer) {
				mc.repository.EXPECT().SaveGame(gomock.Any()).Return(nil)
			},
			input: input{
				playerOne: makePlayerOne(),
				playerTwo: makePlayerTwo(),
			},
			assert: func(t *testing.T, _ input, actual output) {
				assert.NoError(t, actual.err)
				assert.NotNil(t, actual.game)
			},
		},
	}

	for _, tc := range testCases {
		useCases, mc := MakeGameUseCases(t)
		defer mc.controller.Finish()
		tc.before(tc.input, &mc)
		actual, err := useCases.NewGame(tc.input.playerOne, tc.input.playerTwo)
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
			description: "Should throw error if game's movement execution returns error",
			input:       input{gameID: uuid.New()},
			before: func(in input, mc *mockgenContainer) {
				game := mocks.FakeGame()
				game.Board = &structs.Board{}
				mc.repository.EXPECT().FindGame(gomock.Eq(in.gameID)).Return(&game, nil)
			},
			assert: func(t *testing.T, in input, actual output) {
				assert.Nil(t, actual.game)
				assert.Equal(t, domain.ErrNoCheckerAtSelectedPosition, actual.err)
			},
		},
		{
			description: "Should return correct board if valid movement is made",
			input:       input{gameID: uuid.New(), from: 17, movements: []int{9}},
			before: func(in input, mc *mockgenContainer) {
				game := mocks.FakeGame()
				game.ID = in.gameID
				mc.repository.EXPECT().FindGame(gomock.Eq(in.gameID)).Return(&game, nil)
			},
			assert: func(t *testing.T, in input, actual output) {
				wantPlayerOne := mocks.FakePlayerOne()
				wantPlayerTwo := mocks.FakePlayerTwo()
				wantBoard := mocks.FakeBoard()
				wantBoard[26] = wantBoard[17]
				wantBoard[17] = nil
				want := &structs.Game{
					ID:              in.gameID,
					Board:           &wantBoard,
					PlayerOne:       &wantPlayerOne,
					PlayerTwo:       &wantPlayerTwo,
					IsPlayerOneTurn: false,
				}
				assert.Equal(t, want, actual.game)
				assert.NoError(t, actual.err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			useCases, mc := MakeGameUseCases(t)
			defer mc.controller.Finish()
			tc.before(tc.input, &mc)
			game, err := useCases.Move(tc.input.gameID, tc.input.from, tc.input.movements)
			tc.assert(t, tc.input, output{game: game, err: err})
		})
	}
}
