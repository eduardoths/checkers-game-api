package structs_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/domain"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/eduardoths/checkers-game/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	playerOne := mocks.FakePlayerOne()
	playerTwo := mocks.FakePlayerTwo()
	board := *structs.NewBoard(&playerOne, &playerTwo)
	want := mocks.FakeBoard()
	t.Run("It should fill player one correctly", func(t *testing.T) {
		assert.Equal(t, want[:23], board[:23])
	})

	t.Run("It should fill middle rows with nil", func(t *testing.T) {
		assert.Equal(t, want[24:40], board[24:40])
	})

	t.Run("It should player two rows correctly", func(t *testing.T) {
		assert.Equal(t, want[41:], board[41:])
	})
}

func TestGetCheckerFromPos(t *testing.T) {
	checker := mocks.FakeChecker()
	filledBoard := structs.Board{}
	for i := 0; i < len(filledBoard); i++ {
		newChecker := mocks.FakeChecker()
		filledBoard[i] = &newChecker
	}

	type testCase struct {
		board structs.Board
		pos   int
		want  *structs.Checker
	}
	testCases := map[string]testCase{
		"should return nil if pos is before board": {
			board: filledBoard,
			pos:   structs.BOARD_INIT - 1,
			want:  nil,
		},
		"should return nil if pos is after board": {
			board: filledBoard,
			pos:   structs.BOARD_END + 1,
			want:  nil,
		},
		"should return board from position": {
			board: structs.Board{nil, &checker},
			pos:   1,
			want:  &checker,
		},
		"should return nil if there's no checker at the selected position": {
			board: structs.Board{nil, &checker},
			pos:   0,
			want:  nil,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			actual := tc.board.GetCheckerFromPos(tc.pos)
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestMoveCheckerErrors(t *testing.T) {
	type testCase struct {
		board  structs.Board
		from   int
		moveBy int
		want   error
	}

	testCases := map[string]testCase{
		"it should not allow to move to the same position": {
			board:  structs.Board{nil, &structs.Checker{}},
			from:   1,
			moveBy: 0,
			want:   domain.ErrInvalidMovement,
		},
		"it should allow to move to the next diagonal": {
			board:  structs.Board{nil, &structs.Checker{}},
			from:   1,
			moveBy: 7,
			want:   nil,
		},
		"it should return error if column is invalid": {
			board:  mocks.FakeBoard(),
			from:   0,
			moveBy: 7,
			want:   domain.ErrInvalidMovement,
		},
		"it should return error if position is before board init": {
			board:  mocks.FakeBoard(),
			from:   structs.BOARD_INIT - 1,
			moveBy: 7,
			want:   domain.ErrInvalidMovement,
		},
		"it should return error if position is after board end": {
			board:  mocks.FakeBoard(),
			from:   structs.BOARD_END + 1,
			moveBy: 7,
			want:   domain.ErrInvalidMovement,
		},
		"it should return error if there's a checker on it": {
			board:  mocks.FakeBoard(),
			from:   1,
			moveBy: 7,
			want:   domain.ErrInvalidMovement,
		},
	}
	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			actual := tc.board.Move(tc.from, tc.moveBy)
			assert.Equal(t, tc.want, actual)
		})
	}
}
