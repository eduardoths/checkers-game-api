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

func TestExecuteMovementErrors(t *testing.T) {
	type input struct {
		from   int
		moveBy []int
	}

	type testCase struct {
		game  structs.Game
		input input
		want  error
	}
	playerOne := mocks.FakePlayerOne()
	playerTwo := mocks.FakePlayerTwo()
	playerOneMovingBackwards := structs.Board{}
	playerOneMovingBackwards[8] = &structs.Checker{Owner: &playerOne, IsKing: false}

	playerTwoMovingBackwards := structs.Board{nil, &structs.Checker{Owner: &playerTwo, IsKing: false}}

	playerOneKingMovingBackwards := structs.Board{}
	playerOneKingMovingBackwards[8] = &structs.Checker{Owner: &playerOne, IsKing: true}

	playerTwoKingMovingBackwards := structs.Board{nil, &structs.Checker{Owner: &playerTwo, IsKing: true}}

	multipleMovementsBoard := structs.Board{}
	multipleMovementsBoard[1] = &structs.Checker{Owner: &playerOne}
	multipleMovementsBoard[10] = &structs.Checker{Owner: &playerTwo}

	validMultipleMovementsBoard := structs.Board{}
	validMultipleMovementsBoard[1] = &structs.Checker{Owner: &playerOne, IsKing: true}
	validMultipleMovementsBoard[10] = &structs.Checker{Owner: &playerTwo}
	validMultipleMovementsBoard[5] = &structs.Checker{Owner: &playerTwo}

	notPlayerOneTurn := mocks.FakeGame()
	notPlayerTwoTurn := mocks.FakeGame()
	playerOneMovingBackwardsGame := structs.Game{
		Board:           &playerOneMovingBackwards,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: true,
	}
	playerTwoMovingBackwardsGame := structs.Game{
		Board:           &playerTwoMovingBackwards,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: false,
	}
	playerOneKingMovingBackwardsGame := structs.Game{
		Board:           &playerOneKingMovingBackwards,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: true,
	}
	playerTwoKingMovingBackwardsGame := structs.Game{
		Board:           &playerTwoKingMovingBackwards,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: false,
	}
	multipleMovementsWithNoJumpingOne := structs.Game{
		Board:           &multipleMovementsBoard,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: true,
	}
	validMultipleMovements := structs.Game{
		Board:           &validMultipleMovementsBoard,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		IsPlayerOneTurn: true,
	}

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
		"Should return error if player one non king tries to move backwards": {
			game:  playerOneMovingBackwardsGame,
			input: input{from: 8, moveBy: []int{-7}},
			want:  domain.ErrInvalidMovement,
		},
		"Should return error if player two non king tries to move backwards": {
			game:  playerTwoMovingBackwardsGame,
			input: input{from: 1, moveBy: []int{7}},
			want:  domain.ErrInvalidMovement,
		},
		"Should return nilif player one king tries to move backwards": {
			game:  playerOneKingMovingBackwardsGame,
			input: input{from: 8, moveBy: []int{-7}},
			want:  nil,
		},
		"Should return nil if player two king tries to move backwards": {
			game:  playerTwoKingMovingBackwardsGame,
			input: input{from: 1, moveBy: []int{7}},
			want:  nil,
		},
		"Should return error if sequence of multiple movements has a non jumping one": {
			game:  multipleMovementsWithNoJumpingOne,
			input: input{from: 1, moveBy: []int{18, 9}},
			want:  domain.ErrInvalidMovement,
		},
		"Should not return error if sequence of multiple movements is valid": {
			game:  validMultipleMovements,
			input: input{from: 1, moveBy: []int{18, -14}},
			want:  nil,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			actualErr := tc.game.ExecuteMovements(tc.input.from, tc.input.moveBy)
			assert.Equal(t, tc.want, actualErr)
		})
	}
}
