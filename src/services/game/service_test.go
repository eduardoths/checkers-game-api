package game_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/services/game"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	PlayerOneID = uuid.MustParse("a607b8b5-e8a7-4ce9-b001-8f1b3b154490")
	PlayerTwoID = uuid.MustParse("a607b8b5-e8a7-4ce9-b001-8f1b3b154490")
)

func MakeGameService() *game.GameService {
	return game.NewGameService()
}

func TestNewGame(t *testing.T) {
	playerOne := &structs.Player{ID: PlayerOneID}
	playerTwo := &structs.Player{ID: PlayerTwoID}
	t.Run("Should return non nil game", func(t *testing.T) {
		service := MakeGameService()

		game, err := service.NewGame(playerOne, playerTwo)

		assert.NoError(t, err)
		assert.NotNil(t, game, "assert game is not nil")
	})

	t.Run("Should have correct player one set", func(t *testing.T) {
		service := MakeGameService()

		actual, err := service.NewGame(playerOne, playerTwo)

		assert.NoError(t, err)
		assert.Equal(t, playerOne, actual.PlayerOne)
	})

	t.Run("Should have correct player two set", func(t *testing.T) {
		service := MakeGameService()

		actual, err := service.NewGame(playerOne, playerTwo)

		assert.NoError(t, err)
		assert.Equal(t, playerTwo, actual.PlayerTwo)
	})
}
