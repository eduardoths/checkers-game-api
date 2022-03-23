package game

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

func (this *GameUseCases) Move(gameID uuid.UUID, from int, movements []int) (*structs.Game, error) {
	game, err := this.repository.FindGame(gameID)
	if err != nil {
		return nil, err
	}

	if err := game.ExecuteMovements(from, movements); err != nil {
		return nil, err
	}

	return game, nil
}
