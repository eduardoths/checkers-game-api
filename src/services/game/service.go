package game

import (
	"errors"

	"github.com/eduardoths/checkers-game/src/domain"
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

func (this *GameService) Move(gameID uuid.UUID, from int, movements []int) (*structs.Game, error) {
	game, err := this.repository.FindGame(gameID)
	if err != nil {
		return nil, err
	}

	currentPlayerID := game.CurrentPlayerID()

	source := game.Board.GetCheckerFromPos(from)
	if source == nil {
		return nil, domain.ErrNoCheckerAtSelectedPosition
	}

	if source.Owner.ID != currentPlayerID {
		return nil, domain.ErrNotPlayersTurn
	}

	isValidPosition := game.Board.IsValidPosition(from + movements[0])
	if !isValidPosition {
		return nil, domain.ErrInvalidMovement
	}

	return nil, nil
}
