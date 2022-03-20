package structs_test

import (
	"testing"

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
