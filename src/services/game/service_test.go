package game_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/services/game"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	t.Run("Should return non nil game", func(t *testing.T) {
		service := game.NewGameService()
		game, err := service.NewGame(nil, nil)

		assert.NoError(t, err)
		assert.Equal(t, &structs.Game{}, game, "assert game")
	})
}
