package game

import "github.com/eduardoths/checkers-game/src/structs"

type GameService struct{}

func NewGameService() *GameService {
	return &GameService{}
}

func (this *GameService) NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error) {
	return &structs.Game{PlayerOne: playerOne}, nil
}
