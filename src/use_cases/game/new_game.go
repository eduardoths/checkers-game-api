package game

import (
	"errors"

	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

func (this *GameUseCases) NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error) {
	if playerOne == nil || playerTwo == nil {
		return nil, errors.New("invalid_field:player")
	}
	if playerOne.ID == playerTwo.ID {
		return nil, errors.New("invalid_field:player")
	}
	game := &structs.Game{
		ID:              uuid.New(),
		PlayerOne:       playerOne,
		PlayerTwo:       playerTwo,
		Board:           structs.NewBoard(playerOne, playerTwo),
		IsPlayerOneTurn: true,
	}
	if err := this.repository.SaveGame(game); err != nil {
		return nil, err
	}
	return game, nil
}
