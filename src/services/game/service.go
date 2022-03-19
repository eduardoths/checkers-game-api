package game

import (
	"errors"

	"github.com/eduardoths/checkers-game/src/interfaces"
	"github.com/eduardoths/checkers-game/src/repositories"
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

type GameService struct {
	repository interfaces.GameRepository
}

func NewGameService(repos repositories.RepositoriesContainer) *GameService {
	return &GameService{
		repository: repos.GameRepository,
	}
}

func (this *GameService) NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error) {
	if playerOne == nil || playerTwo == nil {
		return nil, errors.New("invalid_field:player is nil")
	}
	if playerOne.ID == playerTwo.ID {
		return nil, errors.New("invalid_field:both players are the same")
	}
	game := &structs.Game{
		ID:              uuid.New(),
		PlayerOne:       playerOne,
		PlayerTwo:       playerTwo,
		Board:           structs.NewBoard(playerOne, playerTwo),
		IsPlayerOneTurn: true,
	}
	return game, nil
}

func (this *GameService) Move(gameID uuid.UUID, from int, movements []int) (*structs.Game, error) {
	_, err := this.repository.FindGame(gameID)
	if err != nil {
		return nil, err
	}
	if from < structs.BOARD_INIT || from > structs.BOARD_END {
		return nil, errors.New("invalid_field:checker position is outside of board")
	}
	return nil, errors.New("invalid_field:no checker at selected position")
}
