package structs_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/domain"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/eduardoths/checkers-game/src/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCurrentPlayerID(t *testing.T) {
	playerOne := mocks.FakePlayerOne()
	playerTwo := mocks.FakePlayerTwo()
	type testCase struct {
		game structs.Game
		want uuid.UUID
	}
	testCases := map[string]testCase{
		"should return player one's if if IsPlayerOneTurn equals true": {
			game: structs.Game{
				PlayerOne:       &playerOne,
				PlayerTwo:       &playerTwo,
				IsPlayerOneTurn: true,
			},
			want: playerOne.ID,
		},
		"should return player two's  id if IsPlayerOneTurn equals false": {
			game: structs.Game{
				PlayerOne:       &playerOne,
				PlayerTwo:       &playerTwo,
				IsPlayerOneTurn: false,
			},
			want: playerTwo.ID,
		},
	}
	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			actual := tc.game.CurrentPlayerID()
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestExecuteMovement(t *testing.T) {
	type input struct {
		from   int
		moveBy []int
	}

	type testCase struct {
		game  structs.Game
		input input
		want  error
	}

	notPlayerOneTurn := mocks.FakeGame()
	notPlayerTwoTurn := mocks.FakeGame()

	notPlayerOneTurn.IsPlayerOneTurn = false

	testCases := map[string]testCase{
		"Should throw error if checker is nil at selected pos": {
			game:  mocks.FakeGame(),
			input: input{from: structs.BOARD_INIT},
			want:  domain.ErrNoCheckerAtSelectedPosition,
		},
		"Should throw error if it's not the player one turn": {
			game:  notPlayerOneTurn,
			input: input{from: 17, moveBy: []int{9}},
			want:  domain.ErrNotPlayersTurn,
		},
		"Should throw error if it's not the player two turns": {
			game:  notPlayerTwoTurn,
			want:  domain.ErrNotPlayersTurn,
			input: input{from: 40, moveBy: []int{-7}},
		},
		"It should throw error if movement validation fails": {
			game:  mocks.FakeGame(),
			want:  domain.ErrInvalidMovement,
			input: input{from: 1, moveBy: []int{-1}},
		},
		"Should throw error if movements array's size is zero": {
			game:  mocks.FakeGame(),
			input: input{from: 1},
			want:  domain.ErrInvalidFieldMovementsArray,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			actualErr := tc.game.ExecuteMovements(tc.input.from, tc.input.moveBy)
			assert.Equal(t, tc.want, actualErr)
		})
	}
}
