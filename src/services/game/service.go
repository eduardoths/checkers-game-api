package game

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

type GameService struct{}

func NewGameService() *GameService {
	return &GameService{}
}

func (this *GameService) NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error) {
	game := &structs.Game{
		ID:              uuid.New(),
		PlayerOne:       playerOne,
		PlayerTwo:       playerTwo,
		Board:           structs.NewBoard(playerOne, playerTwo),
		IsPlayerOneTurn: true,
	}
	return game, nil
}
