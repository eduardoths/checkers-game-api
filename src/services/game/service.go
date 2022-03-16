package game

import "github.com/eduardoths/checkers-game/src/structs"

type GameService struct{}

func NewGameService() *GameService {
	return &GameService{}
}

func (this *GameService) NewGame(PlayerOne, PlayerTwo *structs.Player) (*structs.Game, error) {
	return &structs.Game{PlayerOne: PlayerOne}, nil
}
