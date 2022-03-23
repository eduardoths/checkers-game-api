package structs

import (
	"github.com/eduardoths/checkers-game/src/domain"
	"github.com/google/uuid"
)

type Game struct {
	ID              uuid.UUID
	Board           *Board
	PlayerOne       *Player
	PlayerTwo       *Player
	IsPlayerOneTurn bool
}

func (this *Game) CurrentPlayerID() uuid.UUID {
	if this.IsPlayerOneTurn {
		return this.PlayerOne.ID
	}
	return this.PlayerTwo.ID
}

func (this *Game) ExecuteMovements(from int, movements []int) error {
	source := this.Board.GetCheckerFromPos(from)
	if source == nil {
		return domain.ErrNoCheckerAtSelectedPosition
	}

	currentPlayerID := this.CurrentPlayerID()
	if source.Owner.ID != currentPlayerID {
		return domain.ErrNotPlayersTurn
	}

	if len(movements) == 0 {
		return domain.ErrInvalidFieldMovementsArray
	}
	validDirection := source.IsKing || (this.IsPlayerOneTurn == (movements[0] >= 0))
	if !validDirection {
		return domain.ErrInvalidMovement
	}
	if err := this.Board.Move(from, movements[0]); err != nil {
		return err
	}
	this.IsPlayerOneTurn = !this.IsPlayerOneTurn
	return nil
}
