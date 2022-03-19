package structs_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/eduardoths/checkers-game/src/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCurrentPlayerID(t *testing.T) {
	playerOne := mocks.FakePlayer()
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
