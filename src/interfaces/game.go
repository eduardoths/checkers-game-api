package interfaces

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

type GameService interface {
	NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error)
	Move(gameID uuid.UUID, from int, movements []int) (*structs.Game, error)
}

type GameRepository interface {
	FindGame(id uuid.UUID) (*structs.Game, error)
}
